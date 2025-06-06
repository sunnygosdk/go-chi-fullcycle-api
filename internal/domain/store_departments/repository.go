package store_departments

// StoreDepartmentsRepository is the interface for store departments repository.
type StoreDepartmentsRepository interface {
	// Create creates a new store departments.
	Create(storeDepartments *storeDepartments) error
}
