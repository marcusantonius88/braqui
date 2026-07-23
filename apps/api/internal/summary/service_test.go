package summary

import (
	"context"
	"sync"
	"testing"
	"time"

	"github.com/marcusantonius88/braqui/apps/api/internal/domain"
)

type mockSummaryPetRepo struct {
	mu   sync.Mutex
	pets []*domain.Pet
}

func (r *mockSummaryPetRepo) FindAll(_ context.Context) ([]*domain.Pet, error) {
	r.mu.Lock()
	defer r.mu.Unlock()
	return r.pets, nil
}

type mockSummaryEventRepo struct {
	mu     sync.Mutex
	events []*domain.Event
}

func (r *mockSummaryEventRepo) FindByPetID(_ context.Context, _ string) ([]*domain.Event, error) {
	r.mu.Lock()
	defer r.mu.Unlock()
	return r.events, nil
}

type mockSummaryUserRepo struct {
	mu    sync.Mutex
	users map[string]*domain.User
}

func newMockSummaryUserRepo() *mockSummaryUserRepo {
	return &mockSummaryUserRepo{users: make(map[string]*domain.User)}
}

func (r *mockSummaryUserRepo) FindByID(_ context.Context, id string) (*domain.User, error) {
	r.mu.Lock()
	defer r.mu.Unlock()
	u, ok := r.users[id]
	if !ok {
		return nil, domain.ErrNotFound
	}
	return u, nil
}

type mockSummaryTelegram struct {
	mu       sync.Mutex
	messages []string
}

func (t *mockSummaryTelegram) SendMessage(_ context.Context, _ int64, text string) error {
	t.mu.Lock()
	defer t.mu.Unlock()
	t.messages = append(t.messages, text)
	return nil
}

type mockSummaryLogger struct{}

func (mockSummaryLogger) Info(_ string, _ map[string]any)  {}
func (mockSummaryLogger) Error(_ string, _ map[string]any) {}

func TestService_NoPets(t *testing.T) {
	pets := &mockSummaryPetRepo{}
	events := &mockSummaryEventRepo{}
	users := newMockSummaryUserRepo()
	tg := &mockSummaryTelegram{}
	svc := NewService(pets, events, users, tg, mockSummaryLogger{})

	if err := svc.GenerateAndSend(context.Background()); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if len(tg.messages) > 0 {
		t.Fatalf("expected no messages, got %d", len(tg.messages))
	}
}

func TestService_NoEvents(t *testing.T) {
	pets := &mockSummaryPetRepo{
		pets: []*domain.Pet{{ID: "pet-1", UserID: "user-1", Name: "Thor"}},
	}
	events := &mockSummaryEventRepo{}
	users := newMockSummaryUserRepo()
	users.users["user-1"] = &domain.User{ID: "user-1", TelegramID: 12345}
	tg := &mockSummaryTelegram{}
	svc := NewService(pets, events, users, tg, mockSummaryLogger{})

	if err := svc.GenerateAndSend(context.Background()); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if len(tg.messages) != 1 {
		t.Fatalf("expected 1 message, got %d", len(tg.messages))
	}
	expected := "Nenhum evento foi registrado esta semana para o Thor."
	if tg.messages[0] != expected {
		t.Fatalf("expected:\n%s\ngot:\n%s", expected, tg.messages[0])
	}
}

func recentDay(daysAgo int) time.Time {
	return time.Now().AddDate(0, 0, -daysAgo)
}

