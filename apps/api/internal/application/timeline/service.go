package timeline

import (
	"context"
	"fmt"
	"sort"
	"time"

	"github.com/marcusantonius88/braqui/apps/api/internal/domain"
)

type PetRepository interface {
	FindByUserID(ctx context.Context, userID string) ([]*domain.Pet, error)
}

type EventRepository interface {
	FindByPetID(ctx context.Context, petID string) ([]*domain.Event, error)
}

type Logger interface {
	Info(msg string, fields map[string]any)
	Error(msg string, fields map[string]any)
}

type Service struct {
	pets   PetRepository
	events EventRepository
	log    Logger
}

func NewService(pets PetRepository, events EventRepository, log Logger) *Service {
	return &Service{pets: pets, events: events, log: log}
}

func (s *Service) GetTimeline(ctx context.Context, userID string) (string, error) {
	pets, err := s.pets.FindByUserID(ctx, userID)
	if err != nil {
		s.log.Error("failed to find pets", map[string]any{"user_id": userID, "error": err.Error()})
		return "", fmt.Errorf("find pets: %w", err)
	}
	if len(pets) == 0 {
		return "Você ainda não cadastrou nenhum pet. Envie uma mensagem para começar!", nil
	}

	pet := pets[0]
	events, err := s.events.FindByPetID(ctx, pet.ID)
	if err != nil {
		s.log.Error("failed to query events", map[string]any{"user_id": userID, "error": err.Error()})
		return "Não consegui consultar o histórico agora. Tente novamente em alguns minutos.", nil
	}

	if len(events) == 0 {
		return fmt.Sprintf("Ainda não encontrei eventos registrados para o %s 🐶", pet.Name), nil
	}

	sort.Slice(events, func(i, j int) bool {
		return events[i].Timestamp.After(events[j].Timestamp)
	})

	if len(events) > 10 {
		events = events[:10]
	}

	lines := formatEvents(events)

	reply := fmt.Sprintf("📋 Histórico recente do %s\n\n%s", pet.Name, lines)

	s.log.Info("timeline queried", map[string]any{"user_id": userID, "pet_id": pet.ID, "count": len(events)})
	return reply, nil
}

func formatEvents(events []*domain.Event) string {
	var result string
	for i, e := range events {
		when := relativeDate(e.Timestamp)
		label := eventLabel(e.Type)
		if i == len(events)-1 {
			result += fmt.Sprintf("• %s - %s", when, label)
		} else {
			result += fmt.Sprintf("• %s - %s\n", when, label)
		}
	}
	return result
}

func relativeDate(t time.Time) string {
	now := time.Now()
	t = t.Truncate(24 * time.Hour)
	today := now.Truncate(24 * time.Hour)

	diff := today.Sub(t)
	days := int(diff.Hours() / 24)

	switch {
	case days == 0:
		return "Hoje"
	case days == 1:
		return "Ontem"
	default:
		return fmt.Sprintf("%d dias atrás", days)
	}
}

func eventLabel(eventType string) string {
	switch eventType {
	case "vomit":
		return "Vômito"
	case "itching":
		return "Coceira"
	case "panting":
		return "Ofegante"
	case "fatigue":
		return "Cansaço"
	case "cough":
		return "Tosse"
	case "diarrhea":
		return "Diarreia"
	case "medication_given":
		return "Medicação"
	case "vet_visit":
		return "Consulta veterinária"
	case "diarrhoea":
		return "Diarreia"
	default:
		return eventType
	}
}
