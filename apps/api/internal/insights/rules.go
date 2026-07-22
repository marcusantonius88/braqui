package insights

import (
	"time"

	"github.com/marcusantonius88/braqui/apps/api/internal/domain"
)

type Rule struct {
	Type       string
	EventType  string
	MinCount   int
	WindowDays int
	MessageFn  func(petName string, count int) string
	Absence    bool
}

var rules = []Rule{
	{
		Type:       InsightVomit,
		EventType:  EventVomit,
		MinCount:   3,
		WindowDays: 30,
		MessageFn: func(petName string, count int) string {
			return petName + " apresentou " + formatCount(count, "episódio", "episódios") + " de vômito nos últimos 30 dias."
		},
	},
	{
		Type:       InsightItching,
		EventType:  EventItching,
		MinCount:   3,
		WindowDays: 30,
		MessageFn: func(petName string, count int) string {
			return petName + " apresentou " + formatCount(count, "episódio", "episódios") + " de coceira nos últimos 30 dias."
		},
	},
	{
		Type:       InsightPanting,
		EventType:  EventPanting,
		MinCount:   5,
		WindowDays: 15,
		MessageFn: func(petName string, count int) string {
			return petName + " apresentou " + formatCount(count, "registro", "registros") + " de ofegância nos últimos 15 dias."
		},
	},
	{
		Type:       InsightMedication,
		EventType:  EventMedication,
		MinCount:   1,
		WindowDays: 60,
		Absence:    true,
		MessageFn: func(petName string, count int) string {
			return "Não encontrei registros recentes de medicação para " + petName + "."
		},
	},
}

func formatCount(count int, singular, plural string) string {
	if count == 1 {
		return "1 " + singular
	}
	return itoa(count) + " " + plural
}

func itoa(n int) string {
	if n == 0 {
		return "0"
	}
	var buf [12]byte
	i := len(buf)
	for n > 0 {
		i--
		buf[i] = byte('0' + n%10)
		n /= 10
	}
	return string(buf[i:])
}

func evaluateRule(petName string, events []*domain.Event, rule Rule) *Insight {
	now := time.Now()
	cutoff := now.AddDate(0, 0, -rule.WindowDays)

	count := 0
	for _, e := range events {
		if e.Type == rule.EventType && !e.Timestamp.Before(cutoff) {
			count++
		}
	}

	if rule.Absence {
		if count < rule.MinCount {
			return &Insight{
				Type:        rule.Type,
				Message:     rule.MessageFn(petName, count),
				GeneratedAt: now,
			}
		}
		return nil
	}

	if count >= rule.MinCount {
		return &Insight{
			Type:        rule.Type,
			Message:     rule.MessageFn(petName, count),
			GeneratedAt: now,
		}
	}

	return nil
}
