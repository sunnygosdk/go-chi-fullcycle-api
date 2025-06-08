package repository

import "github.com/sunnygosdk/go-chi-fullcycle-api/internal/domain/entity"

// StoreDepartmentsRepository is the interface for store departments repository.
type StoreDepartmentsRepository interface {
	// Create creates a new store departments.
	Create(storeDepartments *entity.StoreDepartments) error
}
