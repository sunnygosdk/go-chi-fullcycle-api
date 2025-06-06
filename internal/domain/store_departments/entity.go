package store_departments

import (
	"time"

	"github.com/sunnygosdk/go-chi-fullcycle-api/pkg/entity"
)

// storeDepartments represents a store departments entity.
type storeDepartments struct {
	ID           entity.ID
	StoreID      entity.ID
	DepartmentID entity.ID
	CreatedAt    time.Time
	UpdatedAt    time.Time
	DeletedAt    *time.Time
}
