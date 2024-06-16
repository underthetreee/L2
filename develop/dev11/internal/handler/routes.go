package handler

import (
	"calendar/internal/middleware"
	"net/http"
)

func InitRoutes(mux *http.ServeMux, service eventServicer) {
	eventHandler := NewEventHandler(service)

	mux.HandleFunc("POST /create_event", middleware.Log(eventHandler.handleCreate))
	mux.HandleFunc("POST /update_event", middleware.Log(eventHandler.handleUpdate))
	mux.HandleFunc("POST /delete_event", middleware.Log(eventHandler.handleDelete))
	mux.HandleFunc("GET /events_for_day", middleware.Log(eventHandler.handleGetByDay))
	mux.HandleFunc("GET /events_for_week", middleware.Log(eventHandler.handleGetByWeek))
	mux.HandleFunc("GET /events_for_month", middleware.Log(eventHandler.handleGetByMonth))
}
