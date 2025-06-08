package transaction

import (
	"github.com/sunnygosdk/go-chi-fullcycle-api/internal/domain/transaction"
	"github.com/sunnygosdk/go-chi-fullcycle-api/pkg/entity"
)

// CreateTransactionUseCase is the use case for creating a transaction.
type CreateTransactionUseCase struct {
	transactionRepository transaction.TransactionRepository
}

// CreateTransactionUseCaseInput is the input for creating a transaction.
type CreateTransactionUseCaseInput struct {
	StoreID   entity.ID
	ProductID entity.ID
	Quantity  int
}

// Execute creates a new transaction.
//
// Parameters:
//   - input: The input for creating a transaction.
//
// Returns:
//   - error: An error if the transaction creation fails.
func (u *CreateTransactionUseCase) Execute(input *CreateTransactionUseCaseInput) error {
	newTransaction, err := transaction.NewTransaction(input.StoreID, input.ProductID, input.Quantity)
	if err != nil {
		return err
	}
	return u.transactionRepository.Create(newTransaction)
}
