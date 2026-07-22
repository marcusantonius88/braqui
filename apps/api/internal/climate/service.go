package climate

import (
	"context"
	"fmt"
	"sync"
	"time"

	"github.com/marcusantonius88/braqui/apps/api/internal/domain"
)

type PetRepository interface {
	FindAllWithLocation(ctx context.Context) ([]*domain.Pet, error)
}

type Logger interface {
	Info(msg string, fields map[string]any)
	Error(msg string, fields map[string]any)
}

type UserRepository interface {
	FindByID(ctx context.Context, id string) (*domain.User, error)
}

type TelegramGateway interface {
	SendMessage(ctx context.Context, chatID int64, text string) error
}

type Service struct {
	pets    PetRepository
	users   UserRepository
	weather Provider
	tg      TelegramGateway
	log     Logger

	mu           sync.Mutex
	alertedToday map[string]string
}

func NewService(pets PetRepository, users UserRepository, weather Provider, tg TelegramGateway, log Logger) *Service {
	return &Service{
		pets:         pets,
		users:        users,
		weather:      weather,
		tg:           tg,
		log:          log,
		alertedToday: make(map[string]string),
	}
}

func (s *Service) CheckAndAlert(ctx context.Context) error {
	allPets, err := s.pets.FindAllWithLocation(ctx)
	if err != nil {
		return fmt.Errorf("find pets with location: %w", err)
	}

	today := time.Now().Format("2006-01-02")

	for _, pet := range allPets {
		if pet.Location == "" {
			continue
		}

		s.mu.Lock()
		alertDate, exists := s.alertedToday[pet.ID]
		s.mu.Unlock()
		if exists && alertDate == today {
			continue
		}

		data, err := s.weather.GetCurrentWeather(ctx, pet.Location)
		if err != nil {
			s.log.Error("climate: fetch weather", map[string]any{"pet_id": pet.ID, "city": pet.Location, "error": err.Error()})
			continue
		}

		risk := EvaluateRisk(data.Temperature)
		if risk == RiskNone {
			continue
		}

		user, err := s.users.FindByID(ctx, pet.UserID)
		if err != nil {
			s.log.Error("climate: find user", map[string]any{"pet_id": pet.ID, "error": err.Error()})
			continue
		}

		msg := FormatAlert(pet.Location, data.Temperature, risk)
		if err := s.tg.SendMessage(ctx, user.TelegramID, msg); err != nil {
			s.log.Error("climate: send alert", map[string]any{"pet_id": pet.ID, "chat_id": user.TelegramID, "error": err.Error()})
			continue
		}

		s.mu.Lock()
		s.alertedToday[pet.ID] = today
		s.mu.Unlock()

		s.log.Info("climate alert sent", map[string]any{"pet_id": pet.ID, "city": pet.Location, "temp": data.Temperature, "risk": risk.String()})
	}

	return nil
}
