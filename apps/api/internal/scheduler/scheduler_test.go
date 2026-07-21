package scheduler

import (
	"context"
	"sync"
	"testing"
	"time"
)

type mockLogger struct {
	mu  sync.Mutex
	msg []string
}

func (l *mockLogger) Info(msg string, _ map[string]any) {
	l.mu.Lock()
	defer l.mu.Unlock()
	l.msg = append(l.msg, msg)
}

func (l *mockLogger) Error(msg string, _ map[string]any) {
	l.mu.Lock()
	defer l.mu.Unlock()
	l.msg = append(l.msg, msg)
}

type testJob struct {
	name    string
	execute func(ctx context.Context) error
	count   int
	mu      sync.Mutex
}

func (j *testJob) Name() string { return j.name }

func (j *testJob) Execute(ctx context.Context) error {
	j.mu.Lock()
	j.count++
	j.mu.Unlock()
	if j.execute != nil {
		return j.execute(ctx)
	}
	return nil
}

func TestScheduler_Disabled(t *testing.T) {
	log := &mockLogger{}
	s := New(log, false)
	s.Start(context.Background())

	if s.running {
		t.Fatal("scheduler should not be running when disabled")
	}

	log.mu.Lock()
	defer log.mu.Unlock()
	for _, m := range log.msg {
		if m == "scheduler disabled" {
			return
		}
	}
	t.Fatal("expected 'scheduler disabled' log")
}

func TestScheduler_RegisterAndRun(t *testing.T) {
	log := &mockLogger{}
	s := New(log, true)

	job := &testJob{name: "test-job"}
	s.Register(job, Hourly)

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	s.Start(ctx)
	defer s.Stop()

	time.Sleep(50 * time.Millisecond)

	job.mu.Lock()
	count := job.count
	job.mu.Unlock()
	if count < 1 {
		t.Fatalf("expected job to execute at least once, got %d", count)
	}

	log.mu.Lock()
	defer log.mu.Unlock()
	for _, m := range log.msg {
		if m == "job completed" {
			return
		}
	}
	t.Fatal("expected 'job completed' log")
}

func TestScheduler_JobFailure(t *testing.T) {
	log := &mockLogger{}
	s := New(log, true)

	job := &testJob{
		name: "failing-job",
		execute: func(_ context.Context) error {
			return errTest
		},
	}
	s.Register(job, Hourly)

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	s.Start(ctx)
	defer s.Stop()

	time.Sleep(50 * time.Millisecond)

	log.mu.Lock()
	defer log.mu.Unlock()
	for _, m := range log.msg {
		if m == "job failed" {
			return
		}
	}
	t.Fatal("expected 'job failed' log")
}

var errTest = &simpleError{msg: "test error"}

type simpleError struct{ msg string }

func (e *simpleError) Error() string { return e.msg }

func TestScheduler_JobPanic(t *testing.T) {
	log := &mockLogger{}
	s := New(log, true)

	job := &testJob{
		name: "panicking-job",
		execute: func(_ context.Context) error {
			panic("simulated panic")
		},
	}
	s.Register(job, Hourly)

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	s.Start(ctx)
	defer s.Stop()

	time.Sleep(50 * time.Millisecond)

	log.mu.Lock()
	defer log.mu.Unlock()
	for _, m := range log.msg {
		if m == "job panicked" {
			return
		}
	}
	t.Fatal("expected 'job panicked' log")
}

func TestScheduler_MultipleJobs(t *testing.T) {
	log := &mockLogger{}
	s := New(log, true)

	job1 := &testJob{name: "job-1"}
	job2 := &testJob{name: "job-2"}
	s.Register(job1, Hourly)
	s.Register(job2, Daily)

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	s.Start(ctx)
	defer s.Stop()

	time.Sleep(50 * time.Millisecond)

	job1.mu.Lock()
	count1 := job1.count
	job1.mu.Unlock()
	job2.mu.Lock()
	count2 := job2.count
	job2.mu.Unlock()

	if count1 < 1 {
		t.Fatalf("expected job-1 to execute, got %d", count1)
	}
	if count2 < 1 {
		t.Fatalf("expected job-2 to execute, got %d", count2)
	}
}

func TestScheduler_StartTwice(t *testing.T) {
	log := &mockLogger{}
	s := New(log, true)

	job := &testJob{name: "double-start-job"}
	s.Register(job, Hourly)

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	s.Start(ctx)
	s.Start(ctx)
	defer s.Stop()

	time.Sleep(50 * time.Millisecond)
}

func TestFrequency_Duration(t *testing.T) {
	if Hourly.Duration() != time.Hour {
		t.Fatalf("hourly: expected %v, got %v", time.Hour, Hourly.Duration())
	}
	if Daily.Duration() != 24*time.Hour {
		t.Fatalf("daily: expected %v, got %v", 24*time.Hour, Daily.Duration())
	}
	if Weekly.Duration() != 7*24*time.Hour {
		t.Fatalf("weekly: expected %v, got %v", 7*24*time.Hour, Weekly.Duration())
	}
}
