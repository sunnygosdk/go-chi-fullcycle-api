package model

import "errors"

var (
	ErrProductIDisRequired    = errors.New("id is required")
	ErrInvalidProductID       = errors.New("invalid product ID")
	ErrProductNameRequired    = errors.New("name is required")
	ErrProductPriceLessOrZero = errors.New("price must be greater than 0")
)
