package pet

import (
	"context"
	"encoding/json"
	"fmt"
	"strconv"
	"strings"

	"github.com/marcusantonius88/braqui/apps/api/internal/domain"
)

type PetRepository interface {
	Create(ctx context.Context, pet *domain.Pet) error
	FindByUserID(ctx context.Context, userID string) ([]*domain.Pet, error)
}

type StateManager interface {
	Start(ctx context.Context, userID, flow, step string, payload any) (*domain.ConversationState, error)
	Get(ctx context.Context, userID string) (*domain.ConversationState, error)
	Advance(ctx context.Context, userID, nextStep string, payload any) (*domain.ConversationState, error)
	Complete(ctx context.Context, userID string) error
}

type Onboarder struct {
	petRepo PetRepository
	states  StateManager
}

func NewOnboarder(petRepo PetRepository, states StateManager) *Onboarder {
	return &Onboarder{petRepo: petRepo, states: states}
}

const (
	flowRegisterPet = "register_pet"

	stepAskName   = "ask_name"
	stepAskBreed  = "ask_breed"
	stepAskAge    = "ask_age"
	stepAskWeight = "ask_weight"
	stepAskCity   = "ask_city"
)

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

	state, err := o.states.Get(ctx, userID)
	if err != nil {
		return "", fmt.Errorf("get state: %w", err)
	}

	if state == nil {
		state, err = o.states.Start(ctx, userID, flowRegisterPet, stepAskName, onboardingData{})
		if err != nil {
			return "", fmt.Errorf("start state: %w", err)
		}
		return "Qual o nome do seu cão?", nil
	}

	return o.advance(ctx, state, text)
}

func (o *Onboarder) advance(ctx context.Context, state *domain.ConversationState, text string) (string, error) {
	var data onboardingData
	if len(state.Payload) > 0 {
		json.Unmarshal(state.Payload, &data)
	}

	switch state.Step {
	case stepAskName:
		if strings.TrimSpace(text) == "" {
			return "Por favor, digite o nome do seu cão.", nil
		}
		data.Name = strings.TrimSpace(text)
		if _, err := o.states.Advance(ctx, state.UserID, stepAskBreed, data); err != nil {
			return "", fmt.Errorf("advance to breed: %w", err)
		}
		return "Qual a raça do " + data.Name + "?", nil

	case stepAskBreed:
		if strings.TrimSpace(text) == "" {
			return "Por favor, digite a raça do seu cão.", nil
		}
		data.Breed = strings.TrimSpace(text)
		if _, err := o.states.Advance(ctx, state.UserID, stepAskAge, data); err != nil {
			return "", fmt.Errorf("advance to age: %w", err)
		}
		return "Qual a idade do " + data.Name + "? (em anos)", nil

	case stepAskAge:
		age, err := strconv.Atoi(strings.TrimSpace(text))
		if err != nil || age <= 0 || age > 40 {
			return "Por favor, digite uma idade válida (número de 1 a 40).", nil
		}
		data.Age = age
		if _, err := o.states.Advance(ctx, state.UserID, stepAskWeight, data); err != nil {
			return "", fmt.Errorf("advance to weight: %w", err)
		}
		return "Qual o peso do " + data.Name + "? (em kg)", nil

	case stepAskWeight:
		text = strings.TrimSpace(text)
		text = strings.ReplaceAll(text, ",", ".")
		weight, err := strconv.ParseFloat(text, 64)
		if err != nil || weight <= 0 || weight > 100 {
			return "Por favor, digite um peso válido (ex: 12.5).", nil
		}
		data.Weight = weight
		if _, err := o.states.Advance(ctx, state.UserID, stepAskCity, data); err != nil {
			return "", fmt.Errorf("advance to city: %w", err)
		}
		return "Em qual cidade você mora?", nil

	case stepAskCity:
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
		if err := o.states.Complete(ctx, state.UserID); err != nil {
			return "", fmt.Errorf("complete state: %w", err)
		}
		return "Perfeito 🐶\n\nCadastro do " + data.Name + " concluído com sucesso!", nil

	default:
		return "", nil
	}
}
