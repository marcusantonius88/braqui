package telegram

import (
	"context"
	"encoding/json"
	"io"
	"net/http"
)

type TelegramGateway interface {
	SendMessage(ctx context.Context, chatID int64, text string) error
}

type Logger interface {
	Info(msg string, fields map[string]any)
	Error(msg string, fields map[string]any)
}

type Handler struct {
	gateway TelegramGateway
	log     Logger
}

func NewHandler(gateway TelegramGateway, log Logger) *Handler {
	return &Handler{gateway: gateway, log: log}
}

type update struct {
	UpdateID int64   `json:"update_id"`
	Message  *message `json:"message,omitempty"`
}

type message struct {
	MessageID int64   `json:"message_id"`
	From      from    `json:"from"`
	Chat      chat    `json:"chat"`
	Text      string  `json:"text"`
}

type from struct {
	ID        int64  `json:"id"`
	FirstName string `json:"first_name"`
}

type chat struct {
	ID int64 `json:"id"`
}

func (h *Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}

	body, err := io.ReadAll(r.Body)
	if err != nil {
		h.log.Error("failed to read webhook body", map[string]any{"error": err.Error()})
		http.Error(w, "bad request", http.StatusBadRequest)
		return
	}

	var upd update
	if err := json.Unmarshal(body, &upd); err != nil {
		h.log.Error("failed to parse webhook payload", map[string]any{"error": err.Error()})
		http.Error(w, "bad request", http.StatusBadRequest)
		return
	}

	if upd.Message == nil {
		http.Error(w, "ok", http.StatusOK)
		return
	}

	chatID := upd.Message.Chat.ID
	telegramID := upd.Message.From.ID
	firstName := upd.Message.From.FirstName
	text := upd.Message.Text

	h.log.Info("webhook received", map[string]any{"chat_id": chatID, "text": text, "from": firstName})

	reply := h.processMessage(telegramID, firstName, text)

	if reply != "" {
		if err := h.gateway.SendMessage(r.Context(), chatID, reply); err != nil {
			h.log.Error("failed to send message", map[string]any{"chat_id": chatID, "error": err.Error()})
		}
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"status": "ok"})
}

func (h *Handler) processMessage(telegramID int64, firstName, text string) string {
	switch text {
	case "/start":
		return "Olá, " + firstName + "! Eu sou o Braqui, seu assistente para cuidar da saúde do seu cão braquicefálico. Use /help para ver os comandos disponíveis."
	case "/help":
		return "Comandos disponíveis:\n/start - Iniciar conversa\n/help - Mostrar esta ajuda"
	default:
		return ""
	}
}
