package pet

import (
	"context"
	"encoding/json"
	"sync"
	"testing"

	"github.com/marcusantonius88/braqui/apps/api/internal/domain"
)

type mockPetRepo struct {
	mu   sync.Mutex
	pets map[string]*domain.Pet
}

func newMockPetRepo() *mockPetRepo {
	return &mockPetRepo{pets: make(map[string]*domain.Pet)}
}

func (r *mockPetRepo) Create(ctx context.Context, pet *domain.Pet) error {
	r.mu.Lock()
	defer r.mu.Unlock()
	pet.ID = "pet-1"
	r.pets[pet.ID] = pet
	return nil
}

func (r *mockPetRepo) FindByUserID(ctx context.Context, userID string) ([]*domain.Pet, error) {
	r.mu.Lock()
	defer r.mu.Unlock()
	var result []*domain.Pet
	for _, p := range r.pets {
		if p.UserID == userID {
			result = append(result, p)
		}
	}
	return result, nil
}

type mockStateManager struct {
	mu     sync.Mutex
	states map[string]*domain.ConversationState
}

func newMockStateManager() *mockStateManager {
	return &mockStateManager{states: make(map[string]*domain.ConversationState)}
}

func (m *mockStateManager) Start(ctx context.Context, userID, flow, step string, payload any) (*domain.ConversationState, error) {
	m.mu.Lock()
	defer m.mu.Unlock()
	raw, _ := json.Marshal(payload)
	state := &domain.ConversationState{
		ID:      "state-" + userID,
		UserID:  userID,
		Flow:    flow,
		Step:    step,
		Payload: raw,
	}
	m.states[userID] = state
	return state, nil
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

func (m *mockStateManager) Advance(ctx context.Context, userID, nextStep string, payload any) (*domain.ConversationState, error) {
	m.mu.Lock()
	defer m.mu.Unlock()
	s, ok := m.states[userID]
	if !ok {
		return nil, domain.ErrNotFound
	}
	raw, _ := json.Marshal(payload)
	s.Step = nextStep
	s.Payload = raw
	return s, nil
}

func (m *mockStateManager) Complete(ctx context.Context, userID string) error {
	m.mu.Lock()
	defer m.mu.Unlock()
	s, ok := m.states[userID]
	if !ok {
		return domain.ErrNotFound
	}
	s.Step = ""
	return nil
}

func TestOnboarder_HasPet(t *testing.T) {
	petRepo := newMockPetRepo()
	states := newMockStateManager()
	o := NewOnboarder(petRepo, states)
	ctx := context.Background()

	petRepo.Create(ctx, &domain.Pet{UserID: "user-1", Name: "Rex"})

	reply, err := o.Process(ctx, "user-1", "anything")
	if err != nil {
		t.Fatalf("process: %v", err)
	}
	if reply != "" {
		t.Fatalf("expected empty reply for user with pet, got: %s", reply)
	}
}

func TestOnboarder_FirstStep(t *testing.T) {
	petRepo := newMockPetRepo()
	states := newMockStateManager()
	o := NewOnboarder(petRepo, states)
	ctx := context.Background()

	reply, err := o.Process(ctx, "user-2", "/start")
	if err != nil {
		t.Fatalf("process: %v", err)
	}
	if reply != "Qual o nome do seu cão?" {
		t.Fatalf("expected name question, got: %s", reply)
	}

	state, _ := states.Get(ctx, "user-2")
	if state == nil || state.Flow != flowRegisterPet || state.Step != stepAskName {
		t.Fatalf("expected flow register_pet / step ask_name, got flow=%s step=%s", state.Flow, state.Step)
	}
}

func TestOnboarder_FullFlow(t *testing.T) {
	petRepo := newMockPetRepo()
	states := newMockStateManager()
	o := NewOnboarder(petRepo, states)
	ctx := context.Background()

	states.Start(ctx, "user-3", flowRegisterPet, stepAskName, onboardingData{})

	steps := []struct {
		input string
		want  string
	}{
		{"Thor", "Qual a raça do Thor?"},
		{"Buldogue Francês", "Qual a idade do Thor? (em anos)"},
		{"3", "Qual o peso do Thor? (em kg)"},
		{"12.5", "Em qual cidade você mora?"},
		{"João Pessoa", "Perfeito"},
	}

	for _, step := range steps {
		reply, err := o.Process(ctx, "user-3", step.input)
		if err != nil {
			t.Fatalf("process(%q): %v", step.input, err)
		}
		if !contains(reply, step.want) {
			t.Fatalf("process(%q): expected to contain %q, got %q", step.input, step.want, reply)
		}
	}

	pets, _ := petRepo.FindByUserID(ctx, "user-3")
	if len(pets) != 1 {
		t.Fatal("expected 1 pet to be created")
	}
	if pets[0].Name != "Thor" {
		t.Fatalf("expected Thor, got %s", pets[0].Name)
	}
	if pets[0].Breed != "Buldogue Francês" {
		t.Fatalf("expected Buldogue Francês, got %s", pets[0].Breed)
	}
	if pets[0].Location != "João Pessoa" {
		t.Fatalf("expected João Pessoa, got %s", pets[0].Location)
	}
}

func TestOnboarder_EmptyName(t *testing.T) {
	petRepo := newMockPetRepo()
	states := newMockStateManager()
	o := NewOnboarder(petRepo, states)
	ctx := context.Background()

	states.Start(ctx, "user-4", flowRegisterPet, stepAskName, onboardingData{})

	reply, err := o.Process(ctx, "user-4", "")
	if err != nil {
		t.Fatalf("process: %v", err)
	}
	if !contains(reply, "Por favor") {
		t.Fatalf("expected validation error, got: %s", reply)
	}
}

func TestOnboarder_InvalidAge(t *testing.T) {
	petRepo := newMockPetRepo()
	states := newMockStateManager()
	o := NewOnboarder(petRepo, states)
	ctx := context.Background()

	states.Start(ctx, "user-5", flowRegisterPet, stepAskAge, onboardingData{Name: "Rex", Breed: "SRD", Age: 3})

	reply, err := o.Process(ctx, "user-5", "abc")
	if err != nil {
		t.Fatalf("process: %v", err)
	}
	if !contains(reply, "Por favor") {
		t.Fatalf("expected validation error, got: %s", reply)
	}
}

func TestOnboarder_InvalidWeight(t *testing.T) {
	petRepo := newMockPetRepo()
	states := newMockStateManager()
	o := NewOnboarder(petRepo, states)
	ctx := context.Background()

	states.Start(ctx, "user-6", flowRegisterPet, stepAskWeight, onboardingData{Name: "Rex", Breed: "SRD", Age: 3, Weight: 12})

	reply, err := o.Process(ctx, "user-6", "-5")
	if err != nil {
		t.Fatalf("process: %v", err)
	}
	if !contains(reply, "Por favor") {
		t.Fatalf("expected validation error, got: %s", reply)
	}
}

func contains(s, substr string) bool {
	return len(s) >= len(substr) && containsStr(s, substr)
}

func containsStr(s, substr string) bool {
	for i := 0; i <= len(s)-len(substr); i++ {
		if s[i:i+len(substr)] == substr {
			return true
		}
	}
	return false
}
