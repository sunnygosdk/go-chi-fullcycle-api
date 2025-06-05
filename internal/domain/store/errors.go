package domain

import "errors"

var (
	ErrStoreIDisRequired    = errors.New("store id is required")
	ErrStoreInvalidID       = errors.New("store id is invalid")
	ErrStoreNameRequired    = errors.New("store name is required")
	ErrStoreAddressRequired = errors.New("store address is required")
	ErrStoreContactRequired = errors.New("store contact is required")
)
