package repository

import "github.com/sunnygosdk/go-chi-fullcycle-api/internal/domain/entity"

// StoreDepartmentMapRepository is the interface for store department map repository.
type StoreDepartmentMapRepository interface {
	// Create creates a new store department map.
	Create(storeDepartmentMap *entity.StoreDepartmentMap) error
}
