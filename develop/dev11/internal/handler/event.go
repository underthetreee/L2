package handler

import (
	"calendar/internal/model"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/google/uuid"
)

type eventCreateParam struct {
	UserID      uuid.UUID `json:"user_id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Date        time.Time `json:"date"`
}

type eventServicer interface {
	Add(uuid.UUID, model.Event) error
	Update(uuid.UUID, uuid.UUID, model.Event) error
}

type EventHandler struct {
	service eventServicer
}

func NewEventHandler(service eventServicer) *EventHandler {
	return &EventHandler{
		service: service,
	}
}

func (h *EventHandler) handleCreate(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		fmt.Println("method")
		return
	}

	defer r.Body.Close()
	var param eventCreateParam
	if err := json.NewDecoder(r.Body).Decode(&param); err != nil {
		fmt.Println("decoder")
		return
	}

	event := model.Event{
		ID:          uuid.New(),
		Title:       param.Title,
		Description: param.Description,
		Date:        param.Date,
	}

	if err := h.service.Add(param.UserID, event); err != nil {
		fmt.Println("add service")
		return
	}
	fmt.Println(event)
}
