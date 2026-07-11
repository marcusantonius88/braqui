package repositories

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/marcusantonius88/braqui/apps/api/internal/domain"
)

type EventRepository struct {
	pool *pgxpool.Pool
}

func NewEventRepository(pool *pgxpool.Pool) *EventRepository {
	return &EventRepository{pool: pool}
}

func (r *EventRepository) Create(ctx context.Context, event *domain.Event) error {
	return r.pool.QueryRow(ctx,
		`INSERT INTO events (pet_id, type, description, timestamp)
		 VALUES ($1, $2, $3, $4)
		 RETURNING id, created_at`,
		event.PetID, event.Type, event.Description, event.Timestamp,
	).Scan(&event.ID, &event.CreatedAt)
}

func (r *EventRepository) FindByID(ctx context.Context, id string) (*domain.Event, error) {
	event := &domain.Event{}
	err := r.pool.QueryRow(ctx,
		`SELECT id, pet_id, type, description, timestamp, created_at
		 FROM events WHERE id = $1`, id,
	).Scan(&event.ID, &event.PetID, &event.Type, &event.Description, &event.Timestamp, &event.CreatedAt)
	if err != nil {
		if err == pgx.ErrNoRows {
			return nil, domain.ErrNotFound
		}
		return nil, fmt.Errorf("find event by id: %w", err)
	}
	return event, nil
}

func (r *EventRepository) FindByPetID(ctx context.Context, petID string) ([]*domain.Event, error) {
	rows, err := r.pool.Query(ctx,
		`SELECT id, pet_id, type, description, timestamp, created_at
		 FROM events WHERE pet_id = $1 ORDER BY timestamp DESC`, petID)
	if err != nil {
		return nil, fmt.Errorf("find events by pet id: %w", err)
	}
	defer rows.Close()

	var events []*domain.Event
	for rows.Next() {
		event := &domain.Event{}
		if err := rows.Scan(&event.ID, &event.PetID, &event.Type, &event.Description, &event.Timestamp, &event.CreatedAt); err != nil {
			return nil, fmt.Errorf("scan event: %w", err)
		}
		events = append(events, event)
	}
	if events == nil {
		return []*domain.Event{}, nil
	}
	return events, nil
}
