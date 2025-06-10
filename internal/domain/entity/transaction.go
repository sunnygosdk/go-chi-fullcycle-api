package entity

import (
	"errors"
	"time"

	"github.com/sunnygosdk/go-chi-fullcycle-api/pkg/entity"
)

// Error messages for transaction validation
var (
	ErrorTransactionQuantityIsZero         = errors.New("transaction: quantity must not be zero")
	ErrorTransactionInvalidProductID       = errors.New("transaction: invalid product ID")
	ErrorTransactionInvalidStockID         = errors.New("transaction: invalid stock ID")
	ErrorTransactionInvalidTransactionType = errors.New("transaction: invalid transaction type")
	ErrorTransactionAtLeastOneField        = errors.New("transaction: at least one field must be provided")
	ErrorTransactionIsDeleted              = errors.New("transaction: transaction is already deleted")
)

// Transaction represents a transaction within an store.
type Transaction struct {
	ID              entity.ID
	Quantity        int
	TransactionType TransactionType
	ProductID       entity.ID
	StockID         entity.ID
	CreatedAt       time.Time
	UpdatedAt       time.Time
	DeletedAt       *time.Time
}

// TransactionType is a type for transaction type.
type TransactionType string

// TransactionTypes
const (
	TransactionTypeIn  TransactionType = "IN"
	TransactionTypeOut TransactionType = "OUT"
)

// NewTransaction creates a new transaction instance with the provided quantity, transaction type, product ID, and stock ID.
// It initializes the transaction with the given quantity, transaction type, product ID, and stock ID,
// and sets the creation and update timestamps to the current time.
//
// The function also validates the transaction before returning it. If validation fails,
// it returns an error.
//
// Parameters:
//   - quantity: Quantity of the transaction.
//   - transactionType: Transaction type of the transaction.
//   - productID: Product ID of the transaction.
//   - stockID: Stock ID of the transaction.
//
// Returns:
//   - *Transaction: A pointer to the newly created and validated transaction.
//   - error: An error if the transaction validation fails.
func NewTransaction(quantity int, transactionType TransactionType, productID string, stockID string) (*Transaction, error) {
	err := validateTransactionQuantity(quantity)
	if err != nil {
		return nil, err
	}

	prodID, err := entity.ParseID(productID)
	if err != nil {
		return nil, ErrorTransactionInvalidProductID
	}

	stoID, err := entity.ParseID(stockID)
	if err != nil {
		return nil, ErrorTransactionInvalidStockID
	}

	transaction := &Transaction{
		ID:              entity.NewID(),
		Quantity:        quantity,
		TransactionType: transactionType,
		ProductID:       prodID,
		StockID:         stoID,
		CreatedAt:       time.Now(),
		UpdatedAt:       time.Now(),
		DeletedAt:       nil,
	}

	return transaction, nil
}

// Update updates the transaction with the provided values.
// It validates the transaction before updating it. If validation fails,
// it returns an error.
//
// Parameters:
//   - quantity: Quantity of the transaction.
//   - transactionType: Transaction type of the transaction.
//   - productID: Product ID of the transaction.
//   - stockID: Stock ID of the transaction.
//
// Returns:
//   - error: An error if the transaction validation fails.
func (t *Transaction) Update(quantity *int, transactionType *TransactionType, productID *string, stockID *string) error {
	if quantity == nil && transactionType == nil && productID == nil && stockID == nil {
		return ErrorTransactionAtLeastOneField
	}

	if quantity != nil {
		err := validateTransactionQuantity(*quantity)
		if err != nil {
			return err
		}
		t.Quantity = *quantity
	}

	if transactionType != nil {
		t.TransactionType = *transactionType
	}

	if productID != nil {
		prodID, err := entity.ParseID(*productID)
		if err != nil {
			return ErrorTransactionInvalidProductID
		}
		t.ProductID = prodID
	}

	if stockID != nil {
		stoID, err := entity.ParseID(*stockID)
		if err != nil {
			return ErrorTransactionInvalidStockID
		}
		t.StockID = stoID
	}

	t.UpdatedAt = time.Now()
	return nil
}

// Delete marks the transaction as deleted by setting the deletedAt timestamp to the current time.
func (t *Transaction) Delete() error {
	if t.DeletedAt != nil {
		return ErrorTransactionIsDeleted
	}

	deletedAt := time.Now()
	t.DeletedAt = &deletedAt
	return nil
}

// validateTransactionQuantity validates the transaction quantity.
// It returns an error if the transaction quantity is zero.
func validateTransactionQuantity(quantity int) error {
	if quantity == 0 {
		return ErrorTransactionQuantityIsZero
	}
	return nil
}
