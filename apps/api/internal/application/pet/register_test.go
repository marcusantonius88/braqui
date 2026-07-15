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
	state.ID = "state-1"
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

func TestOnboarder_HasPet(t *testing.T) {
	petRepo := newMockPetRepo()
	stateRepo := newMockStateRepo()
	o := NewOnboarder(petRepo, stateRepo)
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
	stateRepo := newMockStateRepo()
	o := NewOnboarder(petRepo, stateRepo)
	ctx := context.Background()

	reply, err := o.Process(ctx, "user-2", "/start")
	if err != nil {
		t.Fatalf("process: %v", err)
	}
	if reply != "Qual o nome do seu cão?" {
		t.Fatalf("expected name question, got: %s", reply)
	}

	state, _ := stateRepo.FindByUserID(ctx, "user-2")
	if state == nil || state.State != "onboarding_name" {
		t.Fatalf("expected state onboarding_name, got %v", state)
	}
}

func TestOnboarder_FullFlow(t *testing.T) {
	petRepo := newMockPetRepo()
	stateRepo := newMockStateRepo()
	o := NewOnboarder(petRepo, stateRepo)
	ctx := context.Background()

	stateRepo.Create(ctx, &domain.ConversationState{
		UserID: "user-3",
		State:  "onboarding_name",
		Data:   []byte("{}"),
	})

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
	stateRepo := newMockStateRepo()
	o := NewOnboarder(petRepo, stateRepo)
	ctx := context.Background()

	stateRepo.Create(ctx, &domain.ConversationState{
		UserID: "user-4",
		State:  "onboarding_name",
		Data:   []byte("{}"),
	})

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
	stateRepo := newMockStateRepo()
	o := NewOnboarder(petRepo, stateRepo)
	ctx := context.Background()

	data, _ := json.Marshal(onboardingData{Name: "Rex", Breed: "SRD"})
	stateRepo.Create(ctx, &domain.ConversationState{
		UserID: "user-5",
		State:  "onboarding_age",
		Data:   data,
	})

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
	stateRepo := newMockStateRepo()
	o := NewOnboarder(petRepo, stateRepo)
	ctx := context.Background()

	data, _ := json.Marshal(onboardingData{Name: "Rex", Breed: "SRD", Age: 3})
	stateRepo.Create(ctx, &domain.ConversationState{
		UserID: "user-6",
		State:  "onboarding_weight",
		Data:   data,
	})

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
