package domain

import "errors"

var (
	ErrRoleIDisRequired = errors.New("role id is required")
	ErrRoleInvalidID    = errors.New("role id is invalid")
	ErrRoleNameRequired = errors.New("role name is required")
)
