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
}

func (m *mockGateway) SendMessage(ctx context.Context, chatID int64, text string) error {
	m.lastChatID = chatID
	m.lastText = text
	return nil
}

type mockIdentificer struct {
	userID string
	isNew  bool
	err    error
}

func (m *mockIdentificer) Identify(ctx context.Context, telegramID int64, firstName, username string) (string, bool, error) {
	return m.userID, m.isNew, m.err
}

type mockLogger struct{}

func (m *mockLogger) Info(msg string, fields map[string]any) {}
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

func TestHandler_ExistingUserStart(t *testing.T) {
	gw := &mockGateway{}
	users := &mockIdentificer{userID: "user-1", isNew: false}
	h := NewHandler(gw, users, &mockLogger{})

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
}

func TestHandler_NewUser(t *testing.T) {
	gw := &mockGateway{}
	users := &mockIdentificer{userID: "user-new", isNew: true}
	h := NewHandler(gw, users, &mockLogger{})

	body := makeUpdateWithUser(t, 67890, "Maria", "", "/start")
	req := httptest.NewRequest(http.MethodPost, "/telegram/webhook", bytes.NewReader(body))
	rec := httptest.NewRecorder()
	h.ServeHTTP(rec, req)

	if rec.Code != http.StatusOK {
		t.Fatalf("expected 200, got %d", rec.Code)
	}
	if gw.lastChatID != 67890 {
		t.Fatalf("expected chat_id 67890, got %d", gw.lastChatID)
	}
	if !contains(gw.lastText, "Maria") {
		t.Fatalf("expected welcome to mention Maria, got: %s", gw.lastText)
	}
	if !contains(gw.lastText, "nome do seu cão") {
		t.Fatalf("expected welcome to ask for pet name, got: %s", gw.lastText)
	}
}

func TestHandler_NewUserSendsWelcome(t *testing.T) {
	gw := &mockGateway{}
	users := &mockIdentificer{userID: "user-2", isNew: true}
	h := NewHandler(gw, users, &mockLogger{})

	body := makeUpdateWithUser(t, 111, "Ana", "aninha", "")
	req := httptest.NewRequest(http.MethodPost, "/telegram/webhook", bytes.NewReader(body))
	rec := httptest.NewRecorder()
	h.ServeHTTP(rec, req)

	if rec.Code != http.StatusOK {
		t.Fatalf("expected 200, got %d", rec.Code)
	}
	if gw.lastText == "" {
		t.Fatal("expected welcome message for new user")
	}
}

func TestHandler_UnknownCommand(t *testing.T) {
	gw := &mockGateway{}
	users := &mockIdentificer{userID: "user-1", isNew: false}
	h := NewHandler(gw, users, &mockLogger{})

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
	h := NewHandler(&mockGateway{}, &mockIdentificer{}, &mockLogger{})
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
	h := NewHandler(&mockGateway{}, &mockIdentificer{}, &mockLogger{})
	req := httptest.NewRequest(http.MethodGet, "/telegram/webhook", nil)
	rec := httptest.NewRecorder()
	h.ServeHTTP(rec, req)

	if rec.Code != http.StatusMethodNotAllowed {
		t.Fatalf("expected 405, got %d", rec.Code)
	}
}

func TestHandler_InvalidJSON(t *testing.T) {
	h := NewHandler(&mockGateway{}, &mockIdentificer{}, &mockLogger{})
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
