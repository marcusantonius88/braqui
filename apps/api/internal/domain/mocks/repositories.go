package mocks

import (
	"context"
	"fmt"
	"sync"
	"time"

	"github.com/marcusantonius88/braqui/apps/api/internal/domain"
)

type UserRepository struct {
	mu    sync.Mutex
	users map[string]*domain.User
}

func NewUserRepository() *UserRepository {
	return &UserRepository{users: make(map[string]*domain.User)}
}

func (r *UserRepository) Create(ctx context.Context, user *domain.User) error {
	r.mu.Lock()
	defer r.mu.Unlock()
	id := fmt.Sprintf("user-%d", len(r.users)+1)
	user.ID = id
	user.CreatedAt = time.Now()
	user.UpdatedAt = time.Now()
	r.users[id] = user
	return nil
}

func (r *UserRepository) FindByID(ctx context.Context, id string) (*domain.User, error) {
	r.mu.Lock()
	defer r.mu.Unlock()
	user, ok := r.users[id]
	if !ok {
		return nil, domain.ErrNotFound
	}
	return user, nil
}

func (r *UserRepository) FindByTelegramID(ctx context.Context, telegramID int64) (*domain.User, error) {
	r.mu.Lock()
	defer r.mu.Unlock()
	for _, u := range r.users {
		if u.TelegramID == telegramID {
			return u, nil
		}
	}
	return nil, domain.ErrNotFound
}

type PetRepository struct {
	mu   sync.Mutex
	pets map[string]*domain.Pet
}

func NewPetRepository() *PetRepository {
	return &PetRepository{pets: make(map[string]*domain.Pet)}
}

func (r *PetRepository) Create(ctx context.Context, pet *domain.Pet) error {
	r.mu.Lock()
	defer r.mu.Unlock()
	id := fmt.Sprintf("pet-%d", len(r.pets)+1)
	pet.ID = id
	pet.CreatedAt = time.Now()
	pet.UpdatedAt = time.Now()
	r.pets[id] = pet
	return nil
}

func (r *PetRepository) FindByID(ctx context.Context, id string) (*domain.Pet, error) {
	r.mu.Lock()
	defer r.mu.Unlock()
	pet, ok := r.pets[id]
	if !ok {
		return nil, domain.ErrNotFound
	}
	return pet, nil
}

func (r *PetRepository) FindByUserID(ctx context.Context, userID string) ([]*domain.Pet, error) {
	r.mu.Lock()
	defer r.mu.Unlock()
	var result []*domain.Pet
	for _, p := range r.pets {
		if p.UserID == userID {
			result = append(result, p)
		}
	}
	if result == nil {
		return []*domain.Pet{}, nil
	}
	return result, nil
}

type EventRepository struct {
	mu     sync.Mutex
	events map[string]*domain.Event
}

func NewEventRepository() *EventRepository {
	return &EventRepository{events: make(map[string]*domain.Event)}
}

func (r *EventRepository) Create(ctx context.Context, event *domain.Event) error {
	r.mu.Lock()
	defer r.mu.Unlock()
	id := fmt.Sprintf("event-%d", len(r.events)+1)
	event.ID = id
	event.CreatedAt = time.Now()
	r.events[id] = event
	return nil
}

func (r *EventRepository) FindByID(ctx context.Context, id string) (*domain.Event, error) {
	r.mu.Lock()
	defer r.mu.Unlock()
	event, ok := r.events[id]
	if !ok {
		return nil, domain.ErrNotFound
	}
	return event, nil
}

func (r *EventRepository) FindByPetID(ctx context.Context, petID string) ([]*domain.Event, error) {
	r.mu.Lock()
	defer r.mu.Unlock()
	var result []*domain.Event
	for _, e := range r.events {
		if e.PetID == petID {
			result = append(result, e)
		}
	}
	if result == nil {
		return []*domain.Event{}, nil
	}
	return result, nil
}

type ReminderRepository struct {
	mu        sync.Mutex
	reminders map[string]*domain.Reminder
}

func NewReminderRepository() *ReminderRepository {
	return &ReminderRepository{reminders: make(map[string]*domain.Reminder)}
}

func (r *ReminderRepository) Create(ctx context.Context, reminder *domain.Reminder) error {
	r.mu.Lock()
	defer r.mu.Unlock()
	id := fmt.Sprintf("reminder-%d", len(r.reminders)+1)
	reminder.ID = id
	reminder.CreatedAt = time.Now()
	reminder.UpdatedAt = time.Now()
	r.reminders[id] = reminder
	return nil
}

func (r *ReminderRepository) FindByID(ctx context.Context, id string) (*domain.Reminder, error) {
	r.mu.Lock()
	defer r.mu.Unlock()
	reminder, ok := r.reminders[id]
	if !ok {
		return nil, domain.ErrNotFound
	}
	return reminder, nil
}

func (r *ReminderRepository) FindByPetID(ctx context.Context, petID string) ([]*domain.Reminder, error) {
	r.mu.Lock()
	defer r.mu.Unlock()
	var result []*domain.Reminder
	for _, rem := range r.reminders {
		if rem.PetID == petID {
			result = append(result, rem)
		}
	}
	if result == nil {
		return []*domain.Reminder{}, nil
	}
	return result, nil
}

func (r *ReminderRepository) FindPendingDueBefore(_ context.Context, due time.Time) ([]*domain.Reminder, error) {
	r.mu.Lock()
	defer r.mu.Unlock()
	var result []*domain.Reminder
	for _, rem := range r.reminders {
		if rem.Status == domain.ReminderStatusPending && !rem.DueDate.After(due) {
			result = append(result, rem)
		}
	}
	if result == nil {
		return []*domain.Reminder{}, nil
	}
	return result, nil
}

func (r *ReminderRepository) UpdateStatus(_ context.Context, id, status string) error {
	r.mu.Lock()
	defer r.mu.Unlock()
	rem, ok := r.reminders[id]
	if !ok {
		return domain.ErrNotFound
	}
	rem.Status = status
	rem.UpdatedAt = time.Now()
	return nil
}

type ConversationStateRepository struct {
	mu     sync.Mutex
	states map[string]*domain.ConversationState
}

func NewConversationStateRepository() *ConversationStateRepository {
	return &ConversationStateRepository{states: make(map[string]*domain.ConversationState)}
}

func (r *ConversationStateRepository) Create(ctx context.Context, state *domain.ConversationState) error {
	r.mu.Lock()
	defer r.mu.Unlock()
	id := fmt.Sprintf("state-%d", len(r.states)+1)
	state.ID = id
	state.CreatedAt = time.Now()
	state.UpdatedAt = time.Now()
	r.states[id] = state
	return nil
}

func (r *ConversationStateRepository) FindByUserID(ctx context.Context, userID string) (*domain.ConversationState, error) {
	r.mu.Lock()
	defer r.mu.Unlock()
	for _, s := range r.states {
		if s.UserID == userID {
			return s, nil
		}
	}
	return nil, domain.ErrNotFound
}

func (r *ConversationStateRepository) Update(ctx context.Context, state *domain.ConversationState) error {
	r.mu.Lock()
	defer r.mu.Unlock()
	if _, ok := r.states[state.ID]; !ok {
		return domain.ErrNotFound
	}
	state.UpdatedAt = time.Now()
	r.states[state.ID] = state
	return nil
}
