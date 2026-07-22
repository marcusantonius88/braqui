package repositories

import (
	"context"
	"log"
	"os"
	"testing"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/marcusantonius88/braqui/apps/api/internal/domain"
)

func getTestPool(t *testing.T) *pgxpool.Pool {
	t.Helper()
	dsn := os.Getenv("DATABASE_URL")
	if dsn == "" {
		dsn = "postgres://postgres:postgres@localhost:5432/braqui?sslmode=disable"
	}
	pool, err := pgxpool.New(context.Background(), dsn)
	if err != nil {
		t.Fatalf("connect: %v", err)
	}
	t.Cleanup(pool.Close)
	return pool
}

func truncateAll(t *testing.T, pool *pgxpool.Pool) {
	t.Helper()
	ctx := context.Background()
	tables := []string{"conversation_states", "reminders", "events", "pets", "users"}
	for _, table := range tables {
		if _, err := pool.Exec(ctx, "DELETE FROM "+table); err != nil {
			log.Printf("truncate %s: %v", table, err)
		}
	}
}

func TestUserRepository_CRUD(t *testing.T) {
	pool := getTestPool(t)
	truncateAll(t, pool)
	repo := NewUserRepository(pool)
	ctx := context.Background()

	user := &domain.User{TelegramID: 12345, FirstName: "João"}
	if err := repo.Create(ctx, user); err != nil {
		t.Fatalf("create: %v", err)
	}
	if user.ID == "" {
		t.Fatal("expected id to be set")
	}

	found, err := repo.FindByID(ctx, user.ID)
	if err != nil {
		t.Fatalf("find by id: %v", err)
	}
	if found.FirstName != "João" {
		t.Fatalf("expected João, got %s", found.FirstName)
	}

	found, err = repo.FindByTelegramID(ctx, 12345)
	if err != nil {
		t.Fatalf("find by telegram id: %v", err)
	}
	if found.FirstName != "João" {
		t.Fatalf("expected João, got %s", found.FirstName)
	}
}

func TestUserRepository_NotFound(t *testing.T) {
	pool := getTestPool(t)
	truncateAll(t, pool)
	repo := NewUserRepository(pool)
	ctx := context.Background()

	_, err := repo.FindByID(ctx, "00000000-0000-0000-0000-000000000000")
	if err != domain.ErrNotFound {
		t.Fatalf("expected ErrNotFound, got %v", err)
	}

	_, err = repo.FindByTelegramID(ctx, 99999)
	if err != domain.ErrNotFound {
		t.Fatalf("expected ErrNotFound, got %v", err)
	}
}

func TestPetRepository_CRUD(t *testing.T) {
	pool := getTestPool(t)
	truncateAll(t, pool)

	userRepo := NewUserRepository(pool)
	petRepo := NewPetRepository(pool)
	ctx := context.Background()

	user := &domain.User{TelegramID: 1, FirstName: "Maria"}
	if err := userRepo.Create(ctx, user); err != nil {
		t.Fatalf("create user: %v", err)
	}

	pet := &domain.Pet{UserID: user.ID, Name: "Rex", Breed: "Bulldog", Age: 3, Weight: 22.5, Location: "São Paulo"}
	if err := petRepo.Create(ctx, pet); err != nil {
		t.Fatalf("create pet: %v", err)
	}
	if pet.ID == "" {
		t.Fatal("expected id to be set")
	}

	found, err := petRepo.FindByID(ctx, pet.ID)
	if err != nil {
		t.Fatalf("find by id: %v", err)
	}
	if found.Name != "Rex" {
		t.Fatalf("expected Rex, got %s", found.Name)
	}

	pets, err := petRepo.FindByUserID(ctx, user.ID)
	if err != nil {
		t.Fatalf("find by user id: %v", err)
	}
	if len(pets) != 1 {
		t.Fatalf("expected 1 pet, got %d", len(pets))
	}
}

func TestPetRepository_NotFound(t *testing.T) {
	pool := getTestPool(t)
	truncateAll(t, pool)
	repo := NewPetRepository(pool)
	ctx := context.Background()

	_, err := repo.FindByID(ctx, "00000000-0000-0000-0000-000000000000")
	if err != domain.ErrNotFound {
		t.Fatalf("expected ErrNotFound, got %v", err)
	}
}

func TestEventRepository_CRUD(t *testing.T) {
	pool := getTestPool(t)
	truncateAll(t, pool)

	userRepo := NewUserRepository(pool)
	petRepo := NewPetRepository(pool)
	eventRepo := NewEventRepository(pool)
	ctx := context.Background()

	user := &domain.User{TelegramID: 2, FirstName: "Pedro"}
	if err := userRepo.Create(ctx, user); err != nil {
		t.Fatalf("create user: %v", err)
	}
	pet := &domain.Pet{UserID: user.ID, Name: "Toto", Breed: "SRD", Age: 5, Weight: 10, Location: "RJ"}
	if err := petRepo.Create(ctx, pet); err != nil {
		t.Fatalf("create pet: %v", err)
	}

	event := &domain.Event{PetID: pet.ID, Type: "feeding", Description: "comeu ração", Source: "manual", Timestamp: time.Now()}
	if err := eventRepo.Create(ctx, event); err != nil {
		t.Fatalf("create event: %v", err)
	}
	if event.ID == "" {
		t.Fatal("expected id to be set")
	}

	found, err := eventRepo.FindByID(ctx, event.ID)
	if err != nil {
		t.Fatalf("find by id: %v", err)
	}
	if found.Type != "feeding" {
		t.Fatalf("expected feeding, got %s", found.Type)
	}

	events, err := eventRepo.FindByPetID(ctx, pet.ID)
	if err != nil {
		t.Fatalf("find by pet id: %v", err)
	}
	if len(events) != 1 {
		t.Fatalf("expected 1 event, got %d", len(events))
	}
}

