package handler

import (
	"calendar/internal/model"
	"encoding/json"
	"log/slog"
	"net/http"
	"time"

	"github.com/google/uuid"
)

type eventUpdateParam struct {
	UserID      uuid.UUID `json:"user_id"`
	ID          uuid.UUID `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Date        time.Time `json:"date"`
}

func (h *EventHandler) handleUpdate(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		slog.Error("update response", "err", "method not allowed")
		sendErrorResponse(w, http.StatusMethodNotAllowed, "method not allowed")
		return
	}

	defer r.Body.Close()
	var param eventUpdateParam
	if err := json.NewDecoder(r.Body).Decode(&param); err != nil {
		slog.Error("update response", "err", err)
		sendErrorResponse(w, http.StatusBadRequest, "bad request")
		return
	}

	event := model.Event{
		ID:          param.ID,
		Title:       param.Title,
		Description: param.Description,
		Date:        param.Date,
	}

	if err := event.Validate(); err != nil {
		slog.Error("update response", "err", err)
		sendErrorResponse(w, http.StatusBadRequest, "bad request")
		return
	}

	if err := h.service.Update(param.UserID, event); err != nil {
		slog.Error("update response", "err", err)
		sendErrorResponse(w, http.StatusInternalServerError, "internal server error")
		return
	}

	resp := response{
		Result: event,
	}
	sendJSONResponse(w, http.StatusOK, resp)
}
