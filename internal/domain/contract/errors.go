package contract

import "errors"

// Error messages for contract validation
var (
	// ErrContractIDisRequired is returned when the contract ID is required but not provided.
	ErrContractIDisRequired = errors.New("contract ID is required")

	// ErrContractInvalidID is returned when the contract ID is invalid.
	ErrContractInvalidID = errors.New("contract ID is invalid")

	// ErrContractEmployeeIDisRequired is returned when the contract employee ID is required but not provided.
	ErrContractEmployeeIDisRequired = errors.New("contract employee ID is required")

	// ErrContractInvalidEmployeeID is returned when the contract employee ID is invalid.
	ErrContractInvalidEmployeeID = errors.New("contract employee ID is invalid")

	// ErrContractRoleIDisRequired is returned when the contract role ID is required but not provided.
	ErrContractRoleIDisRequired = errors.New("contract role ID is required")

	// ErrContractInvalidRoleID is returned when the contract role ID is invalid.
	ErrContractInvalidRoleID = errors.New("contract role ID is invalid")

	// ErrContractDepartmentIDisRequired is returned when the contract department ID is required but not provided.
	ErrContractDepartmentIDisRequired = errors.New("contract department ID is required")

	// ErrContractInvalidDepartmentID is returned when the contract department ID is invalid.
	ErrContractInvalidDepartmentID = errors.New("contract department ID is invalid")

	// ErrContractStoreIDisRequired is returned when the contract store ID is required but not provided.
	ErrContractStoreIDisRequired = errors.New("contract store ID is required")

	// ErrContractInvalidStoreID is returned when the contract store ID is invalid.
	ErrContractInvalidStoreID = errors.New("contract store ID is invalid")

	// ErrContractStartDateRequired is returned when the contract start date is required but not provided.
	ErrContractStartDateRequired = errors.New("contract start date is required")

	// ErrContractEndDateRequired is returned when the contract end date is required but not provided.
	ErrContractEndDateRequired = errors.New("contract end date is required")
)
