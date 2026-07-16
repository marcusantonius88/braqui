package repositories

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/marcusantonius88/braqui/apps/api/internal/domain"
)

type ConversationStateRepository struct {
	pool *pgxpool.Pool
}

func NewConversationStateRepository(pool *pgxpool.Pool) *ConversationStateRepository {
	return &ConversationStateRepository{pool: pool}
}

func (r *ConversationStateRepository) Create(ctx context.Context, state *domain.ConversationState) error {
	return r.pool.QueryRow(ctx,
		`INSERT INTO conversation_states (user_id, current_flow, current_step, payload)
		 VALUES ($1, $2, $3, $4)
		 RETURNING id, created_at, updated_at`,
		state.UserID, state.Flow, state.Step, state.Payload,
	).Scan(&state.ID, &state.CreatedAt, &state.UpdatedAt)
}

func (r *ConversationStateRepository) FindByUserID(ctx context.Context, userID string) (*domain.ConversationState, error) {
	state := &domain.ConversationState{}
	err := r.pool.QueryRow(ctx,
		`SELECT id, user_id, current_flow, current_step, payload, created_at, updated_at
		 FROM conversation_states WHERE user_id = $1`, userID,
	).Scan(&state.ID, &state.UserID, &state.Flow, &state.Step, &state.Payload, &state.CreatedAt, &state.UpdatedAt)
	if err != nil {
		if err == pgx.ErrNoRows {
			return nil, domain.ErrNotFound
		}
		return nil, fmt.Errorf("find conversation state by user id: %w", err)
	}
	return state, nil
}

func (r *ConversationStateRepository) Update(ctx context.Context, state *domain.ConversationState) error {
	tag, err := r.pool.Exec(ctx,
		`UPDATE conversation_states SET current_flow = $1, current_step = $2, payload = $3, updated_at = NOW()
		 WHERE id = $4`,
		state.Flow, state.Step, state.Payload, state.ID)
	if err != nil {
		return fmt.Errorf("update conversation state: %w", err)
	}
	if tag.RowsAffected() == 0 {
		return domain.ErrNotFound
	}
	return nil
}
