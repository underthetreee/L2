package store

import (
	"calendar/internal/model"
	"errors"
	"fmt"
	"sync"
	"time"

	"github.com/google/uuid"
)

type UserID = uuid.UUID

type MemStore struct {
	mu   *sync.RWMutex
	data map[UserID][]model.Event
}

func NewMemStore() *MemStore {
	return &MemStore{
		mu:   &sync.RWMutex{},
		data: make(map[UserID][]model.Event),
	}
}

func (c *MemStore) Add(userID uuid.UUID, event model.Event) error {
	c.mu.Lock()
	defer c.mu.Unlock()

	c.data[userID] = append(c.data[userID], event)
	return nil
}

func (c *MemStore) GetByDate(userID uuid.UUID, date time.Time) ([]model.Event, error) {
	c.mu.RLock()
	defer c.mu.RUnlock()

	events, err := c.findEventsByID(userID)
	if err != nil {
		return nil, err
	}

	var filteredEvents []model.Event
	for _, event := range events {
		if event.Date.Year() == date.Year() &&
			event.Date.Month() == date.Month() &&
			event.Date.Day() == date.Day() {
			filteredEvents = append(filteredEvents, event)
		}
	}

	if len(filteredEvents) == 0 {
		return nil, errors.New("events not found")
	}

	return filteredEvents, nil
}

func (c *MemStore) Update(userID uuid.UUID, eventID uuid.UUID, event model.Event) error {
	c.mu.Lock()
	defer c.mu.Unlock()

	events, err := c.findEventsByID(userID)
	if err != nil {
		return err
	}

	found := false
	for _, e := range events {
		if e.ID == eventID {
			e.Title = event.Title
			e.Description = event.Description
			e.Date = event.Date

			found = true
			break
		}
	}
	if !found {
		return fmt.Errorf("event %d not found", eventID)
	}
	return nil
}

func (c *MemStore) findEventsByID(userID uuid.UUID) ([]model.Event, error) {
	c.mu.RLock()
	defer c.mu.RUnlock()

	events, ok := c.data[userID]
	if !ok {
		return nil, fmt.Errorf("user %d not found", userID)
	}
	return events, nil
}
