package contract

import (
	"time"

	"github.com/sunnygosdk/go-chi-fullcycle-api/pkg/entity"
)

// NewContract creates and returns a new contract instance with the provided parameters.
// It initializes the contract with the given IDs (Contract ID, Employee ID, Role ID,
// Department ID, Store ID), start and end dates, and sets the creation and update timestamps
// to the current time.
//
// The function also validates the contract before returning it. If validation fails,
// it returns an error.
//
// Parameters:
//   - EmployeeID: Unique identifier for the employee associated with the contract.
//   - RoleID: Unique identifier for the role assigned to the employee.
//   - DepartmentID: Unique identifier for the department the employee belongs to.
//   - StoreID: Unique identifier for the store where the employee will work.
//   - StartDate: The date the contract begins.
//   - EndDate: The date the contract ends.
//
// Returns:
//   - *contract: A pointer to the newly created and validated contract.
//   - error: An error if the contract validation fails.
func NewContract(EmployeeID entity.ID, RoleID entity.ID, DepartmentID entity.ID, StoreID entity.ID, StartDate time.Time, EndDate time.Time) (*contract, error) {
	contract := &contract{
		ID:           entity.NewID(),
		EmployeeID:   EmployeeID,
		RoleID:       RoleID,
		DepartmentID: DepartmentID,
		StoreID:      StoreID,
		StartDate:    StartDate,
		EndDate:      EndDate,
		CreatedAt:    time.Now(),
		UpdatedAt:    time.Now(),
	}

	contract, err := contract.validate()
	if err != nil {
		return nil, err
	}

	return contract, nil
}

// validate validates the contract instance.
// It checks if all required fields are set and if the IDs are valid.
// If any validation fails, it returns an error.
//
// Returns:
//   - *contract: A pointer to the validated contract.
//   - error: An error if the validation fails.
func (c *contract) validate() (*contract, error) {
	if c.ID.String() == "" {
		return nil, ErrContractIDisRequired
	}

	_, err := entity.ParseID(c.ID.String())
	if err != nil {
		return nil, ErrContractInvalidID
	}

	if c.EmployeeID.String() == "" {
		return nil, ErrContractEmployeeIDisRequired
	}

	_, err = entity.ParseID(c.EmployeeID.String())
	if err != nil {
		return nil, ErrContractInvalidEmployeeID
	}

	if c.RoleID.String() == "" {
		return nil, ErrContractRoleIDisRequired
	}

	_, err = entity.ParseID(c.RoleID.String())
	if err != nil {
		return nil, ErrContractInvalidRoleID
	}

	if c.DepartmentID.String() == "" {
		return nil, ErrContractDepartmentIDisRequired
	}

	_, err = entity.ParseID(c.DepartmentID.String())
	if err != nil {
		return nil, ErrContractInvalidDepartmentID
	}

	if c.StoreID.String() == "" {
		return nil, ErrContractStoreIDisRequired
	}

	_, err = entity.ParseID(c.StoreID.String())
	if err != nil {
		return nil, ErrContractInvalidStoreID
	}

	if c.StartDate.IsZero() {
		return nil, ErrContractStartDateRequired
	}

	if c.EndDate.IsZero() {
		return nil, ErrContractEndDateRequired
	}

	return c, nil
}
