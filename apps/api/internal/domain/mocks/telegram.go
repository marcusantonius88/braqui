package mocks

import (
	"context"
	"sync"
)

type TelegramGateway struct {
	mu       sync.Mutex
	Messages []struct{ ChatID int64; Text string }
}

func NewTelegramGateway() *TelegramGateway {
	return &TelegramGateway{}
}

func (g *TelegramGateway) SendMessage(ctx context.Context, chatID int64, text string) error {
	g.mu.Lock()
	defer g.mu.Unlock()
	g.Messages = append(g.Messages, struct{ ChatID int64; Text string }{chatID, text})
	return nil
}
