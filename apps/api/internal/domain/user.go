package domain

import (
	"context"
	"time"
)

type User struct {
	ID         string
	TelegramID int64
	FirstName  string
	Username   string
	CreatedAt  time.Time
	UpdatedAt  time.Time
}

type UserRepository interface {
	Create(ctx context.Context, user *User) error
	FindByID(ctx context.Context, id string) (*User, error)
	FindByTelegramID(ctx context.Context, telegramID int64) (*User, error)
}
