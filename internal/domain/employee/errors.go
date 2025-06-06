package employee

import "errors"

// Error messages for employee validation
var (
	// ErrEmployeeIDisRequired is returned when the employee ID is required but not provided.
	ErrEmployeeIDisRequired = errors.New("employee ID is required")

	// ErrEmployeeInvalidID is returned when the employee ID is invalid.
	ErrEmployeeInvalidID = errors.New("invalid employee ID")

	// ErrEmployeeUsernameRequired is returned when the employee username is required but not provided.
	ErrEmployeeUsernameRequired = errors.New("employee username is required")

	// ErrEmployeeFirstNameRequired is returned when the employee first name is required but not provided.
	ErrEmployeeFirstNameRequired = errors.New("employee first name is required")

	// ErrEmployeeLastNameRequired is returned when the employee last name is required but not provided.
	ErrEmployeeLastNameRequired = errors.New("employee last name is required")

	// ErrEmployeeEmailRequired is returned when the employee email is required but not provided.
	ErrEmployeeEmailRequired = errors.New("employee email is required")

	// ErrEmployeeInvalidEmail is returned when the employee email is invalid.
	ErrEmployeeInvalidEmail = errors.New("invalid employee email")

	// ErrEmployeePasswordRequired is returned when the employee password is required but not provided.
	ErrEmployeePasswordRequired = errors.New("employee password is required")

	// ErrEmployeeWeakPassword is returned when the employee password is weak.
	ErrEmployeeWeakPassword = errors.New("weak password")
)
