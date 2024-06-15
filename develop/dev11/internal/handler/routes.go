package handler

import (
	"calendar/internal/middleware"
	"net/http"
)

func InitRoutes(mux *http.ServeMux, service eventServicer) {
	eventHandler := NewEventHandler(service)

	mux.HandleFunc("POST /create_event", middleware.Log(eventHandler.handleCreate))
}
