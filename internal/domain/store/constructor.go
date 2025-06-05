package domain

import (
	"time"

	"github.com/sunnygosdk/go-chi-fullcycle-api/pkg/entity"
)

func NewStore(ID entity.ID, Name string, Address string, Contact string) (*store, error) {
	store := &store{
		ID:        ID,
		Name:      Name,
		Address:   Address,
		Contact:   Contact,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	store, err := store.validate()
	if err != nil {
		return nil, err
	}

	return store, nil
}

func (s *store) validate() (*store, error) {
	if s.ID.String() == "" {
		return nil, ErrStoreIDisRequired
	}

	_, err := entity.ParseID(s.ID.String())
	if err != nil {
		return nil, ErrStoreInvalidID
	}

	if s.Name == "" {
		return nil, ErrStoreNameRequired
	}

	if s.Address == "" {
		return nil, ErrStoreAddressRequired
	}

	if s.Contact == "" {
		return nil, ErrStoreContactRequired
	}

	return s, nil
}
