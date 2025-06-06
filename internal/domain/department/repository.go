package department

// DepartmentRepository is an interface for department repository.
type DepartmentRepository interface {
	// Create creates a new department.
	Create(department *department) error
}
