package store

import (
	"time"

	"github.com/sunnygosdk/go-chi-fullcycle-api/pkg/entity"
)

// store represents a store entity.
type store struct {
	ID        entity.ID
	Name      string
	Address   string
	Contact   string
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time
}
