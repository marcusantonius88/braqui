package state

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"

	"github.com/marcusantonius88/braqui/apps/api/internal/domain"
)

type StateRepository interface {
	Create(ctx context.Context, state *domain.ConversationState) error
	FindByUserID(ctx context.Context, userID string) (*domain.ConversationState, error)
	Update(ctx context.Context, state *domain.ConversationState) error
}

type Logger interface {
	Info(msg string, fields map[string]any)
	Error(msg string, fields map[string]any)
}

type Manager struct {
	repo StateRepository
	log  Logger
}

func NewManager(repo StateRepository, log Logger) *Manager {
	return &Manager{repo: repo, log: log}
}

func (m *Manager) Start(ctx context.Context, userID, flow, step string, payload any) (*domain.ConversationState, error) {
	raw, err := json.Marshal(payload)
	if err != nil {
		return nil, fmt.Errorf("marshal payload: %w", err)
	}
	state := &domain.ConversationState{
		UserID:  userID,
		Flow:    flow,
		Step:    step,
		Payload: raw,
	}
	if err := m.repo.Create(ctx, state); err != nil {
		m.log.Error("failed to create state", map[string]any{"user_id": userID, "flow": flow, "error": err.Error()})
		return nil, fmt.Errorf("create state: %w", err)
	}
	m.log.Info("flow started", map[string]any{"user_id": userID, "flow": flow, "step": step})
	return state, nil
}

func (m *Manager) Get(ctx context.Context, userID string) (*domain.ConversationState, error) {
	state, err := m.repo.FindByUserID(ctx, userID)
	if err != nil {
		if errors.Is(err, domain.ErrNotFound) {
			return nil, nil
		}
		return nil, fmt.Errorf("find state: %w", err)
	}
	return state, nil
}

func (m *Manager) Advance(ctx context.Context, userID, nextStep string, payload any) (*domain.ConversationState, error) {
	state, err := m.repo.FindByUserID(ctx, userID)
	if err != nil {
		m.log.Error("failed to find state for advance", map[string]any{"user_id": userID, "error": err.Error()})
		return nil, fmt.Errorf("find state for advance: %w", err)
	}
	raw, err := json.Marshal(payload)
	if err != nil {
		return nil, fmt.Errorf("marshal payload: %w", err)
	}
	prevStep := state.Step
	state.Step = nextStep
	state.Payload = raw
	if err := m.repo.Update(ctx, state); err != nil {
		m.log.Error("failed to advance state", map[string]any{"user_id": userID, "step": nextStep, "error": err.Error()})
		return nil, fmt.Errorf("update state: %w", err)
	}
	m.log.Info("step changed", map[string]any{"user_id": userID, "from": prevStep, "to": nextStep})
	return state, nil
}

func (m *Manager) Complete(ctx context.Context, userID string) error {
	state, err := m.repo.FindByUserID(ctx, userID)
	if err != nil {
		m.log.Error("failed to find state for complete", map[string]any{"user_id": userID, "error": err.Error()})
		return fmt.Errorf("find state for complete: %w", err)
	}
	flow := state.Flow
	state.Flow = ""
	state.Step = ""
	if err := m.repo.Update(ctx, state); err != nil {
		m.log.Error("failed to complete state", map[string]any{"user_id": userID, "error": err.Error()})
		return fmt.Errorf("update state on complete: %w", err)
	}
	m.log.Info("flow completed", map[string]any{"user_id": userID, "flow": flow})
	return nil
}
