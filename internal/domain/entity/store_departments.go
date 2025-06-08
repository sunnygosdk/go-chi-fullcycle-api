package entity

import (
	"errors"
	"time"

	"github.com/sunnygosdk/go-chi-fullcycle-api/pkg/entity"
)

var (
	ErrorSDInvalidStoreID      = errors.New("store_departments: invalid store ID")
	ErrorSDInvalidDepartmentID = errors.New("store_departments: invalid department ID")
	ErrorSDIsDeleted           = errors.New("store_departments: store is already deleted")
	ErrorSDAtLeastOneField     = errors.New("store_departments: at least one field must be provided")
)

type StoreDepartments struct {
	ID           entity.ID
	StoreID      entity.ID
	DepartmentID entity.ID
	CreatedAt    time.Time
	UpdatedAt    time.Time
	DeletedAt    *time.Time
}

// NewStoreDepartments creates a new store departments instance with the provided ID, store ID, and department ID.
// It initializes the store departments with the given ID, store ID, and department ID,
// and sets the creation and update timestamps to the current time.
//
// The function also validates the store departments before returning it. If validation fails,
// it returns an error.
//
// Parameters:
//   - storeID: Store ID of the store departments.
//   - departmentID: Department ID of the store departments.
//
// Returns:
//   - *StoreDepartments: A pointer to the newly created and validated store departments.
//   - error: An error if the store departments validation fails.
func NewStoreDepartments(storeID string, departmentID string) (*StoreDepartments, error) {
	stID, err := entity.ParseID(storeID)
	if err != nil {
		return nil, ErrorSDInvalidStoreID
	}

	depID, err := entity.ParseID(departmentID)
	if err != nil {
		return nil, ErrorSDInvalidDepartmentID
	}

	storeDepartments := &StoreDepartments{
		ID:           entity.NewID(),
		StoreID:      stID,
		DepartmentID: depID,
		CreatedAt:    time.Now(),
		UpdatedAt:    time.Now(),
		DeletedAt:    nil,
	}

	return storeDepartments, nil
}

// Update updates the store departments with the provided values.
// It validates the store departments before updating it. If validation fails,
// it returns an error.
//
// Parameters:
//   - storeID: Store ID of the store departments.
//   - departmentID: Department ID of the store departments.
//
// Returns:
//   - error: An error if the store departments validation fails.
func (sd *StoreDepartments) Update(storeID *string, departmentID *string) error {
	if storeID == nil && departmentID == nil {
		return ErrorSDAtLeastOneField
	}

	if storeID != nil {
		stID, err := entity.ParseID(*storeID)
		if err != nil {
			return ErrorSDInvalidStoreID
		}
		sd.StoreID = stID
	}

	if departmentID != nil {
		depID, err := entity.ParseID(*departmentID)
		if err != nil {
			return ErrorSDInvalidDepartmentID
		}
		sd.DepartmentID = depID
	}

	sd.UpdatedAt = time.Now()
	return nil
}

// Delete marks the store departments as deleted by setting the deletedAt timestamp to the current time.
// It also validates the store departments before deleting it. If validation fails,
// it returns an error.
func (sd *StoreDepartments) Delete() error {
	if sd.DeletedAt != nil {
		return ErrorSDIsDeleted
	}

	deletedAt := time.Now()
	sd.DeletedAt = &deletedAt
	return nil
}
