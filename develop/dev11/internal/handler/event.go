package handler

import (
	"calendar/internal/model"
	"encoding/json"
	"log/slog"
	"net/http"
	"time"

	"github.com/google/uuid"
)

type response struct {
	Result any    `json:"result,omitempty"`
	Error  string `json:"error,omitempty"`
}

type eventServicer interface {
	Add(uuid.UUID, model.Event) error
	Update(uuid.UUID, model.Event) error
	Delete(uuid.UUID, uuid.UUID) error
	GetByDay(uuid.UUID, time.Time) ([]model.Event, error)
	GetByWeek(uuid.UUID, time.Time) ([]model.Event, error)
	GetByMonth(uuid.UUID, time.Time) ([]model.Event, error)
}

type EventHandler struct {
	service eventServicer
}

func NewEventHandler(service eventServicer) *EventHandler {
	return &EventHandler{
		service: service,
	}
}

func sendJSONResponse(w http.ResponseWriter, statusCode int, data any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	if err := json.NewEncoder(w).Encode(data); err != nil {
		slog.Error("json response", "err", err)
	}
}

func sendErrorResponse(w http.ResponseWriter, statusCode int, message string) {
	resp := response{
		Error: message,
	}
	sendJSONResponse(w, statusCode, resp)
}
