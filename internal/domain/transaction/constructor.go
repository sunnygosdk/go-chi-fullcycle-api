package transaction

import (
	"time"

	"github.com/sunnygosdk/go-chi-fullcycle-api/pkg/entity"
)

// NewTransaction creates a new transaction instance with the provided store ID, product ID, and quantity.
// It initializes the transaction with the given store ID, product ID, and quantity,
// and sets the creation and update timestamps to the current time.
//
// The function also validates the transaction before returning it. If validation fails,
// it returns an error.
//
// Parameters:
//   - StoreID: Store ID of the transaction.
//   - ProductID: Product ID of the transaction.
//   - Quantity: Quantity of the transaction.
//
// Returns:
//   - *transaction: A pointer to the newly created and validated transaction.
//   - error: An error if the transaction validation fails.
func NewTransaction(StoreID entity.ID, ProductID entity.ID, Quantity int) (*transaction, error) {
	transaction := &transaction{
		ID:        entity.NewID(),
		StoreID:   StoreID,
		ProductID: ProductID,
		Quantity:  Quantity,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	transaction, err := transaction.validate()
	if err != nil {
		return nil, err
	}

	return transaction, nil
}

// validate validates the transaction instance.
// It checks if the transaction ID is required and valid,
// and if the transaction product ID is required.
//
// Parameters:
//   - t: The transaction instance to validate.
//
// Returns:
//   - *transaction: A pointer to the validated transaction instance.
//   - error: An error if the transaction validation fails.
func (t *transaction) validate() (*transaction, error) {
	if t.ID.String() == "" {
		return nil, ErrTransactionIDisRequired
	}

	_, err := entity.ParseID(t.ID.String())
	if err != nil {
		return nil, ErrTransactionInvalidID
	}

	if t.ProductID.String() == "" {
		return nil, ErrTransactionProductIDisRequired
	}

	_, err = entity.ParseID(t.ProductID.String())
	if err != nil {
		return nil, ErrTransactionInvalidProductID
	}

	if t.Quantity == 0 {
		return nil, ErrTransactionQuantityEqualZero
	}

	return t, nil
}
