package repository

import "github.com/sunnygosdk/go-chi-fullcycle-api/internal/domain/entity"

// TransactionRepository is the interface for the transaction repository.
type TransactionRepository interface {
	// Create creates a new transaction.
	Create(t *entity.Transaction) error
}
