package reminder

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/marcusantonius88/braqui/apps/api/internal/domain"
)

type PetRepository interface {
	FindByUserID(ctx context.Context, userID string) ([]*domain.Pet, error)
	FindByID(ctx context.Context, id string) (*domain.Pet, error)
}

type ReminderRepository interface {
	Create(ctx context.Context, reminder *domain.Reminder) error
}

type UserRepository interface {
	FindByID(ctx context.Context, id string) (*domain.User, error)
}

type Logger interface {
	Info(msg string, fields map[string]any)
	Error(msg string, fields map[string]any)
}

type TelegramGateway interface {
	SendMessage(ctx context.Context, chatID int64, text string) error
}

type Service struct {
	pets    PetRepository
	reminds ReminderRepository
	users   UserRepository
	tg      TelegramGateway
	log     Logger
}

func NewService(pets PetRepository, reminds ReminderRepository, users UserRepository, tg TelegramGateway, log Logger) *Service {
	return &Service{pets: pets, reminds: reminds, users: users, tg: tg, log: log}
}

func (s *Service) CreateFromText(ctx context.Context, userID, text string) (*domain.Reminder, error) {
	cleaned := strings.TrimPrefix(text, "/remind ")
	cleaned = strings.TrimPrefix(cleaned, "/lembrar ")
	cleaned = strings.TrimPrefix(cleaned, "lembre ")
	cleaned = strings.TrimPrefix(cleaned, "me lembre de ")

	title, dateStr := splitTitleAndDate(cleaned)
	if title == "" {
		return nil, fmt.Errorf("no title found")
	}

	dueDate, err := parseDate(dateStr, time.Now())
	if err != nil {
		return nil, fmt.Errorf("invalid date: %w", err)
	}

	pets, err := s.pets.FindByUserID(ctx, userID)
	if err != nil {
		return nil, fmt.Errorf("find pets: %w", err)
	}
	if len(pets) == 0 {
		return nil, fmt.Errorf("no pets")
	}

	reminder := &domain.Reminder{
		PetID:       pets[0].ID,
		Title:       title,
		Description: text,
		DueDate:     dueDate,
		Status:      domain.ReminderStatusPending,
	}

	if err := s.reminds.Create(ctx, reminder); err != nil {
		return nil, fmt.Errorf("create reminder: %w", err)
	}

	s.log.Info("reminder created", map[string]any{"user_id": userID, "title": title, "due": dueDate.Format("2006-01-02")})
	return reminder, nil
}

func splitTitleAndDate(text string) (title, dateStr string) {
	text = strings.TrimSpace(text)

	keywords := []string{"daqui a", "em", "dia", "amanhã", "amanha", "hoje"}
	lower := strings.ToLower(text)

	bestIdx := -1
	for _, kw := range keywords {
		idx := strings.Index(lower, kw)
		if idx >= 0 && (bestIdx == -1 || idx < bestIdx) {
			bestIdx = idx
		}
	}

	if bestIdx < 0 {
		return text, ""
	}

	title = strings.TrimSpace(text[:bestIdx])
	dateStr = strings.TrimSpace(text[bestIdx:])
	return title, dateStr
}
