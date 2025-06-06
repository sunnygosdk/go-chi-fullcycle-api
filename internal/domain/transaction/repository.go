package transaction

// TransactionRepository is the interface for the transaction repository.
type TransactionRepository interface {
	// Create creates a new transaction.
	Create(t *transaction) error
}
