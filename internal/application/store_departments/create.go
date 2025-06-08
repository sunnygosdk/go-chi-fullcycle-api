package store_departments

import (
	"github.com/sunnygosdk/go-chi-fullcycle-api/internal/domain/store_departments"
	"github.com/sunnygosdk/go-chi-fullcycle-api/pkg/entity"
)

// CreateStoreDepartmentsUseCase is the use case for creating a store departments.
type CreateStoreDepartmentsUseCase struct {
	storeDepartmentRepository store_departments.StoreDepartmentsRepository
}

// CreateStoreDepartmentsUseCaseInput is the input for creating a store departments.
type CreateStoreDepartmentsUseCaseInput struct {
	StoreID      entity.ID
	DepartmentID entity.ID
}

// Execute creates a new store departments.
//
// Parameters:
//   - input: The input for creating a store departments.
//
// Returns:
//   - error: An error if the store departments creation fails.
func (u *CreateStoreDepartmentsUseCase) Execute(input *CreateStoreDepartmentsUseCaseInput) error {
	newStoreDepartments, err := store_departments.NewStoreDepartments(input.StoreID, input.DepartmentID)
	if err != nil {
		return err
	}
	return u.storeDepartmentRepository.Create(newStoreDepartments)
}
