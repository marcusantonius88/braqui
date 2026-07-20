package timeline

import (
	"context"
	"fmt"
	"sync"
	"testing"
	"time"

	"github.com/marcusantonius88/braqui/apps/api/internal/domain"
)

type mockPetRepo struct {
	mu   sync.Mutex
	pets map[string][]*domain.Pet
}

func newMockPetRepo() *mockPetRepo {
	return &mockPetRepo{pets: make(map[string][]*domain.Pet)}
}

func (r *mockPetRepo) FindByUserID(_ context.Context, userID string) ([]*domain.Pet, error) {
	r.mu.Lock()
	defer r.mu.Unlock()
	return r.pets[userID], nil
}

type mockEventRepo struct {
	mu     sync.Mutex
	events map[string][]*domain.Event
}

func newMockEventRepo() *mockEventRepo {
	return &mockEventRepo{events: make(map[string][]*domain.Event)}
}

func (r *mockEventRepo) FindByPetID(_ context.Context, petID string) ([]*domain.Event, error) {
	r.mu.Lock()
	defer r.mu.Unlock()
	return r.events[petID], nil
}

type mockLogger struct{}

func (mockLogger) Info(_ string, _ map[string]any)  {}
func (mockLogger) Error(_ string, _ map[string]any) {}

func TestTimeline_NoPet(t *testing.T) {
	pets := newMockPetRepo()
	events := newMockEventRepo()
	svc := NewService(pets, events, mockLogger{})

	reply, err := svc.GetTimeline(context.Background(), "user-none")
	if err != nil {
		t.Fatalf("get timeline: %v", err)
	}
	if reply != "Você ainda não cadastrou nenhum pet. Envie uma mensagem para começar!" {
		t.Fatalf("unexpected reply: %s", reply)
	}
}

func TestTimeline_NoEvents(t *testing.T) {
	pets := newMockPetRepo()
	pets.pets["user-1"] = []*domain.Pet{{ID: "pet-1", Name: "Thor"}}
	events := newMockEventRepo()
	events.events["pet-1"] = []*domain.Event{}
	svc := NewService(pets, events, mockLogger{})

	reply, err := svc.GetTimeline(context.Background(), "user-1")
	if err != nil {
		t.Fatalf("get timeline: %v", err)
	}
	expected := "Ainda não encontrei eventos registrados para o Thor 🐶"
	if reply != expected {
		t.Fatalf("expected:\n%s\ngot:\n%s", expected, reply)
	}
}

func TestTimeline_WithEvents(t *testing.T) {
	now := time.Now()
	pets := newMockPetRepo()
	pets.pets["user-2"] = []*domain.Pet{{ID: "pet-2", Name: "Rex"}}
	events := newMockEventRepo()
	events.events["pet-2"] = []*domain.Event{
		{Type: "vomit", Timestamp: now.Add(-3 * 24 * time.Hour)},
		{Type: "panting", Timestamp: now.Add(-24 * time.Hour)},
		{Type: "medication_given", Timestamp: now},
	}
	svc := NewService(pets, events, mockLogger{})

	reply, err := svc.GetTimeline(context.Background(), "user-2")
	if err != nil {
		t.Fatalf("get timeline: %v", err)
	}

	expected := "📋 Histórico recente do Rex\n\n• Hoje - Medicação\n• Ontem - Ofegante\n• 3 dias atrás - Vômito"
	if reply != expected {
		t.Fatalf("expected:\n%s\ngot:\n%s", expected, reply)
	}
}

func TestTimeline_Limit10(t *testing.T) {
	now := time.Now()
	pets := newMockPetRepo()
	pets.pets["user-3"] = []*domain.Pet{{ID: "pet-3", Name: "Bolinha"}}

	evts := make([]*domain.Event, 15)
	for i := 0; i < 15; i++ {
		evts[i] = &domain.Event{
			Type:      "vomit",
			Timestamp: now.Add(-time.Duration(i) * 24 * time.Hour),
		}
	}
	events := newMockEventRepo()
	events.events["pet-3"] = evts
	svc := NewService(pets, events, mockLogger{})

	reply, err := svc.GetTimeline(context.Background(), "user-3")
	if err != nil {
		t.Fatalf("get timeline: %v", err)
	}

	expected := "📋 Histórico recente do Bolinha\n\n"
	for i := 0; i < 10; i++ {
		line := ""
		switch {
		case i == 0:
			line = "• Hoje - Vômito"
		case i == 1:
			line = "• Ontem - Vômito"
		default:
			line = fmt.Sprintf("• %d dias atrás - Vômito", i)
		}
		if i < 9 {
			expected += line + "\n"
		} else {
			expected += line
		}
	}

	if reply != expected {
		t.Fatalf("expected:\n%s\n\ngot:\n%s", expected, reply)
	}
}

func TestTimeline_Labels(t *testing.T) {
	now := time.Now()
	pets := newMockPetRepo()
	pets.pets["user-4"] = []*domain.Pet{{ID: "pet-4", Name: "Luna"}}
	events := newMockEventRepo()
	events.events["pet-4"] = []*domain.Event{
		{Type: "itching", Timestamp: now},
		{Type: "fatigue", Timestamp: now},
		{Type: "cough", Timestamp: now},
		{Type: "diarrhea", Timestamp: now},
		{Type: "vet_visit", Timestamp: now},
	}
	svc := NewService(pets, events, mockLogger{})

	reply, err := svc.GetTimeline(context.Background(), "user-4")
	if err != nil {
		t.Fatalf("get timeline: %v", err)
	}

	for _, label := range []string{"Coceira", "Cansaço", "Tosse", "Diarreia", "Consulta veterinária"} {
		if !contains(reply, label) {
			t.Fatalf("expected label %s in reply", label)
		}
	}
}

func TestTimeline_Ordering(t *testing.T) {
	now := time.Now()
	pets := newMockPetRepo()
	pets.pets["user-5"] = []*domain.Pet{{ID: "pet-5", Name: "Toto"}}
	events := newMockEventRepo()
	events.events["pet-5"] = []*domain.Event{
		{Type: "vomit", Timestamp: now.Add(-5 * 24 * time.Hour)},
		{Type: "panting", Timestamp: now.Add(-1 * 24 * time.Hour)},
		{Type: "cough", Timestamp: now.Add(-3 * 24 * time.Hour)},
	}
	svc := NewService(pets, events, mockLogger{})

	reply, err := svc.GetTimeline(context.Background(), "user-5")
	if err != nil {
		t.Fatalf("get timeline: %v", err)
	}

	if !contains(reply, "Ontem - Ofegante") {
		t.Fatal("expected 'Ontem - Ofegante' first")
	}
	if !contains(reply, "3 dias atrás - Tosse") {
		t.Fatal("expected '3 dias atrás - Tosse' second")
	}
	if !contains(reply, "5 dias atrás - Vômito") {
		t.Fatal("expected '5 dias atrás - Vômito' last")
	}
}

func contains(s, substr string) bool {
	return len(s) >= len(substr) && searchString(s, substr)
}

func searchString(s, substr string) bool {
	for i := 0; i <= len(s)-len(substr); i++ {
		if s[i:i+len(substr)] == substr {
			return true
		}
	}
	return false
}
