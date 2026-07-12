package logger

import (
	"bytes"
	"encoding/json"
	"strings"
	"testing"
)

func parseEntry(t *testing.T, buf *bytes.Buffer) map[string]any {
	t.Helper()
	var entry map[string]any
	if err := json.NewDecoder(buf).Decode(&entry); err != nil {
		t.Fatalf("decode: %v", err)
	}
	return entry
}

func TestInfo(t *testing.T) {
	var buf bytes.Buffer
	log := New(&buf, LevelInfo)
	log.Info("hello", nil)

	entry := parseEntry(t, &buf)
	if entry["level"] != "info" {
		t.Fatalf("expected info, got %v", entry["level"])
	}
	if entry["message"] != "hello" {
		t.Fatalf("expected hello, got %v", entry["message"])
	}
	if _, ok := entry["timestamp"]; !ok {
		t.Fatal("expected timestamp")
	}
}

func TestWarn(t *testing.T) {
	var buf bytes.Buffer
	log := New(&buf, LevelInfo)
	log.Warn("warning", nil)

	entry := parseEntry(t, &buf)
	if entry["level"] != "warn" {
		t.Fatalf("expected warn, got %v", entry["level"])
	}
	if entry["message"] != "warning" {
		t.Fatalf("expected warning, got %v", entry["message"])
	}
}

func TestError(t *testing.T) {
	var buf bytes.Buffer
	log := New(&buf, LevelInfo)
	log.Error("something failed", nil)

	entry := parseEntry(t, &buf)
	if entry["level"] != "error" {
		t.Fatalf("expected error, got %v", entry["level"])
	}
	if entry["message"] != "something failed" {
		t.Fatalf("expected something failed, got %v", entry["message"])
	}
}

func TestLevelFiltering(t *testing.T) {
	var buf bytes.Buffer
	log := New(&buf, LevelError)
	log.Info("should be hidden", nil)
	log.Warn("should be hidden too", nil)

	if buf.Len() > 0 {
		t.Fatal("expected no output for filtered levels")
	}

	log.Error("visible", nil)
	entry := parseEntry(t, &buf)
	if entry["message"] != "visible" {
		t.Fatalf("expected visible, got %v", entry["message"])
	}
}

func TestWithFields(t *testing.T) {
	var buf bytes.Buffer
	log := New(&buf, LevelInfo).With("module", "telegram")
	log.Info("message sent", map[string]any{"chat_id": 12345})

	entry := parseEntry(t, &buf)
	if entry["module"] != "telegram" {
		t.Fatalf("expected telegram, got %v", entry["module"])
	}
	if entry["chat_id"] != float64(12345) {
		t.Fatalf("expected 12345, got %v", entry["chat_id"])
	}
}

func TestWithIsImmutable(t *testing.T) {
	var parentBuf bytes.Buffer
	log := New(&parentBuf, LevelInfo)
	child := log.With("scope", "child")
	_ = child

	log.Info("parent", nil)

	e1 := parseEntry(t, &parentBuf)
	if _, ok := e1["scope"]; ok {
		t.Fatal("parent should not have scope field")
	}

	var childBuf bytes.Buffer
	New(&childBuf, LevelInfo).With("scope", "child").Info("child log", nil)
	e2 := parseEntry(t, &childBuf)
	if e2["scope"] != "child" {
		t.Fatalf("expected child scope, got %v", e2["scope"])
	}
}

func TestFields(t *testing.T) {
	var buf bytes.Buffer
	log := New(&buf, LevelInfo)
	log.Info("test", map[string]any{"key1": "value1", "key2": 42})

	entry := parseEntry(t, &buf)
	if entry["key1"] != "value1" {
		t.Fatalf("expected value1, got %v", entry["key1"])
	}
	if entry["key2"] != float64(42) {
		t.Fatalf("expected 42, got %v", entry["key2"])
	}
}

func TestNewlineDelimited(t *testing.T) {
	var buf bytes.Buffer
	log := New(&buf, LevelInfo)
	log.Info("first", nil)
	log.Info("second", nil)

	lines := strings.Split(strings.TrimSpace(buf.String()), "\n")
	if len(lines) != 2 {
		t.Fatalf("expected 2 lines, got %d", len(lines))
	}
}
