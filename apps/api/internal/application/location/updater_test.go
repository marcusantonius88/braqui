package location

import (
	"context"
	"sync"
	"testing"

	"github.com/marcusantonius88/braqui/apps/api/internal/domain"
)

type mockPetRepo struct {
	mu   sync.Mutex
	pets map[string][]*domain.Pet
	byID map[string]*domain.Pet
}

func newMockPetRepo() *mockPetRepo {
	return &mockPetRepo{
		pets: make(map[string][]*domain.Pet),
		byID: make(map[string]*domain.Pet),
	}
}

func (r *mockPetRepo) FindByUserID(_ context.Context, userID string) ([]*domain.Pet, error) {
	r.mu.Lock()
	defer r.mu.Unlock()
	return r.pets[userID], nil
}

func (r *mockPetRepo) UpdateLocation(_ context.Context, petID, city string) error {
	r.mu.Lock()
	defer r.mu.Unlock()
	pet, ok := r.byID[petID]
	if !ok {
		return domain.ErrNotFound
	}
	pet.Location = city
	return nil
}

type mockStateManager struct {
	mu         sync.Mutex
	states     map[string]*domain.ConversationState
}

func newMockStateManager() *mockStateManager {
	return &mockStateManager{states: make(map[string]*domain.ConversationState)}
}

func (m *mockStateManager) Start(_ context.Context, userID, flow, step string, _ any) (*domain.ConversationState, error) {
	m.mu.Lock()
	defer m.mu.Unlock()
	s := &domain.ConversationState{UserID: userID, Flow: flow, Step: step}
	m.states[userID] = s
	return s, nil
}

func (m *mockStateManager) Get(_ context.Context, userID string) (*domain.ConversationState, error) {
	m.mu.Lock()
	defer m.mu.Unlock()
	s, ok := m.states[userID]
	if !ok {
		return nil, nil
	}
	return s, nil
}

func (m *mockStateManager) Advance(_ context.Context, userID, nextStep string, _ any) (*domain.ConversationState, error) {
	m.mu.Lock()
	defer m.mu.Unlock()
	if s, ok := m.states[userID]; ok {
		s.Step = nextStep
		return s, nil
	}
	return nil, nil
}

func (m *mockStateManager) Complete(_ context.Context, userID string) error {
	m.mu.Lock()
	defer m.mu.Unlock()
	delete(m.states, userID)
	return nil
}

type mockLogger struct{}

func (mockLogger) Info(_ string, _ map[string]any)  {}
func (mockLogger) Error(_ string, _ map[string]any) {}

func TestLocationUpdater_NoPet(t *testing.T) {
	pets := newMockPetRepo()
	states := newMockStateManager()
	u := NewUpdater(pets, states, mockLogger{})

	reply, err := u.Process(context.Background(), "user-none", "/mudarcidade")
	if err != nil {
		t.Fatalf("process: %v", err)
	}
	if reply != "Você precisa cadastrar um pet antes de configurar a cidade." {
		t.Fatalf("expected no-pet message, got %s", reply)
	}
}

func TestLocationUpdater_StartsFlow(t *testing.T) {
	pet := &domain.Pet{ID: "pet-1", UserID: "user-1", Name: "Thor", Location: "João Pessoa"}
	pets := newMockPetRepo()
	pets.pets["user-1"] = []*domain.Pet{pet}
	pets.byID["pet-1"] = pet
	states := newMockStateManager()
	u := NewUpdater(pets, states, mockLogger{})

	reply, err := u.Process(context.Background(), "user-1", "/mudarcidade")
	if err != nil {
		t.Fatalf("process: %v", err)
	}
	expected := "Cidade atual: João Pessoa. Para qual cidade deseja mudar?"
	if reply != expected {
		t.Fatalf("expected:\n%s\ngot:\n%s", expected, reply)
	}
}

func TestLocationUpdater_UpdateCity(t *testing.T) {
	pet := &domain.Pet{ID: "pet-2", UserID: "user-2", Name: "Rex", Location: "São Paulo"}
	pets := newMockPetRepo()
	pets.pets["user-2"] = []*domain.Pet{pet}
	pets.byID["pet-2"] = pet
	states := newMockStateManager()
	u := NewUpdater(pets, states, mockLogger{})

	u.Process(context.Background(), "user-2", "/mudarcidade")
	reply, err := u.Process(context.Background(), "user-2", "Recife")
	if err != nil {
		t.Fatalf("process city: %v", err)
	}
	if reply != "Cidade atualizada para Recife! 🐶" {
		t.Fatalf("expected confirmation, got %s", reply)
	}
	if pet.Location != "Recife" {
		t.Fatalf("expected Recife, got %s", pet.Location)
	}
}

func TestLocationUpdater_EmptyCity(t *testing.T) {
	pet := &domain.Pet{ID: "pet-3", UserID: "user-3", Name: "Bolinha"}
	pets := newMockPetRepo()
	pets.pets["user-3"] = []*domain.Pet{pet}
	pets.byID["pet-3"] = pet
	states := newMockStateManager()
	u := NewUpdater(pets, states, mockLogger{})

	u.Process(context.Background(), "user-3", "/mudarcidade")
	reply, err := u.Process(context.Background(), "user-3", "")
	if err != nil {
		t.Fatalf("process empty: %v", err)
	}
	if reply != "Por favor, digite o nome de uma cidade." {
		t.Fatalf("expected empty city message, got %s", reply)
	}
}
