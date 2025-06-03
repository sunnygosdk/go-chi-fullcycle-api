package product

import "github.com/sunnygosdk/go-chi-fullcycle-api/pkg/entity"

func (product *Model) ValidateCreateProduct() error {
	if product.ID.String() == "" {
		return ErrIDisRequired
	}

	_, err := entity.ParseID(product.ID.String())
	if err != nil {
		return ErrInvalidID
	}

	if product.Name == "" {
		return ErrNameRequired
	}

	if product.Price <= 0 {
		return ErrPriceLessOrZero
	}

	return nil
}

func ValidateUpdateName(newName string) error {
	if newName == "" {
		return ErrNameRequired
	}

	return nil
}

func ValidateUpdatePrice(price float64) error {
	if price <= 0 {
		return ErrPriceLessOrZero
	}

	return nil
}
