package health

import (
	"context"
	"encoding/json"
	"net/http"
)

type Pinger interface {
	Ping(ctx context.Context) error
}

type Logger interface {
	Error(msg string, fields map[string]any)
}

type Handler struct {
	db Pinger
	log Logger
}

func NewHandler(db Pinger, log Logger) *Handler {
	return &Handler{db: db, log: log}
}

func (h *Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	status := "ok"

	if h.db != nil {
		if err := h.db.Ping(r.Context()); err != nil {
			status = "degraded"
			if h.log != nil {
				h.log.Error("database ping failed", map[string]any{"error": err.Error()})
			}
		}
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"status": status})
}
