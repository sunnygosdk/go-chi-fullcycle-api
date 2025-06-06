package role

import "errors"

// Error messages for role validation.
var (
	// ErrRoleIDisRequired is returned when the role ID is required but not provided.
	ErrRoleIDisRequired = errors.New("role id is required")

	// ErrRoleInvalidID is returned when the role ID is invalid.
	ErrRoleInvalidID = errors.New("role id is invalid")

	// ErrRoleNameRequired is returned when the role name is required but not provided.
	ErrRoleNameRequired = errors.New("role name is required")

	// ErrRoleTypeRequired is returned when the role type is required but not provided.
	ErrRoleTypeRequired = errors.New("role type is required")

	// ErrRoleInvalidType is returned when the role type is invalid.
	ErrRoleInvalidType = errors.New("role type is invalid")
)
