package repositories

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/marcusantonius88/braqui/apps/api/internal/domain"
)

type ReminderRepository struct {
	pool *pgxpool.Pool
}

func NewReminderRepository(pool *pgxpool.Pool) *ReminderRepository {
	return &ReminderRepository{pool: pool}
}

func (r *ReminderRepository) Create(ctx context.Context, reminder *domain.Reminder) error {
	return r.pool.QueryRow(ctx,
		`INSERT INTO reminders (pet_id, type, description, due_date, repeat_interval)
		 VALUES ($1, $2, $3, $4, $5)
		 RETURNING id, created_at, updated_at`,
		reminder.PetID, reminder.Type, reminder.Description, reminder.DueDate, reminder.RepeatInterval,
	).Scan(&reminder.ID, &reminder.CreatedAt, &reminder.UpdatedAt)
}

func (r *ReminderRepository) FindByID(ctx context.Context, id string) (*domain.Reminder, error) {
	reminder := &domain.Reminder{}
	err := r.pool.QueryRow(ctx,
		`SELECT id, pet_id, type, description, due_date, repeat_interval, created_at, updated_at
		 FROM reminders WHERE id = $1`, id,
	).Scan(&reminder.ID, &reminder.PetID, &reminder.Type, &reminder.Description, &reminder.DueDate, &reminder.RepeatInterval, &reminder.CreatedAt, &reminder.UpdatedAt)
	if err != nil {
		if err == pgx.ErrNoRows {
			return nil, domain.ErrNotFound
		}
		return nil, fmt.Errorf("find reminder by id: %w", err)
	}
	return reminder, nil
}

func (r *ReminderRepository) FindByPetID(ctx context.Context, petID string) ([]*domain.Reminder, error) {
	rows, err := r.pool.Query(ctx,
		`SELECT id, pet_id, type, description, due_date, repeat_interval, created_at, updated_at
		 FROM reminders WHERE pet_id = $1 ORDER BY due_date`, petID)
	if err != nil {
		return nil, fmt.Errorf("find reminders by pet id: %w", err)
	}
	defer rows.Close()

	var reminders []*domain.Reminder
	for rows.Next() {
		reminder := &domain.Reminder{}
		if err := rows.Scan(&reminder.ID, &reminder.PetID, &reminder.Type, &reminder.Description, &reminder.DueDate, &reminder.RepeatInterval, &reminder.CreatedAt, &reminder.UpdatedAt); err != nil {
			return nil, fmt.Errorf("scan reminder: %w", err)
		}
		reminders = append(reminders, reminder)
	}
	if reminders == nil {
		return []*domain.Reminder{}, nil
	}
	return reminders, nil
}
