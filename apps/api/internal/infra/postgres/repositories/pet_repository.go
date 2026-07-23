package repositories

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/marcusantonius88/braqui/apps/api/internal/domain"
)

type PetRepository struct {
	pool *pgxpool.Pool
}

func NewPetRepository(pool *pgxpool.Pool) *PetRepository {
	return &PetRepository{pool: pool}
}

func (r *PetRepository) Create(ctx context.Context, pet *domain.Pet) error {
	return r.pool.QueryRow(ctx,
		`INSERT INTO pets (user_id, name, breed, age, weight, location)
		 VALUES ($1, $2, $3, $4, $5, $6)
		 RETURNING id, created_at, updated_at`,
		pet.UserID, pet.Name, pet.Breed, pet.Age, pet.Weight, pet.Location,
	).Scan(&pet.ID, &pet.CreatedAt, &pet.UpdatedAt)
}

func (r *PetRepository) FindByID(ctx context.Context, id string) (*domain.Pet, error) {
	pet := &domain.Pet{}
	err := r.pool.QueryRow(ctx,
		`SELECT id, user_id, name, breed, age, weight, location, created_at, updated_at
		 FROM pets WHERE id = $1`, id,
	).Scan(&pet.ID, &pet.UserID, &pet.Name, &pet.Breed, &pet.Age, &pet.Weight, &pet.Location, &pet.CreatedAt, &pet.UpdatedAt)
	if err != nil {
		if err == pgx.ErrNoRows {
			return nil, domain.ErrNotFound
		}
		return nil, fmt.Errorf("find pet by id: %w", err)
	}
	return pet, nil
}

func (r *PetRepository) FindByUserID(ctx context.Context, userID string) ([]*domain.Pet, error) {
	rows, err := r.pool.Query(ctx,
		`SELECT id, user_id, name, breed, age, weight, location, created_at, updated_at
		 FROM pets WHERE user_id = $1 ORDER BY created_at`, userID)
	if err != nil {
		return nil, fmt.Errorf("find pets by user id: %w", err)
	}
	defer rows.Close()

	var pets []*domain.Pet
	for rows.Next() {
		pet := &domain.Pet{}
		if err := rows.Scan(&pet.ID, &pet.UserID, &pet.Name, &pet.Breed, &pet.Age, &pet.Weight, &pet.Location, &pet.CreatedAt, &pet.UpdatedAt); err != nil {
			return nil, fmt.Errorf("scan pet: %w", err)
		}
		pets = append(pets, pet)
	}
	if pets == nil {
		return []*domain.Pet{}, nil
	}
	return pets, nil
}

func (r *PetRepository) FindAll(ctx context.Context) ([]*domain.Pet, error) {
	rows, err := r.pool.Query(ctx,
		`SELECT id, user_id, name, breed, age, weight, location, created_at, updated_at
		 FROM pets ORDER BY created_at`)
	if err != nil {
		return nil, fmt.Errorf("find all pets: %w", err)
	}
	defer rows.Close()

	var pets []*domain.Pet
	for rows.Next() {
		pet := &domain.Pet{}
		if err := rows.Scan(&pet.ID, &pet.UserID, &pet.Name, &pet.Breed, &pet.Age, &pet.Weight, &pet.Location, &pet.CreatedAt, &pet.UpdatedAt); err != nil {
			return nil, fmt.Errorf("scan pet: %w", err)
		}
		pets = append(pets, pet)
	}
	if pets == nil {
		return []*domain.Pet{}, nil
	}
	return pets, nil
}

func (r *PetRepository) FindAllWithLocation(ctx context.Context) ([]*domain.Pet, error) {
	rows, err := r.pool.Query(ctx,
		`SELECT id, user_id, name, breed, age, weight, location, created_at, updated_at
		 FROM pets WHERE location IS NOT NULL AND location != '' ORDER BY created_at`)
	if err != nil {
		return nil, fmt.Errorf("find all with location: %w", err)
	}
	defer rows.Close()

	var pets []*domain.Pet
	for rows.Next() {
		pet := &domain.Pet{}
		if err := rows.Scan(&pet.ID, &pet.UserID, &pet.Name, &pet.Breed, &pet.Age, &pet.Weight, &pet.Location, &pet.CreatedAt, &pet.UpdatedAt); err != nil {
			return nil, fmt.Errorf("scan pet: %w", err)
		}
		pets = append(pets, pet)
	}
	if pets == nil {
		return []*domain.Pet{}, nil
	}
	return pets, nil
}

func (r *PetRepository) UpdateLocation(ctx context.Context, petID, city string) error {
	_, err := r.pool.Exec(ctx,
		`UPDATE pets SET location = $1, updated_at = NOW() WHERE id = $2`,
		city, petID)
	if err != nil {
		return fmt.Errorf("update location: %w", err)
	}
	return nil
}
