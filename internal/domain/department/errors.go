package department

import "errors"

// Error messages for department validation
var (
	// ErrDepartmentIDisRequired is returned when the department ID is required but not provided.
	ErrDepartmentIDisRequired = errors.New("department ID is required")

	// ErrDepartmentInvalidID is returned when the department ID is invalid.
	ErrDepartmentInvalidID = errors.New("invalid department ID")

	// ErrDepartmentNameRequired is returned when the department name is required but not provided.
	ErrDepartmentNameRequired = errors.New("department name is required")
)
