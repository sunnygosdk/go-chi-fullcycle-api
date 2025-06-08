package store

import (
	"github.com/sunnygosdk/go-chi-fullcycle-api/internal/domain/entity"
	"github.com/sunnygosdk/go-chi-fullcycle-api/internal/domain/repository"
)

// CreateStoreUseCase is the use case for creating a store.
type CreateStoreUseCase struct {
	storeRepository repository.StoreRepository
}

// CreateStoreUseCaseInput is the input for creating a store.
type CreateStoreUseCaseInput struct {
	Name    string
	Address string
}

// Execute creates a new store.
//
// Parameters:
//   - input: The input for creating a store.
//
// Returns:
//   - error: An error if the store creation fails.
func (u *CreateStoreUseCase) Execute(input *CreateStoreUseCaseInput) error {
	newStore, err := entity.NewStore(input.Name, input.Address)
	if err != nil {
		return err
	}
	return u.storeRepository.Create(newStore)
}
