package domain

import (
	"context"
	"time"
)

type Reminder struct {
	ID             string
	PetID          string
	Type           string
	Description    string
	DueDate        time.Time
	RepeatInterval string
	CreatedAt      time.Time
	UpdatedAt      time.Time
}

type ReminderRepository interface {
	Create(ctx context.Context, reminder *Reminder) error
	FindByID(ctx context.Context, id string) (*Reminder, error)
	FindByPetID(ctx context.Context, petID string) ([]*Reminder, error)
}
