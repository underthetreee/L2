package service

import (
	"calendar/internal/model"
	"fmt"
	"time"

	"github.com/google/uuid"
)

type eventStorer interface {
	Add(uuid.UUID, model.Event) error
	GetEventsByUserID(uuid.UUID) ([]model.Event, error)
	Update(uuid.UUID, model.Event) error
	Delete(uuid.UUID, uuid.UUID) error
}

type EventService struct {
	storer eventStorer
}

func NewEventService(storer eventStorer) *EventService {
	return &EventService{
		storer: storer,
	}
}

func (s *EventService) Add(userID uuid.UUID, event model.Event) error {
	if err := s.storer.Add(userID, event); err != nil {
		return err
	}
	return nil
}

func (s *EventService) Update(userID uuid.UUID, event model.Event) error {
	if err := s.storer.Update(userID, event); err != nil {
		return err
	}
	return nil
}

func (s *EventService) Delete(userID uuid.UUID, eventID uuid.UUID) error {
	if err := s.storer.Delete(userID, eventID); err != nil {
		return err
	}
	return nil
}

func (s *EventService) GetByDay(userID uuid.UUID, date time.Time) ([]model.Event, error) {
	events, err := s.storer.GetEventsByUserID(userID)
	if err != nil {
		return nil, err
	}

	var filteredEvents []model.Event
	for _, event := range events {
		if event.Date.Year() == date.Year() && event.Date.Month() == date.Month() && event.Date.Day() == date.Day() {
			filteredEvents = append(filteredEvents, event)
		}
	}

	if len(filteredEvents) == 0 {
		return nil, fmt.Errorf("events not found")
	}

	return filteredEvents, nil
}

func (s *EventService) GetByMonth(userID uuid.UUID, date time.Time) ([]model.Event, error) {
	events, err := s.storer.GetEventsByUserID(userID)
	if err != nil {
		return nil, err
	}

	var filteredEvents []model.Event
	for _, event := range events {
		if event.Date.Year() == date.Year() && event.Date.Month() == date.Month() {
			filteredEvents = append(filteredEvents, event)
		}
	}

	if len(filteredEvents) == 0 {
		return nil, fmt.Errorf("events not found")
	}

	return filteredEvents, nil
}

func (s *EventService) GetByWeek(userID uuid.UUID, date time.Time) ([]model.Event, error) {
	events, err := s.storer.GetEventsByUserID(userID)
	if err != nil {
		return nil, err
	}

	startOfWeek := date.AddDate(0, 0, -int(date.Weekday()))
	endOfWeek := startOfWeek.AddDate(0, 0, 7)

	var filteredEvents []model.Event
	for _, event := range events {
		if event.Date.After(startOfWeek) && event.Date.Before(endOfWeek) {
			filteredEvents = append(filteredEvents, event)
		}
	}

	if len(filteredEvents) == 0 {
		return nil, fmt.Errorf("events not found")
	}

	return filteredEvents, nil
}
