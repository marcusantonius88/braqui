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
	messages   []struct{ ChatID int64; Text string }
}

func (m *mockGateway) SendMessage(ctx context.Context, chatID int64, text string) error {
	m.lastChatID = chatID
	m.lastText = text
	m.messages = append(m.messages, struct{ ChatID int64; Text string }{chatID, text})
	return nil
}

type mockLogger struct {
	lastMsg  string
	lastLvl  string
}

func (m *mockLogger) Info(msg string, fields map[string]any) {
	m.lastMsg = msg
	m.lastLvl = "info"
}

func (m *mockLogger) Error(msg string, fields map[string]any) {
	m.lastMsg = msg
	m.lastLvl = "error"
}

func makeUpdate(t *testing.T, chatID int64, firstName, text string) []byte {
	t.Helper()
	upd := update{
		UpdateID: 1,
		Message: &message{
			MessageID: 1,
			From:      from{ID: chatID, FirstName: firstName},
			Chat:      chat{ID: chatID},
			Text:      text,
		},
	}
	b, _ := json.Marshal(upd)
	return b
}

func TestHandler_StartCommand(t *testing.T) {
	gw := &mockGateway{}
	h := NewHandler(gw, &mockLogger{})
	body := makeUpdate(t, 12345, "João", "/start")

	req := httptest.NewRequest(http.MethodPost, "/telegram/webhook", bytes.NewReader(body))
	rec := httptest.NewRecorder()

	h.ServeHTTP(rec, req)

	if rec.Code != http.StatusOK {
		t.Fatalf("expected 200, got %d", rec.Code)
	}

	if gw.lastChatID != 12345 {
		t.Fatalf("expected chat_id 12345, got %d", gw.lastChatID)
	}
	if gw.lastText == "" {
		t.Fatal("expected a reply message")
	}
	if !contains(gw.lastText, "João") {
		t.Fatalf("expected reply to mention João, got: %s", gw.lastText)
	}
}

func TestHandler_HelpCommand(t *testing.T) {
	gw := &mockGateway{}
	h := NewHandler(gw, &mockLogger{})
	body := makeUpdate(t, 12345, "Maria", "/help")

	req := httptest.NewRequest(http.MethodPost, "/telegram/webhook", bytes.NewReader(body))
	rec := httptest.NewRecorder()

	h.ServeHTTP(rec, req)

	if rec.Code != http.StatusOK {
		t.Fatalf("expected 200, got %d", rec.Code)
	}

	if !contains(gw.lastText, "/start") {
		t.Fatalf("expected help to mention /start, got: %s", gw.lastText)
	}
}

func TestHandler_UnknownCommand(t *testing.T) {
	gw := &mockGateway{}
	h := NewHandler(gw, &mockLogger{})
	body := makeUpdate(t, 12345, "Pedro", "some random text")

	req := httptest.NewRequest(http.MethodPost, "/telegram/webhook", bytes.NewReader(body))
	rec := httptest.NewRecorder()

	h.ServeHTTP(rec, req)

	if rec.Code != http.StatusOK {
		t.Fatalf("expected 200, got %d", rec.Code)
	}

	if gw.lastText != "" {
		t.Fatalf("expected no reply for unknown command, got: %s", gw.lastText)
	}
}

func TestHandler_NoMessage(t *testing.T) {
	gw := &mockGateway{}
	h := NewHandler(gw, &mockLogger{})

	upd := update{UpdateID: 1}
	b, _ := json.Marshal(upd)

	req := httptest.NewRequest(http.MethodPost, "/telegram/webhook", bytes.NewReader(b))
	rec := httptest.NewRecorder()

	h.ServeHTTP(rec, req)

	if rec.Code != http.StatusOK {
		t.Fatalf("expected 200, got %d", rec.Code)
	}
	if gw.lastText != "" {
		t.Fatal("expected no reply for update without message")
	}
}

func TestHandler_WrongMethod(t *testing.T) {
	h := NewHandler(&mockGateway{}, &mockLogger{})

	req := httptest.NewRequest(http.MethodGet, "/telegram/webhook", nil)
	rec := httptest.NewRecorder()

	h.ServeHTTP(rec, req)

	if rec.Code != http.StatusMethodNotAllowed {
		t.Fatalf("expected 405, got %d", rec.Code)
	}
}

func TestHandler_InvalidJSON(t *testing.T) {
	h := NewHandler(&mockGateway{}, &mockLogger{})

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
