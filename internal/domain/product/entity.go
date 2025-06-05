package domain

import (
	"time"

	"github.com/sunnygosdk/go-chi-fullcycle-api/pkg/entity"
)

type product struct {
	ID           entity.ID
	Name         string
	Price        float64
	DepartmentID entity.ID
	CreatedAt    time.Time
	UpdatedAt    time.Time
	DeletedAt    *time.Time
}
