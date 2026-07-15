package telegram

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

type mockGateway struct {
	lastChatID int64
	lastText   string
	messages   []string
}

func (m *mockGateway) SendMessage(ctx context.Context, chatID int64, text string) error {
	m.lastChatID = chatID
	m.lastText = text
	m.messages = append(m.messages, text)
	return nil
}

type mockIdentificer struct {
	userID string
}

func (m *mockIdentificer) Identify(ctx context.Context, telegramID int64, firstName, username string) (string, bool, error) {
	return m.userID, false, nil
}

type mockOnboarder struct {
	reply string
	err   error
}

func (m *mockOnboarder) Process(ctx context.Context, userID, text string) (string, error) {
	return m.reply, m.err
}

type mockLogger struct{}

func (m *mockLogger) Info(msg string, fields map[string]any)  {}
func (m *mockLogger) Error(msg string, fields map[string]any) {}

func makeUpdateWithUser(t *testing.T, chatID int64, firstName, username, text string) []byte {
	t.Helper()
	upd := update{
		UpdateID: 1,
		Message: &message{
			MessageID: 1,
			From:      from{ID: chatID, FirstName: firstName, Username: username},
			Chat:      chat{ID: chatID},
			Text:      text,
		},
	}
	b, _ := json.Marshal(upd)
	return b
}

func TestHandler_OnboardingFlow(t *testing.T) {
	gw := &mockGateway{}
	onboarder := &mockOnboarder{reply: "Qual a raça do Thor?"}
	h := NewHandler(gw, &mockIdentificer{userID: "user-1"}, onboarder, &mockLogger{})

	body := makeUpdateWithUser(t, 12345, "João", "joaobot", "Thor")
	req := httptest.NewRequest(http.MethodPost, "/telegram/webhook", bytes.NewReader(body))
	rec := httptest.NewRecorder()
	h.ServeHTTP(rec, req)

	if rec.Code != http.StatusOK {
		t.Fatalf("expected 200, got %d", rec.Code)
	}
	if gw.lastText != "Qual a raça do Thor?" {
		t.Fatalf("expected onboarding reply, got: %s", gw.lastText)
	}
}

func TestHandler_OnboardingComplete(t *testing.T) {
	gw := &mockGateway{}
	onboarder := &mockOnboarder{reply: "Perfeito 🐶"}
	h := NewHandler(gw, &mockIdentificer{userID: "user-1"}, onboarder, &mockLogger{})

	body := makeUpdateWithUser(t, 12345, "João", "", "João Pessoa")
	req := httptest.NewRequest(http.MethodPost, "/telegram/webhook", bytes.NewReader(body))
	rec := httptest.NewRecorder()
	h.ServeHTTP(rec, req)

	if rec.Code != http.StatusOK {
		t.Fatalf("expected 200, got %d", rec.Code)
	}
	if !contains(gw.lastText, "Perfeito") {
		t.Fatalf("expected completion message, got: %s", gw.lastText)
	}
}

func TestHandler_ExistingUserStart(t *testing.T) {
	gw := &mockGateway{}
	onboarder := &mockOnboarder{} // empty reply = has pet
	h := NewHandler(gw, &mockIdentificer{userID: "user-1"}, onboarder, &mockLogger{})

	body := makeUpdateWithUser(t, 12345, "João", "joaobot", "/start")
	req := httptest.NewRequest(http.MethodPost, "/telegram/webhook", bytes.NewReader(body))
	rec := httptest.NewRecorder()
	h.ServeHTTP(rec, req)

	if rec.Code != http.StatusOK {
		t.Fatalf("expected 200, got %d", rec.Code)
	}
	if gw.lastChatID != 12345 {
		t.Fatalf("expected chat_id 12345, got %d", gw.lastChatID)
	}
	if !contains(gw.lastText, "/help") {
		t.Fatalf("expected /start reply to mention /help, got: %s", gw.lastText)
	}
}

func TestHandler_UnknownCommand(t *testing.T) {
	gw := &mockGateway{}
	onboarder := &mockOnboarder{} // empty reply = has pet
	h := NewHandler(gw, &mockIdentificer{userID: "user-1"}, onboarder, &mockLogger{})

	body := makeUpdateWithUser(t, 12345, "Pedro", "", "random text")
	req := httptest.NewRequest(http.MethodPost, "/telegram/webhook", bytes.NewReader(body))
	rec := httptest.NewRecorder()
	h.ServeHTTP(rec, req)

	if rec.Code != http.StatusOK {
		t.Fatalf("expected 200, got %d", rec.Code)
	}
	if gw.lastText != "" {
		t.Fatalf("expected no reply, got: %s", gw.lastText)
	}
}

func TestHandler_NoMessage(t *testing.T) {
	h := NewHandler(&mockGateway{}, &mockIdentificer{}, &mockOnboarder{}, &mockLogger{})
	upd := update{UpdateID: 1}
	b, _ := json.Marshal(upd)

	req := httptest.NewRequest(http.MethodPost, "/telegram/webhook", bytes.NewReader(b))
	rec := httptest.NewRecorder()
	h.ServeHTTP(rec, req)

	if rec.Code != http.StatusOK {
		t.Fatalf("expected 200, got %d", rec.Code)
	}
}

func TestHandler_WrongMethod(t *testing.T) {
	h := NewHandler(&mockGateway{}, &mockIdentificer{}, &mockOnboarder{}, &mockLogger{})
	req := httptest.NewRequest(http.MethodGet, "/telegram/webhook", nil)
	rec := httptest.NewRecorder()
	h.ServeHTTP(rec, req)

	if rec.Code != http.StatusMethodNotAllowed {
		t.Fatalf("expected 405, got %d", rec.Code)
	}
}

func TestHandler_InvalidJSON(t *testing.T) {
	h := NewHandler(&mockGateway{}, &mockIdentificer{}, &mockOnboarder{}, &mockLogger{})
	req := httptest.NewRequest(http.MethodPost, "/telegram/webhook", bytes.NewReader([]byte(`{invalid`)))
	rec := httptest.NewRecorder()
	h.ServeHTTP(rec, req)

	if rec.Code != http.StatusBadRequest {
		t.Fatalf("expected 400, got %d", rec.Code)
	}
}

func contains(s, substr string) bool {
	return len(s) >= len(substr) && containsStr(s, substr)
}

func containsStr(s, substr string) bool {
	for i := 0; i <= len(s)-len(substr); i++ {
		if s[i:i+len(substr)] == substr {
			return true
		}
	}
	return false
}
