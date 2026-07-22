package insights

import (
	"context"
	"fmt"

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
	FindByUserID(ctx context.Context, userID string) ([]*domain.Pet, error)
}

type Service struct {
	pets   PetRepository
	events EventRepository
	log    Logger
}

func NewService(pets PetRepository, events EventRepository, log Logger) *Service {
	return &Service{pets: pets, events: events, log: log}
}

func (s *Service) Generate(ctx context.Context, userID string) (string, error) {
	allPets, err := s.pets.FindByUserID(ctx, userID)
	if err != nil {
		return "", fmt.Errorf("find pets: %w", err)
	}
	if len(allPets) == 0 {
		return "Você precisa cadastrar um pet antes de usar insights.", nil
	}

	pet := allPets[0]
	events, err := s.events.FindByPetID(ctx, pet.ID)
	if err != nil {
		return "", fmt.Errorf("find events: %w", err)
	}
	if len(events) == 0 {
		return "Ainda não há dados suficientes para gerar insights.", nil
	}

	var insights []*Insight
	for _, rule := range rules {
		if ins := evaluateRule(pet.Name, events, rule); ins != nil {
			insights = append(insights, ins)
		}
	}

	if len(insights) == 0 {
		return "Não encontrei padrões relevantes no momento.", nil
	}

	msg := fmt.Sprintf("📊 Insights do %s", pet.Name)
	for _, ins := range insights {
		msg += "\n\n• " + ins.Message
	}

	s.log.Info("insights generated", map[string]any{
		"pet_id":  pet.ID,
		"count":   len(insights),
		"user_id": userID,
	})

	return msg, nil
}
