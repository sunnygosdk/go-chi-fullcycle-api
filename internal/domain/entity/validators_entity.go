package entity

import (
	"github.com/sunnygosdk/go-chi-fullcycle-api/pkg/entity"
)

// ID is an alias for entity.ID
type ID = entity.ID

// InvalidID returns an invalid entity.ID
func InvalidID() ID {
	return entity.InvalidID()
}

// NewID returns a new entity.ID.
func NewID() ID {
	return entity.NewID()
}

// ValidateEntityID validates the provided string ID.
// If the ID is empty or invalid, it returns an invalid ID and an error.
// Otherwise, it returns the parsed entity.ID.
//
// Parameters:
//   - id: The string representation of the entity ID to be validated.
//
// Returns:
//   - entity.ID: The parsed entity ID if valid, otherwise an invalid ID.
//   - error: An error if the ID is empty or invalid.

func ValidateEntityID(id string) (ID, error) {
	if id == "" {
		return InvalidID(), ErrorEntityInvalidID
	}

	idParsed, err := entity.ParseID(id)
	if err != nil {
		return InvalidID(), ErrorEntityInvalidID
	}

	return idParsed, nil
}
