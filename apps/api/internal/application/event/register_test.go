package event

import (
	"context"
	"sync"
	"testing"

	"github.com/marcusantonius88/braqui/apps/api/internal/domain"
	domainai "github.com/marcusantonius88/braqui/apps/api/internal/domain/ai"
)

type mockPetRepo struct {
	mu   sync.Mutex
	pets map[string][]*domain.Pet
}

func newMockPetRepo() *mockPetRepo {
	return &mockPetRepo{pets: make(map[string][]*domain.Pet)}
}

func (r *mockPetRepo) FindByUserID(ctx context.Context, userID string) ([]*domain.Pet, error) {
	r.mu.Lock()
	defer r.mu.Unlock()
	return r.pets[userID], nil
}

type mockEventRepo struct {
	mu     sync.Mutex
	events []*domain.Event
}

func newMockEventRepo() *mockEventRepo {
	return &mockEventRepo{}
}

func (r *mockEventRepo) Create(ctx context.Context, event *domain.Event) error {
	r.mu.Lock()
	defer r.mu.Unlock()
	event.ID = "evt-1"
	r.events = append(r.events, event)
	return nil
}

func (r *mockEventRepo) FindByPetID(ctx context.Context, petID string) ([]*domain.Event, error) {
	r.mu.Lock()
	defer r.mu.Unlock()
	return r.events, nil
}

type mockAI struct {
	result *domainai.InterpretationResult
	err    error
}

func (m *mockAI) Interpret(ctx context.Context, message string) (*domainai.InterpretationResult, error) {
	return m.result, m.err
}

type mockLogger struct{}

func (mockLogger) Info(msg string, fields map[string]any)  {}
func (mockLogger) Error(msg string, fields map[string]any) {}

func TestRegister_NoPet(t *testing.T) {
	pets := newMockPetRepo()
	events := newMockEventRepo()
	ai := &mockAI{}
	r := NewRegisterer(pets, events, ai, mockLogger{})

	result, err := r.Register(context.Background(), "user-1", "Thor vomitou")
	if err != nil {
		t.Fatalf("register: %v", err)
	}
	if result != ResultNoPet {
		t.Fatalf("expected no_pet, got %s", result)
	}
}

func TestRegister_ParserMatch(t *testing.T) {
	pets := newMockPetRepo()
	pets.pets["user-1"] = []*domain.Pet{{ID: "pet-1", Name: "Thor"}}
	events := newMockEventRepo()
	ai := &mockAI{}
	r := NewRegisterer(pets, events, ai, mockLogger{})

	result, err := r.Register(context.Background(), "user-1", "Thor vomitou")
	if err != nil {
		t.Fatalf("register: %v", err)
	}
	if result != ResultRegistered {
		t.Fatalf("expected registered, got %s", result)
	}
	if len(events.events) != 1 {
		t.Fatalf("expected 1 event, got %d", len(events.events))
	}
	if events.events[0].Type != "vomit" {
		t.Fatalf("expected vomit, got %s", events.events[0].Type)
	}
	if events.events[0].Source != "parser" {
		t.Fatalf("expected parser source, got %s", events.events[0].Source)
	}
	if events.events[0].Description != "Thor vomitou" {
		t.Fatalf("expected original message, got %s", events.events[0].Description)
	}
}

func TestRegister_AIFallback(t *testing.T) {
	pets := newMockPetRepo()
	pets.pets["user-2"] = []*domain.Pet{{ID: "pet-2", Name: "Rex"}}
	events := newMockEventRepo()
	ai := &mockAI{
		result: &domainai.InterpretationResult{Type: "fatigue", Confidence: "high", Payload: map[string]any{}},
	}
	r := NewRegisterer(pets, events, ai, mockLogger{})

	result, err := r.Register(context.Background(), "user-2", "Rex parece estranho hoje")
	if err != nil {
		t.Fatalf("register: %v", err)
	}
	if result != ResultRegistered {
		t.Fatalf("expected registered, got %s", result)
	}
	if len(events.events) != 1 {
		t.Fatalf("expected 1 event, got %d", len(events.events))
	}
	if events.events[0].Type != "fatigue" {
		t.Fatalf("expected fatigue, got %s", events.events[0].Type)
	}
	if events.events[0].Source != "ai" {
		t.Fatalf("expected ai source, got %s", events.events[0].Source)
	}
}

func TestRegister_AINotInterpreted(t *testing.T) {
	pets := newMockPetRepo()
	pets.pets["user-3"] = []*domain.Pet{{ID: "pet-3", Name: "Bolinha"}}
	events := newMockEventRepo()
	ai := &mockAI{
		result: &domainai.InterpretationResult{Type: "NOT_INTERPRETED", Confidence: "low", Payload: map[string]any{}},
	}
	r := NewRegisterer(pets, events, ai, mockLogger{})

	result, err := r.Register(context.Background(), "user-3", "mensagem aleatória")
	if err != nil {
		t.Fatalf("register: %v", err)
	}
	if result != ResultNotInterpreted {
		t.Fatalf("expected not_interpreted, got %s", result)
	}
	if len(events.events) != 0 {
		t.Fatalf("expected 0 events, got %d", len(events.events))
	}
}

func TestRegister_ParserMatchSaveDescription(t *testing.T) {
	pets := newMockPetRepo()
	pets.pets["user-4"] = []*domain.Pet{{ID: "pet-4", Name: "Luna"}}
	events := newMockEventRepo()
	ai := &mockAI{}
	r := NewRegisterer(pets, events, ai, mockLogger{})

	msg := "Luna vomitou a ração hoje"
	result, err := r.Register(context.Background(), "user-4", msg)
	if err != nil {
		t.Fatalf("register: %v", err)
	}
	if result != ResultRegistered {
		t.Fatalf("expected registered, got %s", result)
	}
	if events.events[0].Description != msg {
		t.Fatalf("expected original description, got %s", events.events[0].Description)
	}
}

func TestRegister_Diarrhea(t *testing.T) {
	pets := newMockPetRepo()
	pets.pets["user-5"] = []*domain.Pet{{ID: "pet-5", Name: "Toto"}}
	events := newMockEventRepo()
	ai := &mockAI{}
	r := NewRegisterer(pets, events, ai, mockLogger{})

	result, err := r.Register(context.Background(), "user-5", "Toto com diarreia")
	if err != nil {
		t.Fatalf("register: %v", err)
	}
	if result != ResultRegistered {
		t.Fatalf("expected registered, got %s", result)
	}
	if events.events[0].Type != "diarrhea" {
		t.Fatalf("expected diarrhea, got %s", events.events[0].Type)
	}
}
