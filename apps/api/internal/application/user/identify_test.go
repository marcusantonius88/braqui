package user

import (
	"context"
	"fmt"
	"sync"
	"testing"

	"github.com/marcusantonius88/braqui/apps/api/internal/domain"
)

type mockRepo struct {
	mu     sync.Mutex
	users  map[int64]*domain.User
	nextID int
}

func newMockRepo() *mockRepo {
	return &mockRepo{users: make(map[int64]*domain.User)}
}

func (r *mockRepo) FindByTelegramID(ctx context.Context, telegramID int64) (*domain.User, error) {
	r.mu.Lock()
	defer r.mu.Unlock()
	u, ok := r.users[telegramID]
	if !ok {
		return nil, domain.ErrNotFound
	}
	return u, nil
}

func (r *mockRepo) Create(ctx context.Context, user *domain.User) error {
	r.mu.Lock()
	defer r.mu.Unlock()
	r.nextID++
	user.ID = fmt.Sprintf("user-%d", r.nextID)
	r.users[user.TelegramID] = user
	return nil
}

func TestIdentify_ExistingUser(t *testing.T) {
	repo := newMockRepo()
	ctx := context.Background()

	repo.Create(ctx, &domain.User{TelegramID: 123, FirstName: "João"})

	out, err := Identify(ctx, repo, 123, "João", "")
	if err != nil {
		t.Fatalf("identify: %v", err)
	}
	if out.IsNew {
		t.Fatal("expected existing user")
	}
	if out.User.TelegramID != 123 {
		t.Fatalf("expected 123, got %d", out.User.TelegramID)
	}
}

func TestIdentify_NewUser(t *testing.T) {
	repo := newMockRepo()
	ctx := context.Background()

	out, err := Identify(ctx, repo, 456, "Maria", "maria_bot")
	if err != nil {
		t.Fatalf("identify: %v", err)
	}
	if !out.IsNew {
		t.Fatal("expected new user")
	}
	if out.User.FirstName != "Maria" {
		t.Fatalf("expected Maria, got %s", out.User.FirstName)
	}
	if out.User.Username != "maria_bot" {
		t.Fatalf("expected maria_bot, got %s", out.User.Username)
	}
	if out.User.ID == "" {
		t.Fatal("expected generated ID")
	}
}

func TestIdentify_NewUserWithoutUsername(t *testing.T) {
	repo := newMockRepo()
	ctx := context.Background()

	out, err := Identify(ctx, repo, 789, "Pedro", "")
	if err != nil {
		t.Fatalf("identify: %v", err)
	}
	if !out.IsNew {
		t.Fatal("expected new user")
	}
	if out.User.Username != "" {
		t.Fatalf("expected empty username, got %s", out.User.Username)
	}
}
