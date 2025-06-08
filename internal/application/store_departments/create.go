package store_departments

import (
	"github.com/sunnygosdk/go-chi-fullcycle-api/internal/domain/entity"
	"github.com/sunnygosdk/go-chi-fullcycle-api/internal/domain/repository"
)

// CreateStoreDepartmentsUseCase is the use case for creating a store departments.
type CreateStoreDepartmentsUseCase struct {
	storeDepartmentRepository repository.StoreDepartmentsRepository
}

// CreateStoreDepartmentsUseCaseInput is the input for creating a store departments.
type CreateStoreDepartmentsUseCaseInput struct {
	StoreID      string
	DepartmentID string
}

// Execute creates a new store departments.
//
// Parameters:
//   - input: The input for creating a store departments.
//
// Returns:
//   - error: An error if the store departments creation fails.
func (u *CreateStoreDepartmentsUseCase) Execute(input *CreateStoreDepartmentsUseCaseInput) error {
	newStoreDepartments, err := entity.NewStoreDepartments(input.StoreID, input.DepartmentID)
	if err != nil {
		return err
	}
	return u.storeDepartmentRepository.Create(newStoreDepartments)
}
