package store_department_map

import (
	"github.com/sunnygosdk/go-chi-fullcycle-api/internal/domain/entity"
	"github.com/sunnygosdk/go-chi-fullcycle-api/internal/domain/repository"
)

// CreateStoreDepartmentMapUseCase is the use case for creating a store department map.
type CreateStoreDepartmentMapUseCase struct {
	storeDepartmentMapRepository repository.StoreDepartmentMapRepository
}

// CreateStoreDepartmentMapUseCaseInput is the input for creating a store department map.
type CreateStoreDepartmentMapUseCaseInput struct {
	StoreID      string
	DepartmentID string
}

// Execute creates a new store department map.
//
// Parameters:
//   - input: The input for creating a store department map.
//
// Returns:
//   - error: An error if the store department map creation fails.
func (u *CreateStoreDepartmentMapUseCase) Execute(input *CreateStoreDepartmentMapUseCaseInput) error {
	newStoreDepartmentMap, err := entity.NewStoreDepartmentMap(input.StoreID, input.DepartmentID)
	if err != nil {
		return err
	}
	return u.storeDepartmentMapRepository.Create(newStoreDepartmentMap)
}
