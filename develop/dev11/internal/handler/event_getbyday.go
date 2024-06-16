package handler

import (
	"log/slog"
	"net/http"
	"time"

	"github.com/google/uuid"
)

func (h *EventHandler) handleGetByDay(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		slog.Error("getbyday response", "err", "method not allowed")
		sendErrorResponse(w, http.StatusMethodNotAllowed, "method not allowed")
		return
	}

	userID, err := uuid.Parse(r.URL.Query().Get("user_id"))
	if err != nil {
		slog.Error("getbyday response", "err", err)
		sendErrorResponse(w, http.StatusBadRequest, "bad request")
		return
	}

	date, err := time.Parse("2006-01-02", r.URL.Query().Get("date"))
	if err != nil {
		slog.Error("getbyday response", "err", err)
		sendErrorResponse(w, http.StatusBadRequest, "bad request")
		return
	}

	events, err := h.service.GetByDay(userID, date)
	if err != nil {
		slog.Error("getbyday response", "err", err)
		sendErrorResponse(w, http.StatusInternalServerError, "internal server error")
		return
	}
	resp := response{
		Result: events,
	}
	sendJSONResponse(w, http.StatusOK, resp)
}
