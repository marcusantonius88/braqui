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

type UserIdentificer interface {
	Identify(ctx context.Context, telegramID int64, firstName, username string) (userID string, isNew bool, err error)
}

type Router interface {
	Route(ctx context.Context, userID, text string) (reply string, err error)
}

type Logger interface {
	Info(msg string, fields map[string]any)
	Error(msg string, fields map[string]any)
}

type Handler struct {
	gateway TelegramGateway
	users   UserIdentificer
	router  Router
	log     Logger
}

func NewHandler(gateway TelegramGateway, users UserIdentificer, router Router, log Logger) *Handler {
	return &Handler{gateway: gateway, users: users, router: router, log: log}
}

type update struct {
	UpdateID int64    `json:"update_id"`
	Message  *message `json:"message,omitempty"`
}

type message struct {
	MessageID int64  `json:"message_id"`
	From      from   `json:"from"`
	Chat      chat   `json:"chat"`
	Text      string `json:"text"`
}

type from struct {
	ID        int64  `json:"id"`
	FirstName string `json:"first_name"`
	Username  string `json:"username,omitempty"`
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
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(map[string]string{"status": "ok"})
		return
	}

	chatID := upd.Message.Chat.ID
	telegramID := upd.Message.From.ID
	firstName := upd.Message.From.FirstName
	username := upd.Message.From.Username
	text := upd.Message.Text

	h.log.Info("webhook received", map[string]any{"chat_id": chatID, "text": text, "from": firstName})

	userID, _, err := h.users.Identify(r.Context(), telegramID, firstName, username)
	if err != nil {
		h.log.Error("failed to identify user", map[string]any{"chat_id": chatID, "error": err.Error()})
		http.Error(w, "internal error", http.StatusInternalServerError)
		return
	}

	reply, err := h.router.Route(r.Context(), userID, text)
	if err != nil {
		h.log.Error("router error", map[string]any{"user_id": userID, "error": err.Error()})
		http.Error(w, "internal error", http.StatusInternalServerError)
		return
	}

	if reply != "" {
		if err := h.gateway.SendMessage(r.Context(), chatID, reply); err != nil {
			h.log.Error("failed to send message", map[string]any{"chat_id": chatID, "error": err.Error()})
		}
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"status": "ok"})
}
