package scheduler

import (
	"context"
	"fmt"
	"sync"
	"time"
)

type Frequency int

const (
	Hourly Frequency = iota
	Daily
	Weekly
)

func (f Frequency) Duration() time.Duration {
	switch f {
	case Hourly:
		return time.Hour
	case Daily:
		return 24 * time.Hour
	case Weekly:
		return 7 * 24 * time.Hour
	default:
		return 24 * time.Hour
	}
}

type Job interface {
	Name() string
	Execute(ctx context.Context) error
}

type Logger interface {
	Info(msg string, fields map[string]any)
	Error(msg string, fields map[string]any)
}

type registeredJob struct {
	job       Job
	frequency Frequency
}

type Scheduler struct {
	mu      sync.Mutex
	jobs    []registeredJob
	log     Logger
	enabled bool
	cancel  context.CancelFunc
	wg      sync.WaitGroup
	running bool
}

func New(log Logger, enabled bool) *Scheduler {
	return &Scheduler{
		log:     log,
		enabled: enabled,
	}
}

func (s *Scheduler) Register(job Job, frequency Frequency) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.jobs = append(s.jobs, registeredJob{job: job, frequency: frequency})
	s.log.Info("job registered", map[string]any{"job": job.Name(), "frequency": fmt.Sprintf("%v", frequency)})
}

func (s *Scheduler) Start(ctx context.Context) {
	if !s.enabled {
		s.log.Info("scheduler disabled", nil)
		return
	}

	s.mu.Lock()
	if s.running {
		s.mu.Unlock()
		return
	}
	s.running = true
	jobCtx, cancel := context.WithCancel(ctx)
	s.cancel = cancel
	s.mu.Unlock()

	for _, rj := range s.jobs {
		rj := rj
		s.wg.Add(1)
		go s.runJob(jobCtx, rj)
	}

	s.log.Info("scheduler started", map[string]any{"jobs": len(s.jobs)})
}

func (s *Scheduler) Stop() {
	s.mu.Lock()
	if !s.running {
		s.mu.Unlock()
		return
	}
	s.running = false
	if s.cancel != nil {
		s.cancel()
	}
	s.mu.Unlock()

	s.wg.Wait()
	s.log.Info("scheduler stopped", nil)
}

func (s *Scheduler) runJob(ctx context.Context, rj registeredJob) {
	defer s.wg.Done()
	defer func() {
		if r := recover(); r != nil {
			s.log.Error("job panicked", map[string]any{"job": rj.job.Name(), "panic": fmt.Sprintf("%v", r)})
		}
	}()

	s.executeJob(ctx, rj.job)

	interval := rj.frequency.Duration()
	ticker := time.NewTicker(interval)
	defer ticker.Stop()

	for {
		select {
		case <-ctx.Done():
			return
		case <-ticker.C:
			s.executeJob(ctx, rj.job)
		}
	}
}

func (s *Scheduler) executeJob(ctx context.Context, job Job) {
	start := time.Now()
	s.log.Info("job started", map[string]any{"job": job.Name()})

	err := job.Execute(ctx)
	duration := time.Since(start)

	if err != nil {
		s.log.Error("job failed", map[string]any{
			"job":      job.Name(),
			"error":    err.Error(),
			"duration": duration.String(),
		})
		return
	}

	s.log.Info("job completed", map[string]any{
		"job":      job.Name(),
		"duration": duration.String(),
	})
}
