package domain

import (
	"errors"
	"testing"
	"time"
)

func TestUserCreation(t *testing.T) {
	u := &User{TelegramID: 123, FirstName: "João"}
	if u.TelegramID != 123 {
		t.Fatalf("expected 123, got %d", u.TelegramID)
	}
	if u.FirstName != "João" {
		t.Fatalf("expected João, got %s", u.FirstName)
	}
	if u.ID != "" {
		t.Fatal("expected zero value for ID")
	}
	if !u.CreatedAt.IsZero() {
		t.Fatal("expected zero CreatedAt")
	}
}

func TestPetCreation(t *testing.T) {
	p := &Pet{
		UserID:   "user-1",
		Name:     "Rex",
		Breed:    "Bulldog",
		Age:      3,
		Weight:   22.5,
		Location: "SP",
	}
	if p.Name != "Rex" {
		t.Fatalf("expected Rex, got %s", p.Name)
	}
	if p.Age != 3 {
		t.Fatalf("expected 3, got %d", p.Age)
	}
	if p.Breed != "Bulldog" {
		t.Fatalf("expected Bulldog, got %s", p.Breed)
	}
	if p.Weight != 22.5 {
		t.Fatalf("expected 22.5, got %f", p.Weight)
	}
}

func TestEventCreation(t *testing.T) {
	e := &Event{
		PetID:       "pet-1",
		Type:        "feeding",
		Description: "comeu ração",
		Timestamp:   time.Now(),
	}
	if e.Type != "feeding" {
		t.Fatalf("expected feeding, got %s", e.Type)
	}
	if e.Description != "comeu ração" {
		t.Fatalf("expected comeu ração, got %s", e.Description)
	}
}

func TestReminderCreation(t *testing.T) {
	r := &Reminder{
		PetID:          "pet-1",
		Type:           "medication",
		Description:    "antipulga",
		DueDate:        time.Now(),
		RepeatInterval: "30d",
	}
	if r.RepeatInterval != "30d" {
		t.Fatalf("expected 30d, got %s", r.RepeatInterval)
	}
}

func TestConversationStateCreation(t *testing.T) {
	s := &ConversationState{
		UserID: "user-1",
		State:  "awaiting_pet_name",
		Data:   []byte(`{}`),
	}
	if s.State != "awaiting_pet_name" {
		t.Fatalf("expected awaiting_pet_name, got %s", s.State)
	}
	if string(s.Data) != "{}" {
		t.Fatalf("expected {}, got %s", string(s.Data))
	}
}

func TestErrNotFound(t *testing.T) {
	if !errors.Is(ErrNotFound, ErrNotFound) {
		t.Fatal("ErrNotFound should wrap itself")
	}
	if errors.Is(ErrNotFound, ErrInvalidInput) {
		t.Fatal("ErrNotFound should not match ErrInvalidInput")
	}
}

func TestErrInvalidInput(t *testing.T) {
	if !errors.Is(ErrInvalidInput, ErrInvalidInput) {
		t.Fatal("ErrInvalidInput should wrap itself")
	}
	if errors.Is(ErrInvalidInput, ErrNotFound) {
		t.Fatal("ErrInvalidInput should not match ErrNotFound")
	}
}
