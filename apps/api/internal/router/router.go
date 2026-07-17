package router

import (
	"context"
	"fmt"

	"github.com/marcusantonius88/braqui/apps/api/internal/domain"
)

type StateManager interface {
	Get(ctx context.Context, userID string) (*domain.ConversationState, error)
}

type FlowHandler interface {
	Handle(ctx context.Context, userID, text string) (reply string, err error)
}

type CommandHandler interface {
	Handle(ctx context.Context, userID string) (reply string, err error)
}

type Logger interface {
	Info(msg string, fields map[string]any)
	Error(msg string, fields map[string]any)
}

type Router struct {
	states   StateManager
	flows    map[string]FlowHandler
	commands map[string]CommandHandler
	fallback FlowHandler
	log      Logger
}

func NewRouter(states StateManager, log Logger) *Router {
	return &Router{
		states:   states,
		flows:    make(map[string]FlowHandler),
		commands: make(map[string]CommandHandler),
		log:      log,
	}
}

func (r *Router) RegisterFlow(name string, handler FlowHandler) {
	r.flows[name] = handler
}

func (r *Router) RegisterCommand(name string, handler CommandHandler) {
	r.commands[name] = handler
}

func (r *Router) SetDefault(handler FlowHandler) {
	r.fallback = handler
}

type FlowHandlerFunc func(ctx context.Context, userID, text string) (string, error)

func (f FlowHandlerFunc) Handle(ctx context.Context, userID, text string) (string, error) {
	return f(ctx, userID, text)
}

type CommandHandlerFunc func(ctx context.Context, userID string) (string, error)

func (f CommandHandlerFunc) Handle(ctx context.Context, userID string) (string, error) {
	return f(ctx, userID)
}

func (r *Router) Route(ctx context.Context, userID, text string) (string, error) {
	state, err := r.states.Get(ctx, userID)
	if err != nil {
		r.log.Error("failed to get state for routing", map[string]any{"user_id": userID, "error": err.Error()})
		return "", fmt.Errorf("get state: %w", err)
	}

	if state != nil && state.Flow != "" {
		if handler, ok := r.flows[state.Flow]; ok {
			r.log.Info("route by flow", map[string]any{"user_id": userID, "flow": state.Flow})
			return handler.Handle(ctx, userID, text)
		}
		r.log.Error("unknown flow", map[string]any{"user_id": userID, "flow": state.Flow})
	}

	if handler, ok := r.commands[text]; ok {
		r.log.Info("route by command", map[string]any{"user_id": userID, "command": text})
		return handler.Handle(ctx, userID)
	}

	if r.fallback != nil {
		r.log.Info("route fallback", map[string]any{"user_id": userID})
		return r.fallback.Handle(ctx, userID, text)
	}

	return "", nil
}
