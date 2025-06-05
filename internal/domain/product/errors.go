package domain

import "errors"

var (
	ErrProductIDisRequired           = errors.New("product ID is required")
	ErrProductInvalidID              = errors.New("invalid product ID")
	ErrProductDepartmentIDisRequired = errors.New("product department ID is required")
	ErrProductInvalidDepartmentID    = errors.New("invalid product department ID")
	ErrProductNameRequired           = errors.New("product name is required")
	ErrProductPriceLessOrZero        = errors.New("product price must be greater than zero")
)
