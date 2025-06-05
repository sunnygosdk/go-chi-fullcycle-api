package domain

import (
	"time"

	"github.com/sunnygosdk/go-chi-fullcycle-api/pkg/entity"
)

type employee struct {
	ID        entity.ID
	Username  string
	FirstName string
	LastName  string
	Email     string
	Password  string
	RoleID    entity.ID
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time
}
