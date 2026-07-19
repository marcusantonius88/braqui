package state

import (
	"context"
	"sync"
	"testing"

	"github.com/marcusantonius88/braqui/apps/api/internal/domain"
)

type mockLogger struct{}

func (mockLogger) Info(msg string, fields map[string]any)  {}
func (mockLogger) Error(msg string, fields map[string]any) {}

type mockStateRepo struct {
	mu     sync.Mutex
	states map[string]*domain.ConversationState
}

func newMockStateRepo() *mockStateRepo {
	return &mockStateRepo{states: make(map[string]*domain.ConversationState)}
}

func (r *mockStateRepo) Create(ctx context.Context, state *domain.ConversationState) error {
	r.mu.Lock()
	defer r.mu.Unlock()
	state.ID = "state-" + state.UserID
	r.states[state.UserID] = state
	return nil
}

func (r *mockStateRepo) FindByUserID(ctx context.Context, userID string) (*domain.ConversationState, error) {
	r.mu.Lock()
	defer r.mu.Unlock()
	s, ok := r.states[userID]
	if !ok {
		return nil, domain.ErrNotFound
	}
	return s, nil
}

func (r *mockStateRepo) Update(ctx context.Context, state *domain.ConversationState) error {
	r.mu.Lock()
	defer r.mu.Unlock()
	if _, ok := r.states[state.UserID]; !ok {
		return domain.ErrNotFound
	}
	r.states[state.UserID] = state
	return nil
}

func TestManager_Start(t *testing.T) {
	repo := newMockStateRepo()
	m := NewManager(repo, mockLogger{})
	ctx := context.Background()

	s, err := m.Start(ctx, "user-1", "register_pet", "ask_name", map[string]string{})
	if err != nil {
		t.Fatalf("start: %v", err)
	}
	if s.UserID != "user-1" {
		t.Fatalf("expected user-1, got %s", s.UserID)
	}
	if s.Flow != "register_pet" {
		t.Fatalf("expected register_pet, got %s", s.Flow)
	}
	if s.Step != "ask_name" {
		t.Fatalf("expected ask_name, got %s", s.Step)
	}
}

func TestManager_Get_NotFound(t *testing.T) {
	repo := newMockStateRepo()
	m := NewManager(repo, mockLogger{})
	ctx := context.Background()

	s, err := m.Get(ctx, "nonexistent")
	if err != nil {
		t.Fatalf("get: %v", err)
	}
	if s != nil {
		t.Fatal("expected nil for nonexistent user")
	}
}

func TestManager_Get_Found(t *testing.T) {
	repo := newMockStateRepo()
	m := NewManager(repo, mockLogger{})
	ctx := context.Background()

	m.Start(ctx, "user-2", "register_pet", "ask_name", map[string]string{})

	s, err := m.Get(ctx, "user-2")
	if err != nil {
		t.Fatalf("get: %v", err)
	}
	if s == nil {
		t.Fatal("expected non-nil state")
	}
	if s.Step != "ask_name" {
		t.Fatalf("expected ask_name, got %s", s.Step)
	}
}

func TestManager_Advance(t *testing.T) {
	repo := newMockStateRepo()
	m := NewManager(repo, mockLogger{})
	ctx := context.Background()

	m.Start(ctx, "user-3", "register_pet", "ask_name", map[string]string{"name": "Thor"})

	s, err := m.Advance(ctx, "user-3", "ask_breed", map[string]string{"name": "Thor"})
	if err != nil {
		t.Fatalf("advance: %v", err)
	}
	if s.Step != "ask_breed" {
		t.Fatalf("expected ask_breed, got %s", s.Step)
	}
}

func TestManager_Advance_NotFound(t *testing.T) {
	repo := newMockStateRepo()
	m := NewManager(repo, mockLogger{})
	ctx := context.Background()

	_, err := m.Advance(ctx, "nonexistent", "step2", map[string]string{})
	if err == nil {
		t.Fatal("expected error for nonexistent user")
	}
}

func TestManager_Complete(t *testing.T) {
	repo := newMockStateRepo()
	m := NewManager(repo, mockLogger{})
	ctx := context.Background()

	m.Start(ctx, "user-4", "register_pet", "ask_name", map[string]string{})

	if err := m.Complete(ctx, "user-4"); err != nil {
		t.Fatalf("complete: %v", err)
	}

	s, _ := repo.FindByUserID(ctx, "user-4")
	if s.Step != "" {
		t.Fatalf("expected empty step after complete, got %s", s.Step)
	}
	if s.Flow != "" {
		t.Fatalf("expected empty flow after complete, got %s", s.Flow)
	}
}

func TestManager_Complete_NotFound(t *testing.T) {
	repo := newMockStateRepo()
	m := NewManager(repo, mockLogger{})
	ctx := context.Background()

	err := m.Complete(ctx, "nonexistent")
	if err == nil {
		t.Fatal("expected error for nonexistent user")
	}
}
