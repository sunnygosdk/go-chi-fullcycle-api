package store

// StoreRepository is the interface for store repository.
type StoreRepository interface {
	// Create creates a new store.
	Create(store *store) error
}
