package product

import "errors"

var (
	ErrIDisRequired    = errors.New("id is required")
	ErrInvalidID       = errors.New("invalid id")
	ErrNameRequired    = errors.New("name is required")
	ErrPriceLessOrZero = errors.New("price must be greater than 0")
)
