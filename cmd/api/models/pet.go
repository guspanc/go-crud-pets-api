package models

import "github.com/google/uuid"

// Pet model
type Pet struct {
	ID    uuid.UUID `json:"id"`
	Name  string    `json:"name"`
	Breed string    `json:"breed"`
	Age   int       `json:"age"`
	Owner *Owner    `json:"owner"`
}
