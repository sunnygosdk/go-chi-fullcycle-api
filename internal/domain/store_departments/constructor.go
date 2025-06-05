package domain

import (
	"time"

	"github.com/sunnygosdk/go-chi-fullcycle-api/pkg/entity"
)

func NewStoreDepartments(ID entity.ID, StoreID entity.ID, DepartmentID entity.ID) (*storeDepartments, error) {
	storeDepartments := &storeDepartments{
		ID:           ID,
		StoreID:      StoreID,
		DepartmentID: DepartmentID,
		CreatedAt:    time.Now(),
		UpdatedAt:    time.Now(),
	}

	storeDepartments, err := storeDepartments.validate()
	if err != nil {
		return nil, err
	}

	return storeDepartments, nil
}

func (sd *storeDepartments) validate() (*storeDepartments, error) {
	if sd.ID.String() == "" {
		return nil, ErrStoreDepartmentsIDisRequired
	}

	_, err := entity.ParseID(sd.ID.String())
	if err != nil {
		return nil, ErrStoreDepartmentsInvalidID
	}

	if sd.StoreID.String() == "" {
		return nil, ErrStoreDepartmentsStoreIDisRequired
	}

	_, err = entity.ParseID(sd.StoreID.String())
	if err != nil {
		return nil, ErrStoreDepartmentsInvalidStoreID
	}

	if sd.DepartmentID.String() == "" {
		return nil, ErrStoreDepartmentsDepartmentIDisRequired
	}

	_, err = entity.ParseID(sd.DepartmentID.String())
	if err != nil {
		return nil, ErrStoreDepartmentsInvalidDepartmentID
	}

	return sd, nil
}
