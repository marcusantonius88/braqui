package domain

import (
	"context"
	"time"
)

const (
	ReminderStatusPending   = "pending"
	ReminderStatusCompleted = "completed"
	ReminderStatusCancelled = "cancelled"
)

type Reminder struct {
	ID          string
	PetID       string
	Title       string
	Description string
	DueDate     time.Time
	Status      string
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

type ReminderRepository interface {
	Create(ctx context.Context, reminder *Reminder) error
	FindByID(ctx context.Context, id string) (*Reminder, error)
	FindByPetID(ctx context.Context, petID string) ([]*Reminder, error)
	FindPendingDueBefore(ctx context.Context, due time.Time) ([]*Reminder, error)
	UpdateStatus(ctx context.Context, id, status string) error
}
