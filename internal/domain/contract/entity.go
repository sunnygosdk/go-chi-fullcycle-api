package contract

import (
	"time"

	"github.com/sunnygosdk/go-chi-fullcycle-api/pkg/entity"
)

// contract represents an employment contract that links an employee
// to a role within a department and store for a defined time period.
type contract struct {
	ID           entity.ID
	EmployeeID   entity.ID
	RoleID       entity.ID
	DepartmentID entity.ID
	StoreID      entity.ID
	StartDate    time.Time
	EndDate      time.Time
	CreatedAt    time.Time
	UpdatedAt    time.Time
	DeletedAt    *time.Time
}
