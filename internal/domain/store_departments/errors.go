package domain

import "errors"

var (
	ErrStoreDepartmentsIDisRequired           = errors.New("store departments id is required")
	ErrStoreDepartmentsInvalidID              = errors.New("store departments id is invalid")
	ErrStoreDepartmentsStoreIDisRequired      = errors.New("store departments store id is required")
	ErrStoreDepartmentsInvalidStoreID         = errors.New("store departments store id is invalid")
	ErrStoreDepartmentsDepartmentIDisRequired = errors.New("store departments department id is required")
	ErrStoreDepartmentsInvalidDepartmentID    = errors.New("store departments department id is invalid")
)
