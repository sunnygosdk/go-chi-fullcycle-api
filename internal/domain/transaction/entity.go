package transaction

import (
	"time"

	"github.com/sunnygosdk/go-chi-fullcycle-api/pkg/entity"
)

// transaction represents a transaction entity.
type transaction struct {
	ID        entity.ID
	StoreID   entity.ID
	ProductID entity.ID
	Quantity  int
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time
}
