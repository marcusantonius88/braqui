package domain

import (
	"context"
	"time"
)

type Pet struct {
	ID        string
	UserID    string
	Name      string
	Breed     string
	Age       int
	Weight    float64
	Location  string
	CreatedAt time.Time
	UpdatedAt time.Time
}

type PetRepository interface {
	Create(ctx context.Context, pet *Pet) error
	FindByID(ctx context.Context, id string) (*Pet, error)
	FindByUserID(ctx context.Context, userID string) ([]*Pet, error)
	FindAllWithLocation(ctx context.Context) ([]*Pet, error)
	UpdateLocation(ctx context.Context, petID, city string) error
}
