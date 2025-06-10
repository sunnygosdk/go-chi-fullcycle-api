package entity

import (
	"time"

	"github.com/sunnygosdk/go-chi-fullcycle-api/pkg/entity"
)

// Store represents a store.
type Store struct {
	ID        entity.ID
	Name      string
	Address   string
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time
}

// NewStore creates a new store instance with the provided ID, name, and address.
// It initializes the store with the given ID, name, and address,
// and sets the creation and update timestamps to the current time.
//
// Parameters:
//   - id: ID of the store.
//   - name: Name of the store.
//   - address: Address of the store.
//
// Returns:
//   - *Store: A pointer to the newly created store.
//   - error: An error if the store validation fails.
func NewStore(name string, address string) (*Store, error) {

	if err := validateStoreName(name); err != nil {
		return nil, err
	}

	if err := validateStoreAddress(address); err != nil {
		return nil, err
	}

	store := &Store{
		ID:        entity.NewID(),
		Name:      name,
		Address:   address,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		DeletedAt: nil,
	}

	return store, nil
}

// Update updates the store with the provided values.
// It validates the store before updating it. If validation fails,
// it returns an error.
//
// Parameters:
//   - name: Name of the store.
//   - address: Address of the store.
//
// Returns:
//   - error: An error if the store validation fails.
func (s *Store) Update(name *string, address *string) error {
	if name == nil && address == nil {
		return ErrorStoreAtLeastOneField
	}

	if name != nil {
		if err := validateStoreName(*name); err != nil {
			return err
		}
		s.Name = *name
	}

	if address != nil {
		if err := validateStoreAddress(*address); err != nil {
			return err
		}
		s.Address = *address
	}

	s.UpdatedAt = time.Now()
	return nil
}

// Delete marks the store as deleted by setting the deletedAt timestamp to the current time.
// It also validates the store before deleting it. If validation fails,
// it returns an error.
func (s *Store) Delete() error {
	if s.DeletedAt != nil {
		return ErrorStoreIsDeleted
	}

	deletedAt := time.Now()
	s.DeletedAt = &deletedAt
	return nil
}

// validateStoreName validates the name of the store.
// It returns an error if the name is empty.
//
// Parameters:
//   - name: Name of the store.
//
// Returns:
//   - error: An error if the store name validation fails.
func validateStoreName(name string) error {
	if name == "" {
		return ErrorStoreInvalidName
	}

	if len(name) < 3 {
		return ErrorStoreMinLengthName
	}

	return nil
}

// validateStoreAddress validates the address of the store.
// It returns an error if the address is empty.
//
// Parameters:
//   - address: Address of the store.
//
// Returns:
//   - error: An error if the store address validation fails.
func validateStoreAddress(address string) error {
	if address == "" {
		return ErrorStoreInvalidAddress
	}

	if len(address) < 3 {
		return ErrorStoreMinLengthAddress
	}

	return nil
}
