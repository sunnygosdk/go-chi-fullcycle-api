package store_departments

import (
	"time"

	"github.com/sunnygosdk/go-chi-fullcycle-api/pkg/entity"
)

// NewStoreDepartments creates a new store departments instance with the provided ID, store ID, and department ID.
// It initializes the store departments with the given ID, store ID, and department ID,
// and sets the creation and update timestamps to the current time.
//
// The function also validates the store departments before returning it. If validation fails,
// it returns an error.
//
// Parameters:
//   - StoreID: Store ID of the store departments.
//   - DepartmentID: Department ID of the store departments.
//
// Returns:
//   - *storeDepartments: A pointer to the newly created and validated store departments.
//   - error: An error if the store departments validation fails.
func NewStoreDepartments(StoreID entity.ID, DepartmentID entity.ID) (*storeDepartments, error) {
	storeDepartments := &storeDepartments{
		ID:           entity.NewID(),
		StoreID:      StoreID,
		DepartmentID: DepartmentID,
		CreatedAt:    time.Now(),
		UpdatedAt:    time.Now(),
	}

	storeDepartments, err := storeDepartments.validate()
	if err != nil {
		return nil, err
	}

	return storeDepartments, nil
}

// validate validates the store departments instance.
// It checks if the store departments ID is required and valid,
// and if the store departments store ID and department ID are required.
//
// Parameters:
//   - sd: The store departments instance to validate.
//
// Returns:
//   - *storeDepartments: A pointer to the validated store departments instance.
//   - error: An error if the store departments validation fails.
func (sd *storeDepartments) validate() (*storeDepartments, error) {
	if sd.ID.String() == "" {
		return nil, ErrStoreDepartmentsIDisRequired
	}

	_, err := entity.ParseID(sd.ID.String())
	if err != nil {
		return nil, ErrStoreDepartmentsInvalidID
	}

	if sd.StoreID.String() == "" {
		return nil, ErrStoreDepartmentsStoreIDisRequired
	}

	_, err = entity.ParseID(sd.StoreID.String())
	if err != nil {
		return nil, ErrStoreDepartmentsInvalidStoreID
	}

	if sd.DepartmentID.String() == "" {
		return nil, ErrStoreDepartmentsDepartmentIDisRequired
	}

	_, err = entity.ParseID(sd.DepartmentID.String())
	if err != nil {
		return nil, ErrStoreDepartmentsInvalidDepartmentID
	}

	return sd, nil
}
