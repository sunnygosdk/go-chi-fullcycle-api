package department

import "github.com/sunnygosdk/go-chi-fullcycle-api/internal/domain/department"

// CreateDepartmentUseCase is the use case for creating a department.
type CreateDepartmentUseCase struct {
	departmentRepository department.DepartmentRepository
}

// CreateDepartmentUseCaseInput is the input for creating a department.
type CreateDepartmentUseCaseInput struct {
	Name string
}

// Execute creates a new department.
//
// Parameters:
//   - input: The input for creating a department.
//
// Returns:
//   - error: An error if the department creation fails.
func (u *CreateDepartmentUseCase) Execute(input *CreateDepartmentUseCaseInput) error {
	newDepartment, err := department.NewDepartment(input.Name)
	if err != nil {
		return err
	}
	return u.departmentRepository.Create(newDepartment)
}
