package insights

import (
	"context"
	"testing"
	"time"

	"github.com/marcusantonius88/braqui/apps/api/internal/domain"
)

type mockPetRepo struct {
	pets []*domain.Pet
}

func (r *mockPetRepo) FindByUserID(_ context.Context, _ string) ([]*domain.Pet, error) {
	return r.pets, nil
}

type mockEventRepo struct {
	events []*domain.Event
}

func (r *mockEventRepo) FindByPetID(_ context.Context, _ string) ([]*domain.Event, error) {
	return r.events, nil
}

type mockInsightLogger struct{}

func (mockInsightLogger) Info(_ string, _ map[string]any)  {}
func (mockInsightLogger) Error(_ string, _ map[string]any) {}

func recent(daysAgo int) time.Time {
	return time.Now().AddDate(0, 0, -daysAgo)
}

func TestService_NoPet(t *testing.T) {
	svc := NewService(&mockPetRepo{}, &mockEventRepo{}, mockInsightLogger{})
	msg, err := svc.Generate(context.Background(), "user-1")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if msg != "Você precisa cadastrar um pet antes de usar insights." {
		t.Fatalf("expected no-pet message, got %s", msg)
	}
}

func TestService_NoEvents(t *testing.T) {
	svc := NewService(
		&mockPetRepo{pets: []*domain.Pet{{ID: "pet-1", Name: "Thor"}}},
		&mockEventRepo{},
		mockInsightLogger{},
	)
	msg, err := svc.Generate(context.Background(), "user-1")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if msg != "Ainda não há dados suficientes para gerar insights." {
		t.Fatalf("expected no-data message, got %s", msg)
	}
}

func TestService_NoInsights(t *testing.T) {
	svc := NewService(
		&mockPetRepo{pets: []*domain.Pet{{ID: "pet-1", Name: "Thor"}}},
		&mockEventRepo{events: []*domain.Event{
			{Type: "walk", Timestamp: recent(5)},
			{Type: "medication_given", Timestamp: recent(5)},
		}},
		mockInsightLogger{},
	)
	msg, err := svc.Generate(context.Background(), "user-1")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if msg != "Não encontrei padrões relevantes no momento." {
		t.Fatalf("expected no-insights message, got %s", msg)
	}
}

func TestService_VomitInsight(t *testing.T) {
	svc := NewService(
		&mockPetRepo{pets: []*domain.Pet{{ID: "pet-1", Name: "Rex"}}},
		&mockEventRepo{events: []*domain.Event{
			{Type: "vomit", Timestamp: recent(5)},
			{Type: "vomit", Timestamp: recent(10)},
			{Type: "vomit", Timestamp: recent(15)},
			{Type: "medication_given", Timestamp: recent(5)},
		}},
		mockInsightLogger{},
	)
	msg, err := svc.Generate(context.Background(), "user-1")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	expected := "📊 Insights do Rex\n\n• Rex apresentou 3 episódios de vômito nos últimos 30 dias."
	if msg != expected {
		t.Fatalf("expected:\n%s\ngot:\n%s", expected, msg)
	}
}

func TestService_ItchingInsight(t *testing.T) {
	svc := NewService(
		&mockPetRepo{pets: []*domain.Pet{{ID: "pet-1", Name: "Thor"}}},
		&mockEventRepo{events: []*domain.Event{
			{Type: "itching", Timestamp: recent(2)},
			{Type: "itching", Timestamp: recent(5)},
			{Type: "itching", Timestamp: recent(8)},
			{Type: "medication_given", Timestamp: recent(5)},
		}},
		mockInsightLogger{},
	)
	msg, err := svc.Generate(context.Background(), "user-1")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	expected := "📊 Insights do Thor\n\n• Thor apresentou 3 episódios de coceira nos últimos 30 dias."
	if msg != expected {
		t.Fatalf("expected:\n%s\ngot:\n%s", expected, msg)
	}
}