func TestService_WithEvents(t *testing.T) {
	pets := &mockSummaryPetRepo{
		pets: []*domain.Pet{{ID: "pet-1", UserID: "user-1", Name: "Rex"}},
	}
	events := &mockSummaryEventRepo{events: []*domain.Event{
		{Type: "vomit", Timestamp: recentDay(1)},
		{Type: "vomit", Timestamp: recentDay(2)},
		{Type: "itching", Timestamp: recentDay(3)},
		{Type: "panting", Timestamp: recentDay(4)},
		{Type: "panting", Timestamp: recentDay(5)},
		{Type: "panting", Timestamp: recentDay(6)},
		{Type: "medication_given", Timestamp: recentDay(6)},
	}}
	users := newMockSummaryUserRepo()
	users.users["user-1"] = &domain.User{ID: "user-1", TelegramID: 67890}
	tg := &mockSummaryTelegram{}
	svc := NewService(pets, events, users, tg, mockSummaryLogger{})

	if err := svc.GenerateAndSend(context.Background()); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if len(tg.messages) != 1 {
		t.Fatalf("expected 1 message, got %d", len(tg.messages))
	}
	expected := "📊 Resumo semanal do Rex\n\n• 7 eventos registrados\n\n• 2 episódios de vômito\n• 1 episódio de coceira\n• 3 episódios de ofegância\n• 1 medicação registrada\n\nContinue registrando eventos para que eu possa acompanhar melhor a saúde do Rex 🐶"
	if tg.messages[0] != expected {
		t.Fatalf("expected:\n%s\ngot:\n%s", expected, tg.messages[0])
	}
}

func TestService_MultiplePets(t *testing.T) {
	pets := &mockSummaryPetRepo{
		pets: []*domain.Pet{
			{ID: "pet-1", UserID: "user-1", Name: "Thor"},
			{ID: "pet-2", UserID: "user-2", Name: "Luna"},
		},
	}
	events := &mockSummaryEventRepo{events: []*domain.Event{}}
	users := newMockSummaryUserRepo()
	users.users["user-1"] = &domain.User{ID: "user-1", TelegramID: 111}
	users.users["user-2"] = &domain.User{ID: "user-2", TelegramID: 222}
	tg := &mockSummaryTelegram{}
	svc := NewService(pets, events, users, tg, mockSummaryLogger{})

	if err := svc.GenerateAndSend(context.Background()); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if len(tg.messages) != 2 {
		t.Fatalf("expected 2 messages, got %d", len(tg.messages))
	}
}

func TestService_OldEventsIgnored(t *testing.T) {
	pets := &mockSummaryPetRepo{
		pets: []*domain.Pet{{ID: "pet-1", UserID: "user-1", Name: "Thor"}},
	}
	events := &mockSummaryEventRepo{events: []*domain.Event{
		{Type: "vomit", Timestamp: recentDay(10)},
		{Type: "vomit", Timestamp: recentDay(15)},
	}}
	users := newMockSummaryUserRepo()
	users.users["user-1"] = &domain.User{ID: "user-1", TelegramID: 333}
	tg := &mockSummaryTelegram{}
	svc := NewService(pets, events, users, tg, mockSummaryLogger{})

	if err := svc.GenerateAndSend(context.Background()); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if len(tg.messages) != 1 {
		t.Fatalf("expected 1 message, got %d", len(tg.messages))
	}
	expected := "Nenhum evento foi registrado esta semana para o Thor."
	if tg.messages[0] != expected {
		t.Fatalf("expected no-event message, got %s", tg.messages[0])
	}
}

func TestFormat_SingularPlural(t *testing.T) {
	pets := &mockSummaryPetRepo{
		pets: []*domain.Pet{{ID: "pet-1", UserID: "user-1", Name: "Thor"}},
	}
	events := &mockSummaryEventRepo{events: []*domain.Event{
		{Type: "vomit", Timestamp: recentDay(1)},
	}}
	users := newMockSummaryUserRepo()
	users.users["user-1"] = &domain.User{ID: "user-1", TelegramID: 444}
	tg := &mockSummaryTelegram{}
	svc := NewService(pets, events, users, tg, mockSummaryLogger{})

	if err := svc.GenerateAndSend(context.Background()); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	expected := "📊 Resumo semanal do Thor\n\n• 1 evento registrado\n\n• 1 episódio de vômito\n\nContinue registrando eventos para que eu possa acompanhar melhor a saúde do Thor 🐶"
	if tg.messages[0] != expected {
		t.Fatalf("expected:\n%s\ngot:\n%s", expected, tg.messages[0])
	}
}
