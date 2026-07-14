package telegram

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

type testLogger struct {
	lastMsg string
}

func (t *testLogger) Error(msg string, fields map[string]any) {
	t.lastMsg = msg
}

func TestGateway_SendMessage(t *testing.T) {
	var received struct {
		ChatID int64  `json:"chat_id"`
		Text   string `json:"text"`
	}

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			t.Fatalf("expected POST, got %s", r.Method)
		}
		if r.URL.Path != "/botTESTTOKEN/sendMessage" {
			t.Fatalf("expected /botTESTTOKEN/sendMessage, got %s", r.URL.Path)
		}
		json.NewDecoder(r.Body).Decode(&received)
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte(`{"ok":true}`))
	}))
	defer server.Close()

	gw := &Gateway{
		token:   "TESTTOKEN",
		baseURL: server.URL + "/botTESTTOKEN",
		client:  server.Client(),
		log:     &testLogger{},
	}

	err := gw.SendMessage(context.Background(), 12345, "Hello World")
	if err != nil {
		t.Fatalf("send message: %v", err)
	}

	if received.ChatID != 12345 {
		t.Fatalf("expected chat_id 12345, got %d", received.ChatID)
	}
	if received.Text != "Hello World" {
		t.Fatalf("expected Hello World, got %s", received.Text)
	}
}

func TestGateway_SendMessageError(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte(`{"ok":false,"description":"Unauthorized"}`))
	}))
	defer server.Close()

	gw := &Gateway{
		baseURL: server.URL + "/botTESTTOKEN",
		client:  server.Client(),
		log:     &testLogger{},
	}

	err := gw.SendMessage(context.Background(), 12345, "Hello")
	if err == nil {
		t.Fatal("expected error, got nil")
	}
	if err.Error() != "telegram api error: Unauthorized" {
		t.Fatalf("expected Unauthorized error, got: %v", err)
	}
}

func TestGateway_SendMessageTimeout(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), 50*time.Millisecond)
	defer cancel()

	gw := &Gateway{
		baseURL: "http://localhost:19999",
		client:  &http.Client{Timeout: 50 * time.Millisecond},
		log:     &testLogger{},
	}

	err := gw.SendMessage(ctx, 12345, "Hello")
	if err == nil {
		t.Fatal("expected error, got nil")
	}
}
