package entity

import (
	"time"

	"github.com/sunnygosdk/go-chi-fullcycle-api/pkg/entity"
)

// StoreDepartmentMap represents a store department map.
type StoreDepartmentMap struct {
	ID           entity.ID
	StoreID      entity.ID
	DepartmentID entity.ID
	CreatedAt    time.Time
	UpdatedAt    time.Time
	DeletedAt    *time.Time
}

// NewStoreDepartmentMap creates a new store department map instance with the provided ID, store ID, and department ID.
// It initializes the store department map with the given ID, store ID, and department ID,
// and sets the creation and update timestamps to the current time.
//
// The function also validates the store department map before returning it. If validation fails,
// it returns an error.
//
// Parameters:
//   - storeID: Store ID of the store department map.
//   - departmentID: Department ID of the store department map.
//
// Returns:
//   - *StoreDepartmentMap: A pointer to the newly created and validated store department map.
//   - error: An error if the store department map validation fails.
func NewStoreDepartmentMap(storeID string, departmentID string) (*StoreDepartmentMap, error) {
	stID, err := entity.ParseID(storeID)
	if err != nil {
		return nil, ErrorSDMInvalidStoreID
	}

	depID, err := entity.ParseID(departmentID)
	if err != nil {
		return nil, ErrorSDMInvalidDepartmentID
	}

	storeDepartmentMap := &StoreDepartmentMap{
		ID:           entity.NewID(),
		StoreID:      stID,
		DepartmentID: depID,
		CreatedAt:    time.Now(),
		UpdatedAt:    time.Now(),
		DeletedAt:    nil,
	}

	return storeDepartmentMap, nil
}

// Update updates the store department map with the provided values.
// It validates the store department map before updating it. If validation fails,
// it returns an error.
//
// Parameters:
//   - storeID: Store ID of the store department map.
//   - departmentID: Department ID of the store department map.
//
// Returns:
//   - error: An error if the store department map validation fails.
func (sd *StoreDepartmentMap) Update(storeID *string, departmentID *string) error {
	if storeID == nil && departmentID == nil {
		return ErrorSDMAtLeastOneField
	}

	if storeID != nil {
		stID, err := entity.ParseID(*storeID)
		if err != nil {
			return ErrorSDMInvalidStoreID
		}
		sd.StoreID = stID
	}

	if departmentID != nil {
		depID, err := entity.ParseID(*departmentID)
		if err != nil {
			return ErrorSDMInvalidDepartmentID
		}
		sd.DepartmentID = depID
	}

	sd.UpdatedAt = time.Now()
	return nil
}

// Delete marks the store department map as deleted by setting the deletedAt timestamp to the current time.
// It also validates the store department map before deleting it. If validation fails,
// it returns an error.
func (sd *StoreDepartmentMap) Delete() error {
	if sd.DeletedAt != nil {
		return ErrorSDMIsDeleted
	}

	deletedAt := time.Now()
	sd.DeletedAt = &deletedAt
	return nil
}
