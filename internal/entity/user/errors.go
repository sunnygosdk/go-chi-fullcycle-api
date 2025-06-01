package user

import "errors"

var (
	ErrInvalidID        = errors.New("invalid ID")
	ErrNameRequired     = errors.New("name is required")
	ErrEmailRequired    = errors.New("email is required")
	ErrInvalidEmail     = errors.New("email is invalid")
	ErrPasswordRequired = errors.New("password is required")
	ErrWeakPassword     = errors.New("password must be at least 8 characters and include uppercase, lowercase, number, and special character")
)
