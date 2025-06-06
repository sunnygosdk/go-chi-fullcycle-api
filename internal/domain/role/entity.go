package role

import (
	"time"

	"github.com/sunnygosdk/go-chi-fullcycle-api/pkg/entity"
)

// role represents a role entity.
type role struct {
	ID        entity.ID
	Name      string
	Type      roleType
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time
}
