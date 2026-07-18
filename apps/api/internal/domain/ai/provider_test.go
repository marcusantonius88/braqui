package ai

import (
	"testing"
)

func TestInterpretationResult_Fields(t *testing.T) {
	r := &InterpretationResult{
		Type:       "vomit",
		Confidence: "high",
		Payload:    map[string]any{"note": "depois do jantar"},
	}
	if r.Type != "vomit" {
		t.Fatalf("expected vomit, got %s", r.Type)
	}
	if r.Confidence != "high" {
		t.Fatalf("expected high, got %s", r.Confidence)
	}
	if r.Payload["note"] != "depois do jantar" {
		t.Fatalf("expected note, got %v", r.Payload["note"])
	}
}

func TestInterpretationResult_EmptyPayload(t *testing.T) {
	r := &InterpretationResult{Type: "fatigue", Confidence: "medium", Payload: map[string]any{}}
	if r.Payload == nil {
		t.Fatal("expected non-nil payload")
	}
}
