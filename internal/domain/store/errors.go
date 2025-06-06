package store

import "errors"

// Error messages for store validation
var (
	// ErrStoreIDisRequired is returned when the store ID is required but not provided.
	ErrStoreIDisRequired = errors.New("store id is required")

	// ErrStoreInvalidID is returned when the store ID is invalid.
	ErrStoreInvalidID = errors.New("store id is invalid")

	// ErrStoreNameRequired is returned when the store name is required but not provided.
	ErrStoreNameRequired = errors.New("store name is required")

	// ErrStoreAddressRequired is returned when the store address is required but not provided.
	ErrStoreAddressRequired = errors.New("store address is required")

	// ErrStoreContactRequired is returned when the store contact is required but not provided.
	ErrStoreContactRequired = errors.New("store contact is required")
)
