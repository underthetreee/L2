package handler

import (
	"encoding/json"
	"log/slog"
	"net/http"

	"github.com/google/uuid"
)

type deleteEventParam struct {
	UserID uuid.UUID `json:"user_id"`
	ID     uuid.UUID `json:"id"`
}

func (h *EventHandler) handleDelete(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		slog.Error("delete response", "err", "method not allowed")
		sendErrorResponse(w, http.StatusMethodNotAllowed, "method not allowed")
		return
	}

	defer r.Body.Close()
	var param deleteEventParam
	if err := json.NewDecoder(r.Body).Decode(&param); err != nil {
		slog.Error("delete response", "err", err)
		sendErrorResponse(w, http.StatusBadRequest, "bad request")
		return
	}

	if err := h.service.Delete(param.UserID, param.ID); err != nil {
		slog.Error("delete response", "err", err)
		sendErrorResponse(w, http.StatusInternalServerError, "internal server error")
		return
	}

	resp := response{
		Result: "deleted",
	}
	sendJSONResponse(w, http.StatusOK, resp)
}
