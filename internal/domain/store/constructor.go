package store

import (
	"time"

	"github.com/sunnygosdk/go-chi-fullcycle-api/pkg/entity"
)

// NewStore creates a new store instance with the provided ID, name, address, and contact.
// It initializes the store with the given ID, name, address, and contact,
// and sets the creation and update timestamps to the current time.
//
// The function also validates the store before returning it. If validation fails,
// it returns an error.
//
// Parameters:
//   - Name: Name of the store.
//   - Address: Address of the store.
//   - Contact: Contact information for the store.
//
// Returns:
//   - *store: A pointer to the newly created and validated store.
//   - error: An error if the store validation fails.
func NewStore(Name string, Address string, Contact string) (*store, error) {
	store := &store{
		ID:        entity.NewID(),
		Name:      Name,
		Address:   Address,
		Contact:   Contact,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	store, err := store.validate()
	if err != nil {
		return nil, err
	}

	return store, nil
}

// validate validates the store instance.
// It checks if the store ID is required and valid,
// and if the store name, address, and contact are required.
//
// Parameters:
//   - s: The store instance to validate.
//
// Returns:
//   - *store: A pointer to the validated store instance.
//   - error: An error if the store validation fails.
func (s *store) validate() (*store, error) {
	if s.ID.String() == "" {
		return nil, ErrStoreIDisRequired
	}

	_, err := entity.ParseID(s.ID.String())
	if err != nil {
		return nil, ErrStoreInvalidID
	}

	if s.Name == "" {
		return nil, ErrStoreNameRequired
	}

	if s.Address == "" {
		return nil, ErrStoreAddressRequired
	}

	if s.Contact == "" {
		return nil, ErrStoreContactRequired
	}

	return s, nil
}
