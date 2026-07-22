package location

import (
	"context"
	"fmt"
	"strings"

	"github.com/marcusantonius88/braqui/apps/api/internal/domain"
)

type PetRepository interface {
	FindByUserID(ctx context.Context, userID string) ([]*domain.Pet, error)
	UpdateLocation(ctx context.Context, petID, city string) error
}

type Logger interface {
	Info(msg string, fields map[string]any)
	Error(msg string, fields map[string]any)
}

type StateManager interface {
	Start(ctx context.Context, userID, flow, step string, payload any) (*domain.ConversationState, error)
	Get(ctx context.Context, userID string) (*domain.ConversationState, error)
	Advance(ctx context.Context, userID, nextStep string, payload any) (*domain.ConversationState, error)
	Complete(ctx context.Context, userID string) error
}

type Updater struct {
	pets PetRepository
	state StateManager
	log   Logger
}

func NewUpdater(pets PetRepository, state StateManager, log Logger) *Updater {
	return &Updater{pets: pets, state: state, log: log}
}

func (u *Updater) Process(ctx context.Context, userID, text string) (string, error) {
	existingState, err := u.state.Get(ctx, userID)
	if err != nil {
		return "", fmt.Errorf("get state: %w", err)
	}

	if existingState == nil || existingState.Flow == "" || existingState.Flow != "change_city" {
		pets, err := u.pets.FindByUserID(ctx, userID)
		if err != nil {
			return "", fmt.Errorf("find pets: %w", err)
		}
		if len(pets) == 0 {
			return "Você precisa cadastrar um pet antes de configurar a cidade.", nil
		}

		payload := map[string]any{"pet_id": pets[0].ID, "current_city": pets[0].Location}
		if _, err := u.state.Start(ctx, userID, "change_city", "ask_city", payload); err != nil {
			return "", fmt.Errorf("start state: %w", err)
		}

		current := pets[0].Location
		if current == "" {
			return "Em qual cidade você mora?", nil
		}
		return fmt.Sprintf("Cidade atual: %s. Para qual cidade deseja mudar?", current), nil
	}

	return u.handleCity(ctx, userID, text, existingState)
}

func (u *Updater) handleCity(ctx context.Context, userID, text string, _ *domain.ConversationState) (string, error) {
	city := strings.TrimSpace(text)
	if city == "" {
		return "Por favor, digite o nome de uma cidade.", nil
	}

	pets, err := u.pets.FindByUserID(ctx, userID)
	if err != nil {
		return "", fmt.Errorf("find pets: %w", err)
	}
	if len(pets) == 0 {
		return "Pet não encontrado.", nil
	}

	if err := u.pets.UpdateLocation(ctx, pets[0].ID, city); err != nil {
		u.log.Error("failed to update location", map[string]any{"user_id": userID, "error": err.Error()})
		return "Não consegui atualizar a cidade agora. Tente novamente.", nil
	}

	if err := u.state.Complete(ctx, userID); err != nil {
		return "", fmt.Errorf("complete state: %w", err)
	}

	u.log.Info("location updated", map[string]any{"user_id": userID, "city": city})
	return fmt.Sprintf("Cidade atualizada para %s! 🐶", city), nil
}
