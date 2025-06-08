package repository

import "github.com/sunnygosdk/go-chi-fullcycle-api/internal/domain/entity"

// DepartmentRepository is an interface for department repository.
type DepartmentRepository interface {
	// Create creates a new department.
	Create(department *entity.Department) error
}
