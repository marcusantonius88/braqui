package domain

import (
	"context"
	"time"
)

type Event struct {
	ID          string
	PetID       string
	Type        string
	Description string
	Source      string
	Timestamp   time.Time
	CreatedAt   time.Time
}

type EventRepository interface {
	Create(ctx context.Context, event *Event) error
	FindByID(ctx context.Context, id string) (*Event, error)
	FindByPetID(ctx context.Context, petID string) ([]*Event, error)
}
