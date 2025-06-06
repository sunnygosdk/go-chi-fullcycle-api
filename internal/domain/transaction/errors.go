package transaction

import "errors"

// Error messages for transaction validation.
var (
	// ErrTransactionIDisRequired is returned when the transaction ID is required.
	ErrTransactionIDisRequired = errors.New("transaction id is required")

	// ErrTransactionInvalidID is returned when the transaction ID is invalid.
	ErrTransactionInvalidID = errors.New("transaction id is invalid")

	// ErrTransactionStoreIDisRequired is returned when the transaction store ID is required.
	ErrTransactionStoreIDisRequired = errors.New("transaction store id is required")

	// ErrTransactionInvalidStoreID is returned when the transaction store ID is invalid.
	ErrTransactionInvalidStoreID = errors.New("transaction store id is invalid")

	// ErrTransactionProductIDisRequired is returned when the transaction product ID is required.
	ErrTransactionProductIDisRequired = errors.New("transaction product id is required")

	// ErrTransactionInvalidProductID is returned when the transaction product ID is invalid.
	ErrTransactionInvalidProductID = errors.New("transaction product id is invalid")

	// ErrTransactionQuantityEqualZero is returned when the transaction quantity is equal to zero.
	ErrTransactionQuantityEqualZero = errors.New("transaction quantity no can be equal zero")
)
