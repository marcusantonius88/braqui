package telegram

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

type Logger interface {
	Error(msg string, fields map[string]any)
}

type Gateway struct {
	token   string
	client  *http.Client
	baseURL string
	log     Logger
}

func NewGateway(token string, log Logger) *Gateway {
	return &Gateway{
		token:   token,
		baseURL: fmt.Sprintf("https://api.telegram.org/bot%s", token),
		client: &http.Client{
			Timeout: 10 * time.Second,
		},
		log: log,
	}
}

type sendMessagePayload struct {
	ChatID int64  `json:"chat_id"`
	Text   string `json:"text"`
}

type telegramResponse struct {
	Ok          bool   `json:"ok"`
	Description string `json:"description,omitempty"`
}

func (g *Gateway) SendMessage(ctx context.Context, chatID int64, text string) error {
	payload := sendMessagePayload{ChatID: chatID, Text: text}
	body, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("marshal: %w", err)
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, g.baseURL+"/sendMessage", bytes.NewReader(body))
	if err != nil {
		return fmt.Errorf("create request: %w", err)
	}
	req.Header.Set("Content-Type", "application/json")

	resp, err := g.client.Do(req)
	if err != nil {
		return fmt.Errorf("send request: %w", err)
	}
	defer resp.Body.Close()

	respBody, _ := io.ReadAll(resp.Body)

	var result telegramResponse
	if err := json.Unmarshal(respBody, &result); err != nil {
		return fmt.Errorf("decode response: %w", err)
	}

	if !result.Ok {
		err := fmt.Errorf("telegram api error: %s", result.Description)
		if g.log != nil {
			g.log.Error("telegram api error", map[string]any{"description": result.Description, "chat_id": chatID})
		}
		return err
	}

	return nil
}
