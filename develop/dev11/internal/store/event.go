package store

import (
	"calendar/internal/model"
	"errors"
	"fmt"
	"sync"
	"time"

	"github.com/google/uuid"
)

type MemStore struct {
	mu   *sync.RWMutex
	data map[uuid.UUID][]model.Event
}

func NewMemStore() *MemStore {
	return &MemStore{
		mu:   &sync.RWMutex{},
		data: make(map[uuid.UUID][]model.Event),
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

	events, ok := c.data[userID]
	if !ok {
		return nil, fmt.Errorf("user %s not found", userID)
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

func (c *MemStore) Update(userID uuid.UUID, event model.Event) error {
	c.mu.Lock()
	defer c.mu.Unlock()

	events, ok := c.data[userID]
	if !ok {
		return fmt.Errorf("user %s not found", userID)
	}

	found := false
	for _, e := range events {
		if e.ID == event.ID {
			e.Title = event.Title
			e.Description = event.Description
			e.Date = event.Date

			found = true
			break
		}
	}
	if !found {
		return fmt.Errorf("event %d not found", event.ID)
	}
	c.data[userID] = events
	return nil
}

func (c *MemStore) Delete(userID uuid.UUID, eventID uuid.UUID) error {
	c.mu.Lock()
	defer c.mu.Unlock()

	events, ok := c.data[userID]
	if !ok {
		return fmt.Errorf("user %s not found", userID)
	}

	idx := -1
	for i, e := range events {
		if e.ID == eventID {
			idx = i
			break
		}
	}

	if idx == -1 {
		return fmt.Errorf("event %s not found for user %s", eventID, userID)
	}

	c.data[userID] = append(events[:idx], events[idx+1:]...)
	return nil
}

func (c *MemStore) GetEventsByUserID(userID uuid.UUID) ([]model.Event, error) {
	c.mu.RLock()
	defer c.mu.RUnlock()

	events, ok := c.data[userID]
	if !ok {
		return nil, fmt.Errorf("user %s not found", userID)
	}
	return events, nil
}
