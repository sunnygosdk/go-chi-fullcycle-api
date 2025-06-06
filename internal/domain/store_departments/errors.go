package store_departments

import "errors"

// Error messages for store departments validation
var (
	// ErrStoreDepartmentsIDisRequired is returned when the store departments ID is required but not provided.
	ErrStoreDepartmentsIDisRequired = errors.New("store departments id is required")

	// ErrStoreDepartmentsInvalidID is returned when the store departments ID is invalid.
	ErrStoreDepartmentsInvalidID = errors.New("store departments id is invalid")

	// ErrStoreDepartmentsStoreIDisRequired is returned when the store departments store ID is required but not provided.
	ErrStoreDepartmentsStoreIDisRequired = errors.New("store departments store id is required")

	// ErrStoreDepartmentsInvalidStoreID is returned when the store departments store ID is invalid.
	ErrStoreDepartmentsInvalidStoreID = errors.New("store departments store id is invalid")

	// ErrStoreDepartmentsDepartmentIDisRequired is returned when the store departments department ID is required but not provided.
	ErrStoreDepartmentsDepartmentIDisRequired = errors.New("store departments department id is required")

	// ErrStoreDepartmentsInvalidDepartmentID is returned when the store departments department ID is invalid.
	ErrStoreDepartmentsInvalidDepartmentID = errors.New("store departments department id is invalid")
)
