package domain

import (
	"time"

	"github.com/sunnygosdk/go-chi-fullcycle-api/pkg/entity"
)

func NewRole(ID entity.ID, Name string) (*role, error) {
	role := &role{
		ID:        ID,
		Name:      Name,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	role, err := role.validate()
	if err != nil {
		return nil, err
	}

	return role, nil
}

func (r *role) validate() (*role, error) {
	if r.ID.String() == "" {
		return nil, ErrRoleIDisRequired
	}

	_, err := entity.ParseID(r.ID.String())
	if err != nil {
		return nil, ErrRoleInvalidID
	}

	if r.Name == "" {
		return nil, ErrRoleNameRequired
	}

	return r, nil
}
