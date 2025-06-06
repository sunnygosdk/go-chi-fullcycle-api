package employee

// EmployeeRepository is an interface for employee repository.
type EmployeeRepository interface {
	// Create creates a new employee.
	Create(employee *employee) error
}
