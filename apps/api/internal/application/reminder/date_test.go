package reminder

import (
	"testing"
	"time"
)

func TestParseDate_Hoje(t *testing.T) {
	now := time.Date(2026, 7, 22, 10, 0, 0, 0, time.UTC)
	d, err := parseDate("hoje", now)
	if err != nil {
		t.Fatalf("parse: %v", err)
	}
	expected := time.Date(2026, 7, 22, 0, 0, 0, 0, time.UTC)
	if !d.Equal(expected) {
		t.Fatalf("expected %v, got %v", expected, d)
	}
}

func TestParseDate_Amanha(t *testing.T) {
	now := time.Date(2026, 7, 22, 10, 0, 0, 0, time.UTC)
	d, err := parseDate("amanhã", now)
	if err != nil {
		t.Fatalf("parse: %v", err)
	}
	expected := time.Date(2026, 7, 23, 0, 0, 0, 0, time.UTC)
	if !d.Equal(expected) {
		t.Fatalf("expected %v, got %v", expected, d)
	}
}

func TestParseDate_AmanhaSemAcento(t *testing.T) {
	now := time.Date(2026, 7, 22, 10, 0, 0, 0, time.UTC)
	d, err := parseDate("amanha", now)
	if err != nil {
		t.Fatalf("parse: %v", err)
	}
	expected := time.Date(2026, 7, 23, 0, 0, 0, 0, time.UTC)
	if !d.Equal(expected) {
		t.Fatalf("expected %v, got %v", expected, d)
	}
}

func TestParseDate_DaquiA30Dias(t *testing.T) {
	now := time.Date(2026, 7, 22, 10, 0, 0, 0, time.UTC)
	d, err := parseDate("daqui a 30 dias", now)
	if err != nil {
		t.Fatalf("parse: %v", err)
	}
	expected := time.Date(2026, 8, 21, 0, 0, 0, 0, time.UTC)
	if !d.Equal(expected) {
		t.Fatalf("expected %v, got %v", expected, d)
	}
}

func TestParseDate_DaquiA7Dias(t *testing.T) {
	now := time.Date(2026, 7, 22, 10, 0, 0, 0, time.UTC)
	d, err := parseDate("daqui a 7 dias", now)
	if err != nil {
		t.Fatalf("parse: %v", err)
	}
	expected := time.Date(2026, 7, 29, 0, 0, 0, 0, time.UTC)
	if !d.Equal(expected) {
		t.Fatalf("expected %v, got %v", expected, d)
	}
}

func TestParseDate_Em5Dias(t *testing.T) {
	now := time.Date(2026, 7, 22, 10, 0, 0, 0, time.UTC)
	d, err := parseDate("em 5 dias", now)
	if err != nil {
		t.Fatalf("parse: %v", err)
	}
	expected := time.Date(2026, 7, 27, 0, 0, 0, 0, time.UTC)
	if !d.Equal(expected) {
		t.Fatalf("expected %v, got %v", expected, d)
	}
}

func TestParseDate_Dia15(t *testing.T) {
	now := time.Date(2026, 7, 22, 10, 0, 0, 0, time.UTC)
	d, err := parseDate("dia 15", now)
	if err != nil {
		t.Fatalf("parse: %v", err)
	}
	expected := time.Date(2026, 8, 15, 0, 0, 0, 0, time.UTC)
	if !d.Equal(expected) {
		t.Fatalf("expected %v (august), got %v (july?)", expected, d)
	}
}

func TestParseDate_Dia25_MonthPassed(t *testing.T) {
	now := time.Date(2026, 7, 22, 10, 0, 0, 0, time.UTC)
	d, err := parseDate("dia 25", now)
	if err != nil {
		t.Fatalf("parse: %v", err)
	}
	expected := time.Date(2026, 7, 25, 0, 0, 0, 0, time.UTC)
	if !d.Equal(expected) {
		t.Fatalf("expected %v, got %v", expected, d)
	}
}

func TestParseDate_Dia1(t *testing.T) {
	now := time.Date(2026, 7, 22, 10, 0, 0, 0, time.UTC)
	d, err := parseDate("dia 1", now)
	if err != nil {
		t.Fatalf("parse: %v", err)
	}
	expected := time.Date(2026, 8, 1, 0, 0, 0, 0, time.UTC)
	if !d.Equal(expected) {
		t.Fatalf("expected %v (august), got %v", expected, d)
	}
}

func TestParseDate_Invalid(t *testing.T) {
	now := time.Date(2026, 7, 22, 10, 0, 0, 0, time.UTC)
	_, err := parseDate("qualquer coisa", now)
	if err == nil {
		t.Fatal("expected error for invalid date")
	}
}

func TestParseDate_Dia32_Invalid(t *testing.T) {
	now := time.Date(2026, 7, 22, 10, 0, 0, 0, time.UTC)
	_, err := parseDate("dia 32", now)
	if err == nil {
		t.Fatal("expected error for invalid day 32")
	}
}
