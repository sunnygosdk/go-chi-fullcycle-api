package user

import "errors"

var (
	ErrInvalidID        = errors.New("invalid ID")
	ErrDifferentID      = errors.New("different ID")
	ErrNameRequired     = errors.New("name is required")
	ErrEmailRequired    = errors.New("email is required")
	ErrInvalidEmail     = errors.New("email is invalid")
	ErrPasswordRequired = errors.New("password is required")
	ErrSamePassword     = errors.New("password is the same")
	ErrSameEmail        = errors.New("email is the same")
	ErrSameName         = errors.New("name is the same")
	ErrWeakPassword     = errors.New("password must be at least 8 characters and include uppercase, lowercase, number, and special character")
)
