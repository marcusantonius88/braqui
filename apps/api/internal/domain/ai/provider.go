package ai

import "context"

type InterpretationResult struct {
	Type       string
	Confidence string
	Payload    map[string]any
}

type AIProvider interface {
	Interpret(ctx context.Context, message string) (*InterpretationResult, error)
}