func TestService_PantingInsight(t *testing.T) {
	events := make([]*domain.Event, 6)
	for i := range 5 {
		events[i] = &domain.Event{Type: "panting", Timestamp: recent(i * 3)}
	}
	events[5] = &domain.Event{Type: "medication_given", Timestamp: recent(1)}
	svc := NewService(
		&mockPetRepo{pets: []*domain.Pet{{ID: "pet-1", Name: "Luna"}}},
		&mockEventRepo{events: events},
		mockInsightLogger{},
	)
	msg, err := svc.Generate(context.Background(), "user-1")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	expected := "📊 Insights do Luna\n\n• Luna apresentou 5 registros de ofegância nos últimos 15 dias."
	if msg != expected {
		t.Fatalf("expected:\n%s\ngot:\n%s", expected, msg)
	}
}

func TestService_MedicationAbsence(t *testing.T) {
	svc := NewService(
		&mockPetRepo{pets: []*domain.Pet{{ID: "pet-1", Name: "Bolinha"}}},
		&mockEventRepo{events: []*domain.Event{
			{Type: "walk", Timestamp: recent(1)},
			{Type: "vomit", Timestamp: recent(3)},
		}},
		mockInsightLogger{},
	)
	msg, err := svc.Generate(context.Background(), "user-1")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	expected := "📊 Insights do Bolinha\n\n• Não encontrei registros recentes de medicação para Bolinha."
	if msg != expected {
		t.Fatalf("expected:\n%s\ngot:\n%s", expected, msg)
	}
}

func TestService_MultipleInsights(t *testing.T) {
	svc := NewService(
		&mockPetRepo{pets: []*domain.Pet{{ID: "pet-1", Name: "Thor"}}},
		&mockEventRepo{events: []*domain.Event{
			{Type: "vomit", Timestamp: recent(5)},
			{Type: "vomit", Timestamp: recent(10)},
			{Type: "vomit", Timestamp: recent(15)},
			{Type: "itching", Timestamp: recent(2)},
			{Type: "itching", Timestamp: recent(7)},
			{Type: "itching", Timestamp: recent(12)},
			{Type: "medication_given", Timestamp: recent(3)},
		}},
		mockInsightLogger{},
	)
	msg, err := svc.Generate(context.Background(), "user-1")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	expected := "📊 Insights do Thor\n\n• Thor apresentou 3 episódios de vômito nos últimos 30 dias.\n\n• Thor apresentou 3 episódios de coceira nos últimos 30 dias."
	if msg != expected {
		t.Fatalf("expected:\n%s\ngot:\n%s", expected, msg)
	}
}

func TestService_OldEventsNoInsight(t *testing.T) {
	svc := NewService(
		&mockPetRepo{pets: []*domain.Pet{{ID: "pet-1", Name: "Thor"}}},
		&mockEventRepo{events: []*domain.Event{
			{Type: "vomit", Timestamp: recent(60)},
			{Type: "vomit", Timestamp: recent(90)},
			{Type: "vomit", Timestamp: recent(120)},
			{Type: "medication_given", Timestamp: recent(5)},
		}},
		mockInsightLogger{},
	)
	msg, err := svc.Generate(context.Background(), "user-1")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if msg != "Não encontrei padrões relevantes no momento." {
		t.Fatalf("expected no-insights for old events, got %s", msg)
	}
}

func TestInsight_NoMedicationButHasRecent(t *testing.T) {
	svc := NewService(
		&mockPetRepo{pets: []*domain.Pet{{ID: "pet-1", Name: "Thor"}}},
		&mockEventRepo{events: []*domain.Event{
			{Type: "medication_given", Timestamp: recent(5)},
		}},
		mockInsightLogger{},
	)
	msg, err := svc.Generate(context.Background(), "user-1")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if msg != "Não encontrei padrões relevantes no momento." {
		t.Fatalf("expected no-insights when medication exists, got %s", msg)
	}
}
