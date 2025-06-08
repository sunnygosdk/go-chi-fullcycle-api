package transaction

import (
	"github.com/sunnygosdk/go-chi-fullcycle-api/internal/domain/entity"
	"github.com/sunnygosdk/go-chi-fullcycle-api/internal/domain/repository"
)

// CreateTransactionUseCase is the use case for creating a transaction.
type CreateTransactionUseCase struct {
	transactionRepository repository.TransactionRepository
}

// CreateTransactionUseCaseInput is the input for creating a transaction.
type CreateTransactionUseCaseInput struct {
	Quantity  int
	Type      entity.TransactionType
	StoreID   string
	ProductID string
}

// Execute creates a new transaction.
//
// Parameters:
//   - input: The input for creating a transaction.
//
// Returns:
//   - error: An error if the transaction creation fails.
func (u *CreateTransactionUseCase) Execute(input *CreateTransactionUseCaseInput) error {
	newTransaction, err := entity.NewTransaction(input.Quantity, input.Type, input.ProductID, input.StoreID)
	if err != nil {
		return err
	}
	return u.transactionRepository.Create(newTransaction)
}
