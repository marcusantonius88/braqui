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
		Source:      "manual",
		Timestamp:   time.Now(),
	}
	if e.Type != "feeding" {
		t.Fatalf("expected feeding, got %s", e.Type)
	}
	if e.Description != "comeu ração" {
		t.Fatalf("expected comeu ração, got %s", e.Description)
	}
	if e.Source != "manual" {
		t.Fatalf("expected manual, got %s", e.Source)
	}
}

func TestReminderCreation(t *testing.T) {
	r := &Reminder{
		PetID:       "pet-1",
		Title:       "Dar Simparic",
		Description: "antipulga",
		DueDate:     time.Now(),
		Status:      ReminderStatusPending,
	}
	if r.Title != "Dar Simparic" {
		t.Fatalf("expected Dar Simparic, got %s", r.Title)
	}
	if r.Status != ReminderStatusPending {
		t.Fatalf("expected pending, got %s", r.Status)
	}
}

func TestConversationStateCreation(t *testing.T) {
	s := &ConversationState{
		UserID:  "user-1",
		Flow:    "register_pet",
		Step:    "ask_name",
		Payload: []byte(`{"name":"Thor"}`),
	}
	if s.Flow != "register_pet" {
		t.Fatalf("expected register_pet, got %s", s.Flow)
	}
	if s.Step != "ask_name" {
		t.Fatalf("expected ask_name, got %s", s.Step)
	}
	if string(s.Payload) != `{"name":"Thor"}` {
		t.Fatalf("expected payload, got %s", string(s.Payload))
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
