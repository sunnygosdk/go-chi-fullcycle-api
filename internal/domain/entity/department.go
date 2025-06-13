package entity

import (
	"time"
)

// Department represents a department within an store.
type Department struct {
	ID          ID
	Name        string
	Description string
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   *time.Time
}

// NewDepartment creates a new department instance with the provided name and description.
// It initializes the department with the given name and description, and sets the creation and update timestamps
// to the current time.
//
// The function also validates the department before returning it. If validation fails,
// it returns an error.
//
// Parameters:
//   - name: Name of the department.
//   - description: Description of the department.
//
// Returns:
//   - *Department: A pointer to the newly created and validated department.
//   - error: An error if the department validation fails.
func NewDepartment(name string, description string) (*Department, error) {
	err := validateDepartmentName(name)
	if err != nil {
		return nil, err
	}

	err = validateDepartmentDescription(description)
	if err != nil {
		return nil, err
	}

	department := &Department{
		ID:          NewID(),
		Name:        name,
		Description: description,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
		DeletedAt:   nil,
	}

	return department, nil
}

// Update updates the department with the provided name and description.
// It updates the department with the given name and description, and sets the update timestamp to the current time.
//
// The function also validates the department before updating it. If validation fails,
// it returns an error.
//
// Parameters:
//   - name: Name of the department.
//   - description: Description of the department.
//
// Returns:
//   - error: An error if the department validation fails.
func (d *Department) Update(name, description *string) error {
	if name == nil && description == nil {
		return ErrorDepartmentAtLeastOneField
	}

	if name != nil {
		err := validateDepartmentName(*name)
		if err != nil {
			return err
		}
		d.Name = *name
	}

	if description != nil {
		err := validateDepartmentDescription(*description)
		if err != nil {
			return err
		}
		d.Description = *description
	}

	d.UpdatedAt = time.Now()
	return nil
}

// Delete marks the department as deleted by setting the deletedAt timestamp to the current time.
// It also validates the department before deleting it. If validation fails,
// it returns an error.
func (d *Department) Delete() error {
	err := validateDepartmentIsDeleted(d)
	if err != nil {
		return err
	}

	deletedAt := time.Now()
	d.DeletedAt = &deletedAt
	return nil
}

// validateDepartmentName validates the department name.
// It checks if the department name is required and if it has a minimum length of 2 characters.
//
// Parameters:
//   - name: Name of the department.
//
// Returns:
//   - error: An error if the department name validation fails.
func validateDepartmentName(name string) error {
	if name == "" {
		return ErrorDepartmentNameRequired
	}

	if len(name) < 2 {
		return ErrorDepartmentNameMinLength
	}

	return nil
}

// validateDepartmentDescription validates the department description.
// It checks if the department description is required and if it has a minimum length of 2 characters.
//
// Parameters:
//   - description: Description of the department.
//
// Returns:
//   - error: An error if the department description validation fails.
func validateDepartmentDescription(description string) error {
	if description == "" {
		return ErrorDepartmentDescriptionRequired
	}

	if len(description) < 2 {
		return ErrorDepartmentDescriptionMinLength
	}

	return nil
}

// validateDepartmentIsDeleted validates the department.
// It checks if the department is already deleted.
//
// Parameters:
//   - department: Department to validate.
//
// Returns:
//   - error: An error if the department is already deleted.
func validateDepartmentIsDeleted(department *Department) error {
	if department.DeletedAt != nil {
		return ErrorDepartmentIsDeleted
	}
	return nil
}
