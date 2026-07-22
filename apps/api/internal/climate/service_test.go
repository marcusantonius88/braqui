package climate

import (
	"context"
	"sync"
	"testing"

	"github.com/marcusantonius88/braqui/apps/api/internal/domain"
)

type mockPetRepo struct {
	mu   sync.Mutex
	pets []*domain.Pet
}

func (r *mockPetRepo) FindAllWithLocation(_ context.Context) ([]*domain.Pet, error) {
	r.mu.Lock()
	defer r.mu.Unlock()
	return r.pets, nil
}

type mockUserRepo struct {
	mu    sync.Mutex
	users map[string]*domain.User
}

func newMockUserRepo() *mockUserRepo {
	return &mockUserRepo{users: make(map[string]*domain.User)}
}

func (r *mockUserRepo) FindByID(_ context.Context, id string) (*domain.User, error) {
	r.mu.Lock()
	defer r.mu.Unlock()
	u, ok := r.users[id]
	if !ok {
		return nil, domain.ErrNotFound
	}
	return u, nil
}

type mockWeatherProvider struct {
	data *WeatherData
	err  error
}

func (p *mockWeatherProvider) GetCurrentWeather(_ context.Context, _ string) (*WeatherData, error) {
	return p.data, p.err
}

type mockTelegram struct {
	mu       sync.Mutex
	messages []string
}

func (t *mockTelegram) SendMessage(_ context.Context, _ int64, text string) error {
	t.mu.Lock()
	defer t.mu.Unlock()
	t.messages = append(t.messages, text)
	return nil
}

type mockClimateLogger struct{}

func (mockClimateLogger) Info(_ string, _ map[string]any)  {}
func (mockClimateLogger) Error(_ string, _ map[string]any) {}

func TestService_NoPets(t *testing.T) {
	pets := &mockPetRepo{}
	users := newMockUserRepo()
	weather := &mockWeatherProvider{}
	tg := &mockTelegram{}
	svc := NewService(pets, users, weather, tg, mockClimateLogger{})

	if err := svc.CheckAndAlert(context.Background()); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if len(tg.messages) > 0 {
		t.Fatalf("expected no messages, got %d", len(tg.messages))
	}
}

func TestService_NormalWeather(t *testing.T) {
	pets := &mockPetRepo{
		pets: []*domain.Pet{
			{ID: "pet-1", UserID: "user-1", Name: "Rex", Location: "São Paulo"},
		},
	}
	users := newMockUserRepo()
	users.users["user-1"] = &domain.User{ID: "user-1", TelegramID: 12345}
	weather := &mockWeatherProvider{
		data: &WeatherData{Temperature: 25, Humidity: 60, City: "São Paulo"},
	}
	tg := &mockTelegram{}
	svc := NewService(pets, users, weather, tg, mockClimateLogger{})

	if err := svc.CheckAndAlert(context.Background()); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if len(tg.messages) > 0 {
		t.Fatalf("expected no messages for normal weather, got %d", len(tg.messages))
	}
}

func TestService_HighHeatAlert(t *testing.T) {
	pets := &mockPetRepo{
		pets: []*domain.Pet{
			{ID: "pet-2", UserID: "user-2", Name: "Thor", Location: "João Pessoa"},
		},
	}
	users := newMockUserRepo()
	users.users["user-2"] = &domain.User{ID: "user-2", TelegramID: 67890}
	weather := &mockWeatherProvider{
		data: &WeatherData{Temperature: 33, Humidity: 70, City: "João Pessoa"},
	}
	tg := &mockTelegram{}
	svc := NewService(pets, users, weather, tg, mockClimateLogger{})

	if err := svc.CheckAndAlert(context.Background()); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if len(tg.messages) != 1 {
		t.Fatalf("expected 1 message, got %d", len(tg.messages))
	}
	msg := tg.messages[0]
	if msg != "🐶 Atenção\n\nHoje está muito quente em João Pessoa (33°C).\n\nCães braquicefálicos podem sofrer mais nesses dias." {
		t.Fatalf("unexpected message:\n%s", msg)
	}
}

