package pet

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"strconv"
	"strings"

	"github.com/marcusantonius88/braqui/apps/api/internal/domain"
)

type PetRepository interface {
	Create(ctx context.Context, pet *domain.Pet) error
	FindByUserID(ctx context.Context, userID string) ([]*domain.Pet, error)
}

type ConversationStateRepository interface {
	Create(ctx context.Context, state *domain.ConversationState) error
	FindByUserID(ctx context.Context, userID string) (*domain.ConversationState, error)
	Update(ctx context.Context, state *domain.ConversationState) error
}

type Onboarder struct {
	petRepo   PetRepository
	stateRepo ConversationStateRepository
}

func NewOnboarder(petRepo PetRepository, stateRepo ConversationStateRepository) *Onboarder {
	return &Onboarder{petRepo: petRepo, stateRepo: stateRepo}
}

type onboardingData struct {
	Name   string  `json:"name,omitempty"`
	Breed  string  `json:"breed,omitempty"`
	Age    int     `json:"age,omitempty"`
	Weight float64 `json:"weight,omitempty"`
	City   string  `json:"city,omitempty"`
}

func (o *Onboarder) Process(ctx context.Context, userID, text string) (reply string, err error) {
	pets, err := o.petRepo.FindByUserID(ctx, userID)
	if err != nil {
		return "", fmt.Errorf("find pets: %w", err)
	}
	if len(pets) > 0 {
		return "", nil
	}

	state, err := o.stateRepo.FindByUserID(ctx, userID)
	if err != nil && !errors.Is(err, domain.ErrNotFound) {
		return "", fmt.Errorf("find state: %w", err)
	}

	if errors.Is(err, domain.ErrNotFound) {
		state = &domain.ConversationState{
			UserID: userID,
			State:  "onboarding_name",
			Data:   []byte("{}"),
		}
		if err := o.stateRepo.Create(ctx, state); err != nil {
			return "", fmt.Errorf("create state: %w", err)
		}
		return "Qual o nome do seu cão?", nil
	}

	return o.advance(ctx, state, text)
}

func (o *Onboarder) advance(ctx context.Context, state *domain.ConversationState, text string) (string, error) {
	var data onboardingData
	if len(state.Data) > 0 {
		json.Unmarshal(state.Data, &data)
	}

	switch state.State {
	case "onboarding_name":
		if strings.TrimSpace(text) == "" {
			return "Por favor, digite o nome do seu cão.", nil
		}
		data.Name = strings.TrimSpace(text)
		state.State = "onboarding_breed"
		return saveAndReply(ctx, o.stateRepo, state, data, "Qual a raça do "+data.Name+"?")

	case "onboarding_breed":
		if strings.TrimSpace(text) == "" {
			return "Por favor, digite a raça do seu cão.", nil
		}
		data.Breed = strings.TrimSpace(text)
		state.State = "onboarding_age"
		return saveAndReply(ctx, o.stateRepo, state, data, "Qual a idade do "+data.Name+"? (em anos)")

	case "onboarding_age":
		age, err := strconv.Atoi(strings.TrimSpace(text))
		if err != nil || age <= 0 || age > 40 {
			return "Por favor, digite uma idade válida (número de 1 a 40).", nil
		}
		data.Age = age
		state.State = "onboarding_weight"
		return saveAndReply(ctx, o.stateRepo, state, data, "Qual o peso do "+data.Name+"? (em kg)")

	case "onboarding_weight":
		text = strings.TrimSpace(text)
		text = strings.ReplaceAll(text, ",", ".")
		weight, err := strconv.ParseFloat(text, 64)
		if err != nil || weight <= 0 || weight > 100 {
			return "Por favor, digite um peso válido (ex: 12.5).", nil
		}
		data.Weight = weight
		state.State = "onboarding_city"
		return saveAndReply(ctx, o.stateRepo, state, data, "Em qual cidade você mora?")

	case "onboarding_city":
		if strings.TrimSpace(text) == "" {
			return "Por favor, digite o nome da sua cidade.", nil
		}
		data.City = strings.TrimSpace(text)

		pet := &domain.Pet{
			UserID:   state.UserID,
			Name:     data.Name,
			Breed:    data.Breed,
			Age:      data.Age,
			Weight:   data.Weight,
			Location: data.City,
		}
		if err := o.petRepo.Create(ctx, pet); err != nil {
			return "", fmt.Errorf("create pet: %w", err)
		}

		if err := o.stateRepo.Update(ctx, state); err != nil {
			return "", fmt.Errorf("update state: %w", err)
		}

		return "Perfeito 🐶\n\nCadastro do " + data.Name + " concluído com sucesso!", nil

	default:
		return "", nil
	}
}

func saveAndReply(ctx context.Context, repo ConversationStateRepository, state *domain.ConversationState, data onboardingData, reply string) (string, error) {
	encoded, err := json.Marshal(data)
	if err != nil {
		return "", fmt.Errorf("marshal data: %w", err)
	}
	state.Data = encoded
	if err := repo.Update(ctx, state); err != nil {
		return "", fmt.Errorf("update state: %w", err)
	}
	return reply, nil
}
