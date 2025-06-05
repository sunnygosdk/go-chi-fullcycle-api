package domain

import (
	"time"

	"github.com/sunnygosdk/go-chi-fullcycle-api/pkg/entity"
)

func NewProduct(ID entity.ID, Name string, Price float64, DepartmentID entity.ID) (*product, error) {
	product := &product{
		ID:           ID,
		Name:         Name,
		Price:        Price,
		DepartmentID: DepartmentID,
		CreatedAt:    time.Now(),
		UpdatedAt:    time.Now(),
	}

	product, err := product.validate()
	if err != nil {
		return nil, err
	}

	return product, nil
}

func (p *product) validate() (*product, error) {
	if p.ID.String() == "" {
		return nil, ErrProductIDisRequired
	}

	if p.DepartmentID.String() == "" {
		return nil, ErrProductDepartmentIDisRequired
	}

	_, err := entity.ParseID(p.ID.String())
	if err != nil {
		return nil, ErrProductInvalidID
	}

	_, err = entity.ParseID(p.DepartmentID.String())
	if err != nil {
		return nil, ErrProductInvalidDepartmentID
	}

	if p.Name == "" {
		return nil, ErrProductNameRequired
	}

	if p.Price <= 0 {
		return nil, ErrProductPriceLessOrZero
	}

	return p, nil
}
