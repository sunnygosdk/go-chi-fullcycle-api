package department

import (
	"time"

	"github.com/sunnygosdk/go-chi-fullcycle-api/pkg/entity"
)

// NewDepartment creates a new department instance with the provided ID and name.
// It initializes the department with the given ID and name, and sets the creation and update timestamps
// to the current time.
//
// The function also validates the department before returning it. If validation fails,
// it returns an error.
//
// Parameters:
//   - Name: Name of the department.
//
// Returns:
//   - *department: A pointer to the newly created and validated department.
//   - error: An error if the department validation fails.
func NewDepartment(Name string) (*department, error) {
	department := &department{
		ID:        entity.NewID(),
		Name:      Name,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	department, err := department.validate()
	if err != nil {
		return nil, err
	}

	return department, nil
}

// validate validates the department instance.
// It checks if the department ID is required and valid,
// and if the department name is required.
//
// Parameters:
//   - d: The department instance to validate.
//
// Returns:
//   - *department: A pointer to the validated department instance.
//   - error: An error if the department validation fails.
func (d *department) validate() (*department, error) {
	if d.ID.String() == "" {
		return nil, ErrDepartmentIDisRequired
	}

	_, err := entity.ParseID(d.ID.String())
	if err != nil {
		return nil, ErrDepartmentInvalidID
	}

	if d.Name == "" {
		return nil, ErrDepartmentNameRequired
	}

	return d, nil
}
