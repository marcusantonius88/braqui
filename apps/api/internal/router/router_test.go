package router

import (
	"context"
	"sync"
	"testing"

	"github.com/marcusantonius88/braqui/apps/api/internal/domain"
)

type mockStateManager struct {
	mu     sync.Mutex
	states map[string]*domain.ConversationState
}

func newMockStateManager() *mockStateManager {
	return &mockStateManager{states: make(map[string]*domain.ConversationState)}
}

func (m *mockStateManager) Get(ctx context.Context, userID string) (*domain.ConversationState, error) {
	m.mu.Lock()
	defer m.mu.Unlock()
	s, ok := m.states[userID]
	if !ok {
		return nil, nil
	}
	return s, nil
}

type mockFlowHandler struct {
	reply string
	err   error
}

func (h *mockFlowHandler) Handle(ctx context.Context, userID, text string) (string, error) {
	return h.reply, h.err
}

type mockCommandHandler struct {
	reply string
	err   error
}

func (h *mockCommandHandler) Handle(ctx context.Context, userID string) (string, error) {
	return h.reply, h.err
}

type mockLogger struct{}

func (mockLogger) Info(msg string, fields map[string]any)  {}
func (mockLogger) Error(msg string, fields map[string]any) {}

func TestRouter_RoutesToFlow(t *testing.T) {
	states := newMockStateManager()
	r := NewRouter(states, mockLogger{})

	flow := &mockFlowHandler{reply: "flow response"}
	r.RegisterFlow("register_pet", flow)

	states.states["user-1"] = &domain.ConversationState{UserID: "user-1", Flow: "register_pet", Step: "ask_name"}

	reply, err := r.Route(context.Background(), "user-1", "Thor")
	if err != nil {
		t.Fatalf("route: %v", err)
	}
	if reply != "flow response" {
		t.Fatalf("expected flow response, got %s", reply)
	}
}

func TestRouter_RoutesToCommand(t *testing.T) {
	states := newMockStateManager()
	r := NewRouter(states, mockLogger{})

	cmd := &mockCommandHandler{reply: "command response"}
	r.RegisterCommand("/start", cmd)

	reply, err := r.Route(context.Background(), "user-1", "/start")
	if err != nil {
		t.Fatalf("route: %v", err)
	}
	if reply != "command response" {
		t.Fatalf("expected command response, got %s", reply)
	}
}

func TestRouter_StateHasPriority(t *testing.T) {
	states := newMockStateManager()
	r := NewRouter(states, mockLogger{})

	flow := &mockFlowHandler{reply: "flow response"}
	cmd := &mockCommandHandler{reply: "command response"}
	r.RegisterFlow("register_pet", flow)
	r.RegisterCommand("/start", cmd)

	states.states["user-1"] = &domain.ConversationState{UserID: "user-1", Flow: "register_pet", Step: "ask_name"}

	reply, err := r.Route(context.Background(), "user-1", "/start")
	if err != nil {
		t.Fatalf("route: %v", err)
	}
	if reply != "flow response" {
		t.Fatalf("expected flow response (state has priority), got %s", reply)
	}
}

func TestRouter_Fallback(t *testing.T) {
	states := newMockStateManager()
	r := NewRouter(states, mockLogger{})

	fallback := &mockFlowHandler{reply: "fallback response"}
	r.SetDefault(fallback)

	reply, err := r.Route(context.Background(), "user-1", "random text")
	if err != nil {
		t.Fatalf("route: %v", err)
	}
	if reply != "fallback response" {
		t.Fatalf("expected fallback response, got %s", reply)
	}
}

func TestRouter_EmptyReply(t *testing.T) {
	states := newMockStateManager()
	r := NewRouter(states, mockLogger{})

	reply, err := r.Route(context.Background(), "user-1", "random text")
	if err != nil {
		t.Fatalf("route: %v", err)
	}
	if reply != "" {
		t.Fatalf("expected empty reply, got %s", reply)
	}
}

func TestRouter_UnknownFlow(t *testing.T) {
	states := newMockStateManager()
	r := NewRouter(states, mockLogger{})

	cmd := &mockCommandHandler{reply: "helps"}
	r.RegisterCommand("/help", cmd)

	states.states["user-1"] = &domain.ConversationState{UserID: "user-1", Flow: "unknown_flow", Step: "step1"}

	reply, err := r.Route(context.Background(), "user-1", "/help")
	if err != nil {
		t.Fatalf("route: %v", err)
	}
	if reply != "helps" {
		t.Fatalf("expected command reply for unknown flow, got %s", reply)
	}
}