func TestService_CriticalHeatAlert(t *testing.T) {
	pets := &mockPetRepo{
		pets: []*domain.Pet{
			{ID: "pet-3", UserID: "user-3", Name: "Bolinha", Location: "Recife"},
		},
	}
	users := newMockUserRepo()
	users.users["user-3"] = &domain.User{ID: "user-3", TelegramID: 11111}
	weather := &mockWeatherProvider{
		data: &WeatherData{Temperature: 37, Humidity: 75, City: "Recife"},
	}
	tg := &mockTelegram{}
	svc := NewService(pets, users, weather, tg, mockClimateLogger{})

	if err := svc.CheckAndAlert(context.Background()); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if len(tg.messages) != 1 {
		t.Fatalf("expected 1 message, got %d", len(tg.messages))
	}
	msg := tg.messages[0]
	if msg != "🐶 Atenção\n\nHoje está muito quente em Recife (37°C).\n\nCães braquicefálicos podem sofrer mais nesses dias.\n\n⚠️ Calor crítico! Evite passeios e mantenha seu pet em local arejado." {
		t.Fatalf("unexpected message:\n%s", msg)
	}
}

func TestService_AntiSpam(t *testing.T) {
	pets := &mockPetRepo{
		pets: []*domain.Pet{
			{ID: "pet-4", UserID: "user-4", Name: "Luna", Location: "Rio de Janeiro"},
		},
	}
	users := newMockUserRepo()
	users.users["user-4"] = &domain.User{ID: "user-4", TelegramID: 22222}
	weather := &mockWeatherProvider{
		data: &WeatherData{Temperature: 34, Humidity: 65, City: "Rio de Janeiro"},
	}
	tg := &mockTelegram{}
	svc := NewService(pets, users, weather, tg, mockClimateLogger{})

	if err := svc.CheckAndAlert(context.Background()); err != nil {
		t.Fatalf("first check: %v", err)
	}
	if len(tg.messages) != 1 {
		t.Fatalf("expected 1 message after first check, got %d", len(tg.messages))
	}

	if err := svc.CheckAndAlert(context.Background()); err != nil {
		t.Fatalf("second check: %v", err)
	}
	if len(tg.messages) != 1 {
		t.Fatalf("expected still 1 message after second check (anti-spam), got %d", len(tg.messages))
	}
}

func TestService_EmptyCity(t *testing.T) {
	pets := &mockPetRepo{
		pets: []*domain.Pet{
			{ID: "pet-5", UserID: "user-5", Name: "Dog", Location: ""},
		},
	}
	users := newMockUserRepo()
	users.users["user-5"] = &domain.User{ID: "user-5", TelegramID: 33333}
	weather := &mockWeatherProvider{
		data: &WeatherData{Temperature: 35, Humidity: 60, City: "Nowhere"},
	}
	tg := &mockTelegram{}
	svc := NewService(pets, users, weather, tg, mockClimateLogger{})

	if err := svc.CheckAndAlert(context.Background()); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if len(tg.messages) > 0 {
		t.Fatalf("expected no messages for empty city, got %d", len(tg.messages))
	}
}

func TestService_ProviderError(t *testing.T) {
	pets := &mockPetRepo{
		pets: []*domain.Pet{
			{ID: "pet-6", UserID: "user-6", Name: "Rex", Location: "São Paulo"},
		},
	}
	users := newMockUserRepo()
	users.users["user-6"] = &domain.User{ID: "user-6", TelegramID: 44444}
	weather := &mockWeatherProvider{err: context.DeadlineExceeded}
	tg := &mockTelegram{}
	svc := NewService(pets, users, weather, tg, mockClimateLogger{})

	if err := svc.CheckAndAlert(context.Background()); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if len(tg.messages) > 0 {
		t.Fatalf("expected no messages after provider error, got %d", len(tg.messages))
	}
}

func TestRules_EvaluateRisk(t *testing.T) {
	tests := []struct {
		temp float64
		want RiskLevel
	}{
		{20, RiskNone},
		{29, RiskNone},
		{30, RiskHigh},
		{32, RiskHigh},
		{34, RiskHigh},
		{35, RiskCritical},
		{40, RiskCritical},
	}
	for _, tc := range tests {
		got := EvaluateRisk(tc.temp)
		if got != tc.want {
			t.Errorf("EvaluateRisk(%.0f) = %d, want %d", tc.temp, got, tc.want)
		}
	}
}
