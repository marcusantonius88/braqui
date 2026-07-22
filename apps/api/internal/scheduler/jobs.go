package scheduler

import (
	"context"
	"fmt"
	"time"

	"github.com/marcusantonius88/braqui/apps/api/internal/domain"
)

type ReminderRepository interface {
	FindPendingDueBefore(ctx context.Context, due time.Time) ([]*domain.Reminder, error)
	UpdateStatus(ctx context.Context, id, status string) error
}

type PetRepository interface {
	FindByID(ctx context.Context, id string) (*domain.Pet, error)
}

type UserRepository interface {
	FindByID(ctx context.Context, id string) (*domain.User, error)
}

type TelegramGateway interface {
	SendMessage(ctx context.Context, chatID int64, text string) error
}

type ReminderJob struct {
	reminders ReminderRepository
	pets      PetRepository
	users     UserRepository
	tg        TelegramGateway
	log       Logger
}

func NewReminderJob(reminders ReminderRepository, pets PetRepository, users UserRepository, tg TelegramGateway, log Logger) *ReminderJob {
	return &ReminderJob{reminders: reminders, pets: pets, users: users, tg: tg, log: log}
}

func (j *ReminderJob) Name() string { return "reminder" }

func (j *ReminderJob) Execute(ctx context.Context) error {
	now := time.Now()
	due, err := j.reminders.FindPendingDueBefore(ctx, now)
	if err != nil {
		return fmt.Errorf("find pending: %w", err)
	}

	for _, r := range due {
		pet, err := j.pets.FindByID(ctx, r.PetID)
		if err != nil {
			j.log.Error("reminder: find pet", map[string]any{"reminder_id": r.ID, "error": err.Error()})
			continue
		}

		user, err := j.users.FindByID(ctx, pet.UserID)
		if err != nil {
			j.log.Error("reminder: find user", map[string]any{"reminder_id": r.ID, "error": err.Error()})
			continue
		}

		msg := fmt.Sprintf("🐶 Lembrete do %s\n\nHoje é dia de %s.", pet.Name, r.Title)
		if err := j.tg.SendMessage(ctx, user.TelegramID, msg); err != nil {
			j.log.Error("reminder: send failed", map[string]any{"reminder_id": r.ID, "error": err.Error()})
			continue
		}

		if err := j.reminders.UpdateStatus(ctx, r.ID, "completed"); err != nil {
			j.log.Error("reminder: update status", map[string]any{"reminder_id": r.ID, "error": err.Error()})
		}

		j.log.Info("reminder sent", map[string]any{"reminder_id": r.ID, "title": r.Title, "pet": pet.Name})
	}

	return nil
}

type WeeklySummaryJob struct{}

func (j *WeeklySummaryJob) Name() string { return "weekly_summary" }
func (j *WeeklySummaryJob) Execute(_ context.Context) error {
	return nil
}

type ClimateService interface {
	CheckAndAlert(ctx context.Context) error
}

type ClimateAlertJob struct {
	svc ClimateService
	log Logger
}

func NewClimateAlertJob(svc ClimateService, log Logger) *ClimateAlertJob {
	return &ClimateAlertJob{svc: svc, log: log}
}

func (j *ClimateAlertJob) Name() string { return "climate_alert" }

func (j *ClimateAlertJob) Execute(ctx context.Context) error {
	j.log.Info("climate alert job started", nil)
	if err := j.svc.CheckAndAlert(ctx); err != nil {
		return fmt.Errorf("climate alert: %w", err)
	}
	j.log.Info("climate alert job finished", nil)
	return nil
}
