package reminder

import (
	"context"
	"encoding/json"
	"sync"
	"testing"

	"github.com/marcusantonius88/braqui/apps/api/internal/domain"
)

type mockPetRepoCreator struct {
	mu   sync.Mutex
	pets map[string][]*domain.Pet
}

func (r *mockPetRepoCreator) FindByUserID(_ context.Context, userID string) ([]*domain.Pet, error) {
	r.mu.Lock()
	defer r.mu.Unlock()
	return r.pets[userID], nil
}

func (r *mockPetRepoCreator) FindByID(_ context.Context, id string) (*domain.Pet, error) {
	return nil, nil
}

type mockRemindRepo struct {
	mu       sync.Mutex
	reminder *domain.Reminder
}

func (r *mockRemindRepo) Create(_ context.Context, reminder *domain.Reminder) error {
	r.mu.Lock()
	defer r.mu.Unlock()
	reminder.ID = "rem-1"
	r.reminder = reminder
	return nil
}

type mockStateManagerCreator struct {
	mu         sync.Mutex
	states     map[string]*domain.ConversationState
	advanceTo  string
	completeFn func()
}

func newMockStateManagerCreator() *mockStateManagerCreator {
	return &mockStateManagerCreator{states: make(map[string]*domain.ConversationState)}
}

func (m *mockStateManagerCreator) Start(_ context.Context, userID, flow, step string, payload any) (*domain.ConversationState, error) {
	m.mu.Lock()
	defer m.mu.Unlock()
	raw, _ := json.Marshal(payload)
	s := &domain.ConversationState{UserID: userID, Flow: flow, Step: step, Payload: raw}
	m.states[userID] = s
	return s, nil
}

func (m *mockStateManagerCreator) Get(_ context.Context, userID string) (*domain.ConversationState, error) {
	m.mu.Lock()
	defer m.mu.Unlock()
	s, ok := m.states[userID]
	if !ok {
		return nil, nil
	}
	return s, nil
}

func (m *mockStateManagerCreator) Advance(_ context.Context, userID, nextStep string, payload any) (*domain.ConversationState, error) {
	m.mu.Lock()
	defer m.mu.Unlock()
	s, ok := m.states[userID]
	if !ok {
		return nil, nil
	}
	raw, _ := json.Marshal(payload)
	s.Step = nextStep
	s.Payload = raw
	m.advanceTo = nextStep
	return s, nil
}

func (m *mockStateManagerCreator) Complete(_ context.Context, userID string) error {
	m.mu.Lock()
	defer m.mu.Unlock()
	if m.completeFn != nil {
		m.completeFn()
	}
	delete(m.states, userID)
	return nil
}

func TestCreator_StartFlow(t *testing.T) {
	pets := &mockPetRepoCreator{pets: map[string][]*domain.Pet{"user-1": {{ID: "pet-1", Name: "Thor"}}}}
	reminds := &mockRemindRepo{}
	states := newMockStateManagerCreator()
	log := mockLogger{}

	c := NewCreator(pets, reminds, states, log)
	reply, err := c.Process(context.Background(), "user-1", "/remind")
	if err != nil {
		t.Fatalf("process: %v", err)
	}
	if reply != "O que você quer lembrar?" {
		t.Fatalf("expected prompt, got %s", reply)
	}
}

func TestCreator_NoPet(t *testing.T) {
	pets := &mockPetRepoCreator{pets: map[string][]*domain.Pet{}}
	reminds := &mockRemindRepo{}
	states := newMockStateManagerCreator()
	log := mockLogger{}

	c := NewCreator(pets, reminds, states, log)
	reply, err := c.Process(context.Background(), "user-nopet", "/remind")
	if err != nil {
		t.Fatalf("process: %v", err)
	}
	if reply != "Você precisa cadastrar um pet antes de criar lembretes." {
		t.Fatalf("expected no-pet message, got %s", reply)
	}
}

func TestCreator_FullFlow(t *testing.T) {
	pets := &mockPetRepoCreator{pets: map[string][]*domain.Pet{"user-2": {{ID: "pet-2", Name: "Rex"}}}}
	reminds := &mockRemindRepo{}
	states := newMockStateManagerCreator()
	log := mockLogger{}

	c := NewCreator(pets, reminds, states, log)

	reply, err := c.Process(context.Background(), "user-2", "/remind")
	if err != nil {
		t.Fatalf("process start: %v", err)
	}
	if reply != "O que você quer lembrar?" {
		t.Fatalf("expected ask_title, got %s", reply)
	}

	reply, err = c.Process(context.Background(), "user-2", "Dar Simparic")
	if err != nil {
		t.Fatalf("process title: %v", err)
	}
	if reply != "Para quando?" {
		t.Fatalf("expected ask_date, got %s", reply)
	}

	reply, err = c.Process(context.Background(), "user-2", "amanhã")
	if err != nil {
		t.Fatalf("process date: %v", err)
	}
	if reply != "Lembrete criado! Vou te lembrar de \"dar simparic\" amanhã." {
		t.Fatalf("expected confirmation, got %s", reply)
	}

	if reminds.reminder == nil {
		t.Fatal("expected reminder to be created")
	}
	if reminds.reminder.Title != "Dar Simparic" {
		t.Fatalf("expected title 'Dar Simparic', got %s", reminds.reminder.Title)
	}
}

func TestCreator_InvalidDate(t *testing.T) {
	pets := &mockPetRepoCreator{pets: map[string][]*domain.Pet{"user-3": {{ID: "pet-3", Name: "Bolinha"}}}}
	reminds := &mockRemindRepo{}
	states := newMockStateManagerCreator()
	log := mockLogger{}

	c := NewCreator(pets, reminds, states, log)

	c.Process(context.Background(), "user-3", "/remind")
	reply, err := c.Process(context.Background(), "user-3", "Dar Simparic")
	if err != nil {
		t.Fatalf("process title: %v", err)
	}
	if reply != "Para quando?" {
		t.Fatalf("expected ask_date, got %s", reply)
	}

	reply, err = c.Process(context.Background(), "user-3", "algum dia")
	if err != nil {
		t.Fatalf("process invalid date: %v", err)
	}
	if reply != "Não consegui entender a data informada. Digite algo como \"amanhã\", \"dia 15\" ou \"daqui a 7 dias\"." {
		t.Fatalf("expected invalid date message, got %s", reply)
	}

	if reminds.reminder != nil {
		t.Fatal("expected no reminder to be created for invalid date")
	}
}

type mockLogger struct{}

func (mockLogger) Info(_ string, _ map[string]any)  {}
func (mockLogger) Error(_ string, _ map[string]any) {}
