package domain

import "github.com/sunnygosdk/go-chi-fullcycle-api/pkg/entity"

func NewDepartment(ID entity.ID, Name string) (*department, error) {
	department := &department{
		ID:   ID,
		Name: Name,
	}

	department, err := department.validate()
	if err != nil {
		return nil, err
	}

	return department, nil
}

func (d *department) validate() (*department, error) {
	if d.ID.String() == "" {
		return nil, ErrDepartmentIDisRequired
	}

	_, err := entity.ParseID(d.ID.String())
	if err != nil {
		return nil, ErrDepartmentInvalidID
	}

	if d.Name == "" {
		return nil, ErrDepartmentNameRequired
	}

	return d, nil
}
