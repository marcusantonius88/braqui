package summary

import (
	"context"
	"fmt"
	"time"

	"github.com/marcusantonius88/braqui/apps/api/internal/domain"
)

type Logger interface {
	Info(msg string, fields map[string]any)
	Error(msg string, fields map[string]any)
}

type EventRepository interface {
	FindByPetID(ctx context.Context, petID string) ([]*domain.Event, error)
}

type PetRepository interface {
	FindAll(ctx context.Context) ([]*domain.Pet, error)
}

type UserRepository interface {
	FindByID(ctx context.Context, id string) (*domain.User, error)
}

type TelegramGateway interface {
	SendMessage(ctx context.Context, chatID int64, text string) error
}

type Service struct {
	pets   PetRepository
	events EventRepository
	users  UserRepository
	tg     TelegramGateway
	log    Logger
}

func NewService(pets PetRepository, events EventRepository, users UserRepository, tg TelegramGateway, log Logger) *Service {
	return &Service{pets: pets, events: events, users: users, tg: tg, log: log}
}

func (s *Service) GenerateAndSend(ctx context.Context) error {
	allPets, err := s.pets.FindAll(ctx)
	if err != nil {
		return fmt.Errorf("find all pets: %w", err)
	}

	weekAgo := time.Now().AddDate(0, 0, -7)

	for _, pet := range allPets {
		events, err := s.events.FindByPetID(ctx, pet.ID)
		if err != nil {
			s.log.Error("summary: find events", map[string]any{"pet_id": pet.ID, "error": err.Error()})
			continue
		}

		counts := map[string]int{}
		total := 0
		for _, e := range events {
			if !e.Timestamp.Before(weekAgo) {
				counts[e.Type]++
				total++
			}
		}

		user, err := s.users.FindByID(ctx, pet.UserID)
		if err != nil {
			s.log.Error("summary: find user", map[string]any{"pet_id": pet.ID, "error": err.Error()})
			continue
		}

		msg := formatSummary(pet.Name, counts, total)
		if err := s.tg.SendMessage(ctx, user.TelegramID, msg); err != nil {
			s.log.Error("summary: send failed", map[string]any{"pet_id": pet.ID, "chat_id": user.TelegramID, "error": err.Error()})
			continue
		}

		s.log.Info("summary sent", map[string]any{"pet_id": pet.ID, "pet": pet.Name, "total_events": total})
	}

	return nil
}