func TestEventRepository_NotFound(t *testing.T) {
	pool := getTestPool(t)
	truncateAll(t, pool)
	repo := NewEventRepository(pool)
	ctx := context.Background()

	_, err := repo.FindByID(ctx, "00000000-0000-0000-0000-000000000000")
	if err != domain.ErrNotFound {
		t.Fatalf("expected ErrNotFound, got %v", err)
	}
}

func TestReminderRepository_CRUD(t *testing.T) {
	pool := getTestPool(t)
	truncateAll(t, pool)

	userRepo := NewUserRepository(pool)
	petRepo := NewPetRepository(pool)
	reminderRepo := NewReminderRepository(pool)
	ctx := context.Background()

	user := &domain.User{TelegramID: 3, FirstName: "Ana"}
	if err := userRepo.Create(ctx, user); err != nil {
		t.Fatalf("create user: %v", err)
	}
	pet := &domain.Pet{UserID: user.ID, Name: "Luna", Breed: "Poodle", Age: 2, Weight: 5, Location: "BH"}
	if err := petRepo.Create(ctx, pet); err != nil {
		t.Fatalf("create pet: %v", err)
	}

	dueDate, _ := time.Parse(time.RFC3339, "2026-07-15T10:00:00Z")
	reminder := &domain.Reminder{PetID: pet.ID, Title: "Dar Simparic", Description: "antipulga", DueDate: dueDate, Status: domain.ReminderStatusPending}
	if err := reminderRepo.Create(ctx, reminder); err != nil {
		t.Fatalf("create reminder: %v", err)
	}
	if reminder.ID == "" {
		t.Fatal("expected id to be set")
	}

	found, err := reminderRepo.FindByID(ctx, reminder.ID)
	if err != nil {
		t.Fatalf("find by id: %v", err)
	}
	if found.Title != "Dar Simparic" {
		t.Fatalf("expected Dar Simparic, got %s", found.Title)
	}
	if found.Status != domain.ReminderStatusPending {
		t.Fatalf("expected pending, got %s", found.Status)
	}

	reminders, err := reminderRepo.FindByPetID(ctx, pet.ID)
	if err != nil {
		t.Fatalf("find by pet id: %v", err)
	}
	if len(reminders) != 1 {
		t.Fatalf("expected 1 reminder, got %d", len(reminders))
	}
}

func TestReminderRepository_NotFound(t *testing.T) {
	pool := getTestPool(t)
	truncateAll(t, pool)
	repo := NewReminderRepository(pool)
	ctx := context.Background()

	_, err := repo.FindByID(ctx, "00000000-0000-0000-0000-000000000000")
	if err != domain.ErrNotFound {
		t.Fatalf("expected ErrNotFound, got %v", err)
	}
}

func TestConversationStateRepository_CRUD(t *testing.T) {
	pool := getTestPool(t)
	truncateAll(t, pool)

	userRepo := NewUserRepository(pool)
	convRepo := NewConversationStateRepository(pool)
	ctx := context.Background()

	user := &domain.User{TelegramID: 4, FirstName: "Carlos"}
	if err := userRepo.Create(ctx, user); err != nil {
		t.Fatalf("create user: %v", err)
	}

	state := &domain.ConversationState{UserID: user.ID, Flow: "register_pet", Step: "ask_name", Payload: []byte(`{}`)}
	if err := convRepo.Create(ctx, state); err != nil {
		t.Fatalf("create state: %v", err)
	}
	if state.ID == "" {
		t.Fatal("expected id to be set")
	}

	found, err := convRepo.FindByUserID(ctx, user.ID)
	if err != nil {
		t.Fatalf("find by user id: %v", err)
	}
	if found.Flow != "register_pet" {
		t.Fatalf("expected register_pet, got %s", found.Flow)
	}
	if found.Step != "ask_name" {
		t.Fatalf("expected ask_name, got %s", found.Step)
	}

	found.Step = "idle"
	if err := convRepo.Update(ctx, found); err != nil {
		t.Fatalf("update state: %v", err)
	}

	updated, err := convRepo.FindByUserID(ctx, user.ID)
	if err != nil {
		t.Fatalf("find by user id after update: %v", err)
	}
	if updated.Step != "idle" {
		t.Fatalf("expected idle, got %s", updated.Step)
	}
}

func TestConversationStateRepository_NotFound(t *testing.T) {
	pool := getTestPool(t)
	truncateAll(t, pool)
	repo := NewConversationStateRepository(pool)
	ctx := context.Background()

	state := &domain.ConversationState{ID: "00000000-0000-0000-0000-000000000000", Payload: []byte(`{}`)}
	err := repo.Update(ctx, state)
	if err != domain.ErrNotFound {
		t.Fatalf("expected ErrNotFound, got %v", err)
	}
}
