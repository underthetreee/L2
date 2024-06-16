package model

import (
	"fmt"
	"time"

	"github.com/google/uuid"
)

type Event struct {
	ID          uuid.UUID
	Title       string
	Description string
	Date        time.Time
}

func (e *Event) Validate() error {
	if e.Title == "" {
		return fmt.Errorf("title cannot be empty")
	}

	if e.Description == "" {
		return fmt.Errorf("description cannot be empty")
	}
	return nil
}
