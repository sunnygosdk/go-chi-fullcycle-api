package product

import "errors"

// Error messages for product validation
var (
	// ErrProductIDisRequired is returned when the product ID is required but not provided.
	ErrProductIDisRequired = errors.New("product ID is required")

	// ErrProductInvalidID is returned when the product ID is invalid.
	ErrProductInvalidID = errors.New("invalid product ID")

	// ErrProductDepartmentIDisRequired is returned when the product department ID is required but not provided.
	ErrProductDepartmentIDisRequired = errors.New("product department ID is required")

	// ErrProductInvalidDepartmentID is returned when the product department ID is invalid.
	ErrProductInvalidDepartmentID = errors.New("invalid product department ID")

	// ErrProductNameRequired is returned when the product name is required but not provided.
	ErrProductNameRequired = errors.New("product name is required")

	// ErrProductPriceLessOrZero is returned when the product price is less than or equal to zero.
	ErrProductPriceLessOrZero = errors.New("product price must be greater than zero")
)
