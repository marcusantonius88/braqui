package domain

import (
	"context"
	"time"
)

type ConversationState struct {
	ID        string
	UserID    string
	Flow      string
	Step      string
	Payload   []byte
	CreatedAt time.Time
	UpdatedAt time.Time
}

type ConversationStateRepository interface {
	Create(ctx context.Context, state *ConversationState) error
	FindByUserID(ctx context.Context, userID string) (*ConversationState, error)
	Update(ctx context.Context, state *ConversationState) error
}
