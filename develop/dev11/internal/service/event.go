package service

import (
	"calendar/internal/model"
	"fmt"
	"time"

	"github.com/google/uuid"
)

type eventStorer interface {
	Add(uuid.UUID, model.Event) error
	GetByDate(uuid.UUID, time.Time) ([]model.Event, error)
	Update(uuid.UUID, uuid.UUID, model.Event) error
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
	fmt.Println("service:", event)
	return nil
}

func (s *EventService) Update(uuid.UUID, uuid.UUID, model.Event) error {
	return nil
}
