package ai

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	domainai "github.com/marcusantonius88/braqui/apps/api/internal/domain/ai"
)

type mockLogger struct{}

func (mockLogger) Info(msg string, fields map[string]any)  {}
func (mockLogger) Error(msg string, fields map[string]any) {}

func setEndpoint(t *testing.T, url string) {
	t.Helper()
	old := geminiEndpoint
	geminiEndpoint = url
	t.Cleanup(func() { geminiEndpoint = old })
}

func TestGeminiProvider_Interpret(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			t.Fatalf("expected POST, got %s", r.Method)
		}
		if r.URL.Query().Get("key") != "test-key" {
			t.Fatalf("expected test-key, got %s", r.URL.Query().Get("key"))
		}
		resp := `{"candidates":[{"content":{"parts":[{"text":"{\"type\":\"vomit\",\"confidence\":\"high\",\"payload\":{}}"}]}}]}`
		w.Write([]byte(resp))
	}))
	defer server.Close()
	setEndpoint(t, server.URL)

	p := &GeminiProvider{
		apiKey: "test-key",
		client: server.Client(),
		log:    mockLogger{},
	}

	result, err := p.Interpret(context.Background(), "Thor vomitou")
	if err != nil {
		t.Fatalf("interpret: %v", err)
	}
	if result.Type != "vomit" {
		t.Fatalf("expected vomit, got %s", result.Type)
	}
	if result.Confidence != "high" {
		t.Fatalf("expected high, got %s", result.Confidence)
	}
}

func TestGeminiProvider_InvalidJSON(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(`invalid`))
	}))
	defer server.Close()
	setEndpoint(t, server.URL)

	p := &GeminiProvider{
		apiKey: "test-key",
		client: server.Client(),
		log:    mockLogger{},
	}

	result, err := p.Interpret(context.Background(), "test")
	if err != nil {
		t.Fatalf("interpret: %v", err)
	}
	if result.Type != "NOT_INTERPRETED" {
		t.Fatalf("expected NOT_INTERPRETED, got %s", result.Type)
	}
}

func TestGeminiProvider_EmptyCandidates(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(`{"candidates":[]}`))
	}))
	defer server.Close()
	setEndpoint(t, server.URL)

	p := &GeminiProvider{
		apiKey: "test-key",
		client: server.Client(),
		log:    mockLogger{},
	}

	result, err := p.Interpret(context.Background(), "test")
	if err != nil {
		t.Fatalf("interpret: %v", err)
	}
	if result.Type != "NOT_INTERPRETED" {
		t.Fatalf("expected NOT_INTERPRETED, got %s", result.Type)
	}
}

func TestGeminiProvider_NonOKStatus(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusTooManyRequests)
		w.Write([]byte(`{"error":{"message":"rate limited"}}`))
	}))
	defer server.Close()
	setEndpoint(t, server.URL)

	p := &GeminiProvider{
		apiKey: "test-key",
		client: server.Client(),
		log:    mockLogger{},
	}

	result, err := p.Interpret(context.Background(), "test")
	if err != nil {
		t.Fatalf("interpret: %v", err)
	}
	if result.Type != "NOT_INTERPRETED" {
		t.Fatalf("expected NOT_INTERPRETED, got %s", result.Type)
	}
}

func TestGeminiProvider_Timeout(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		<-r.Context().Done()
	}))
	defer server.Close()
	setEndpoint(t, server.URL)

	p := &GeminiProvider{
		apiKey: "test-key",
		client: server.Client(),
		log:    mockLogger{},
	}

	ctx, cancel := context.WithCancel(context.Background())
	cancel()

	result, err := p.Interpret(ctx, "test")
	if err != nil {
		t.Fatalf("interpret: %v", err)
	}
	if result.Type != "NOT_INTERPRETED" {
		t.Fatalf("expected NOT_INTERPRETED, got %s", result.Type)
	}
}

func TestParseResult(t *testing.T) {
	tests := []struct {
		input string
		want  string
	}{
		{`{"type":"vomit","confidence":"high"}`, "vomit"},
		{`{"type":"fatigue","confidence":"medium"}`, "fatigue"},
		{`{"type":"medication_given","confidence":"low"}`, "medication_given"},
		{`{"type":"vet_visit","confidence":"high","payload":{"note":"checkup"}}`, "vet_visit"},
		{`Texto qualquer {"type":"diarrhea","confidence":"high"} depois`, "diarrhea"},
		{`{"type":"cough","confidence":"high"}`, "cough"},
	}

	for _, tt := range tests {
		result, err := parseResult(tt.input)
		if err != nil {
			t.Fatalf("parseResult(%q): %v", tt.input, err)
		}
		if result.Type != tt.want {
			t.Fatalf("parseResult(%q): expected %s, got %s", tt.input, tt.want, result.Type)
		}
	}
}

func TestParseResult_Invalid(t *testing.T) {
	tests := []string{
		"not json",
		`{"type":""}`,
		`{}`,
	}

	for _, input := range tests {
		_, err := parseResult(input)
		if err == nil {
			t.Fatalf("parseResult(%q): expected error", input)
		}
	}
}

func TestCleanJSON(t *testing.T) {
	tests := []struct {
		input string
		want  string
	}{
		{`{"a":1}`, `{"a":1}`},
		{`text {"a":1} more`, `{"a":1}`},
		{`{"a":{"b":2}}`, `{"a":{"b":2}}`},
		{`no braces`, `{}`},
	}

	for _, tt := range tests {
		got := cleanJSON(tt.input)
		if got != tt.want {
			t.Fatalf("cleanJSON(%q): expected %q, got %q", tt.input, tt.want, got)
		}
	}
}

func TestNotInterpreted(t *testing.T) {
	r := notInterpreted()
	if r.Type != "NOT_INTERPRETED" {
		t.Fatalf("expected NOT_INTERPRETED, got %s", r.Type)
	}
	if r.Confidence != "low" {
		t.Fatalf("expected low, got %s", r.Confidence)
	}
}

func TestInterpretationResult_JSON(t *testing.T) {
	raw := `{"type":"fatigue","confidence":"medium","payload":{}}`
	var r domainai.InterpretationResult
	if err := json.Unmarshal([]byte(raw), &r); err != nil {
		t.Fatalf("unmarshal: %v", err)
	}
	if r.Type != "fatigue" {
		t.Fatalf("expected fatigue, got %s", r.Type)
	}
}
