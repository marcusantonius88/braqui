package reminder

import (
	"context"
	"encoding/json"
	"fmt"
	"strings"
	"time"

	"github.com/marcusantonius88/braqui/apps/api/internal/domain"
)

type StateManager interface {
	Start(ctx context.Context, userID, flow, step string, payload any) (*domain.ConversationState, error)
	Get(ctx context.Context, userID string) (*domain.ConversationState, error)
	Advance(ctx context.Context, userID, nextStep string, payload any) (*domain.ConversationState, error)
	Complete(ctx context.Context, userID string) error
}

type Creator struct {
	pets        PetRepository
	reminds     ReminderRepository
	state       StateManager
	log         Logger
}

func NewCreator(pets PetRepository, reminds ReminderRepository, state StateManager, log Logger) *Creator {
	return &Creator{pets: pets, reminds: reminds, state: state, log: log}
}

func (c *Creator) Process(ctx context.Context, userID, text string) (string, error) {
	existingState, err := c.state.Get(ctx, userID)
	if err != nil {
		return "", fmt.Errorf("get state: %w", err)
	}

	if existingState == nil || existingState.Flow == "" || existingState.Flow != "create_reminder" {
		pets, err := c.pets.FindByUserID(ctx, userID)
		if err != nil {
			return "", fmt.Errorf("find pets: %w", err)
		}
		if len(pets) == 0 {
			return "Você precisa cadastrar um pet antes de criar lembretes.", nil
		}

		_, err = c.state.Start(ctx, userID, "create_reminder", "ask_title", map[string]any{})
		if err != nil {
			return "", fmt.Errorf("start state: %w", err)
		}
		return "O que você quer lembrar?", nil
	}

	switch existingState.Step {
	case "ask_title":
		return c.handleTitle(ctx, userID, text, existingState)
	case "ask_date":
		return c.handleDate(ctx, userID, text, existingState)
	default:
		return "", nil
	}
}

func (c *Creator) handleTitle(ctx context.Context, userID, text string, state *domain.ConversationState) (string, error) {
	payload := map[string]any{"title": text}
	if _, err := c.state.Advance(ctx, userID, "ask_date", payload); err != nil {
		return "", fmt.Errorf("advance: %w", err)
	}
	return "Para quando?", nil
}

func (c *Creator) handleDate(ctx context.Context, userID, text string, state *domain.ConversationState) (string, error) {
	dueDate, err := parseDate(text, time.Now())
	if err != nil {
		return "Não consegui entender a data informada. Digite algo como \"amanhã\", \"dia 15\" ou \"daqui a 7 dias\".", nil
	}

	var payload map[string]any
	if err := json.Unmarshal(state.Payload, &payload); err != nil {
		return "", fmt.Errorf("unmarshal payload: %w", err)
	}

	title, _ := payload["title"].(string)

	pets, err := c.pets.FindByUserID(ctx, userID)
	if err != nil {
		return "", fmt.Errorf("find pets: %w", err)
	}
	if len(pets) == 0 {
		return "Você precisa cadastrar um pet antes de criar lembretes.", nil
	}

	reminder := &domain.Reminder{
		PetID:       pets[0].ID,
		Title:       title,
		Description: fmt.Sprintf("Lembrete: %s", title),
		DueDate:     dueDate,
		Status:      domain.ReminderStatusPending,
	}

	if err := c.reminds.Create(ctx, reminder); err != nil {
		return "", fmt.Errorf("create reminder: %w", err)
	}

	if err := c.state.Complete(ctx, userID); err != nil {
		return "", fmt.Errorf("complete state: %w", err)
	}

	dateStr := dueDate.Format("02/01/2006")
	if dueDate.Truncate(24*time.Hour).Equal(time.Now().Truncate(24 * time.Hour)) {
		dateStr = "hoje"
	} else if dueDate.Truncate(24*time.Hour).Equal(time.Now().Truncate(24*time.Hour).Add(24 * time.Hour)) {
		dateStr = "amanhã"
	}

	c.log.Info("reminder created via flow", map[string]any{"user_id": userID, "title": title, "due": dueDate.Format("2006-01-02")})
	return fmt.Sprintf("Lembrete criado! Vou te lembrar de \"%s\" %s.", strings.ToLower(title), dateStr), nil
}
