package models

import "github.com/google/uuid"

// Owner model
type Owner struct {
	ID       uuid.UUID `json:"id"`
	FullName string    `json:"fullName"`
}
