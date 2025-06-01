package entity

import "github.com/google/uuid"

type ID = uuid.UUID

// NewID returns a new ID.
func NewID() ID {
	return ID(uuid.New())
}

// ParseID returns a new ID from a string.
// If the string is not a valid UUID, it returns an error.
func ParseID(s string) (ID, error) {
	id, err := uuid.Parse(s)
	return ID(id), err
}
