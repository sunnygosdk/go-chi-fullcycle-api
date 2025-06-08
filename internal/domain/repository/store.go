package repository

import "github.com/sunnygosdk/go-chi-fullcycle-api/internal/domain/entity"

// StoreRepository is the interface for store repository.
type StoreRepository interface {
	// Create creates a new store.
	Create(store *entity.Store) error
}
