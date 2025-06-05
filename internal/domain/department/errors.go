package domain

import "errors"

var (
	ErrDepartmentIDisRequired = errors.New("department ID is required")
	ErrDepartmentInvalidID    = errors.New("invalid department ID")
	ErrDepartmentNameRequired = errors.New("department name is required")
)
