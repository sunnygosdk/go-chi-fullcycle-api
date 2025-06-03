package model

import "errors"

var (
	ErrInvalidUserID    = errors.New("invalid user ID")
	ErrUserNameRequired = errors.New("user name is required")
	ErrEmailRequired    = errors.New("email is required")
	ErrInvalidEmail     = errors.New("email is invalid")
	ErrPasswordRequired = errors.New("password is required")
	ErrWeakPassword     = errors.New("password must be at least 8 characters and include uppercase, lowercase, number, and special character")
)
