package event

import (
	"context"
	"fmt"
	"time"

	"github.com/marcusantonius88/braqui/apps/api/internal/domain"
	domainai "github.com/marcusantonius88/braqui/apps/api/internal/domain/ai"
	"github.com/marcusantonius88/braqui/apps/api/internal/parser"
)

type PetRepository interface {
	FindByUserID(ctx context.Context, userID string) ([]*domain.Pet, error)
}

type EventRepository interface {
	Create(ctx context.Context, event *domain.Event) error
	FindByPetID(ctx context.Context, petID string) ([]*domain.Event, error)
}

type AIProvider interface {
	Interpret(ctx context.Context, message string) (*domainai.InterpretationResult, error)
}

type Logger interface {
	Info(msg string, fields map[string]any)
	Error(msg string, fields map[string]any)
}

type Registerer struct {
	pets        PetRepository
	events      EventRepository
	ai          AIProvider
	log         Logger
}

func NewRegisterer(pets PetRepository, events EventRepository, ai AIProvider, log Logger) *Registerer {
	return &Registerer{pets: pets, events: events, ai: ai, log: log}
}

type RegisterResult string

const (
	ResultRegistered      RegisterResult = "registered"
	ResultNotInterpreted  RegisterResult = "not_interpreted"
	ResultNoPet           RegisterResult = "no_pet"
)

func (r *Registerer) Register(ctx context.Context, userID, text string) (RegisterResult, error) {
	pets, err := r.pets.FindByUserID(ctx, userID)
	if err != nil {
		return "", fmt.Errorf("find pets: %w", err)
	}
	if len(pets) == 0 {
		return ResultNoPet, nil
	}

	parseResult := parser.Parse(text)
	eventType := parseResult.Type
	source := "parser"

	if eventType == "NOT_MATCHED" {
		r.log.Info("parser not matched, calling AI", map[string]any{"user_id": userID})
		aiResult, err := r.ai.Interpret(ctx, text)
		if err != nil {
			r.log.Error("ai interpret error", map[string]any{"user_id": userID, "error": err.Error()})
			return ResultNotInterpreted, nil
		}
		if aiResult.Type == "NOT_INTERPRETED" {
			r.log.Info("ai not interpreted", map[string]any{"user_id": userID})
			return ResultNotInterpreted, nil
		}
		eventType = aiResult.Type
		source = "ai"
	}

	pet := pets[0]
	event := &domain.Event{
		PetID:       pet.ID,
		Type:        eventType,
		Description: text,
		Source:      source,
		Timestamp:   time.Now(),
	}

	if err := r.events.Create(ctx, event); err != nil {
		r.log.Error("failed to create event", map[string]any{"user_id": userID, "error": err.Error()})
		return "", fmt.Errorf("create event: %w", err)
	}

	r.log.Info("event registered", map[string]any{"user_id": userID, "type": eventType, "source": source})
	return ResultRegistered, nil
}
