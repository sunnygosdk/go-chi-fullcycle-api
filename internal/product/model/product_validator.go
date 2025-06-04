package model

import "github.com/sunnygosdk/go-chi-fullcycle-api/pkg/entity"

func (product *ProductModel) ValidateCreateProduct() error {
	if product.ID.String() == "" {
		return ErrProductIDisRequired
	}

	_, err := entity.ParseID(product.ID.String())
	if err != nil {
		return ErrInvalidProductID
	}

	if product.Name == "" {
		return ErrProductNameRequired
	}

	if product.Price <= 0 {
		return ErrProductPriceLessOrZero
	}

	return nil
}

func ValidateUpdateName(newName string) error {
	if newName == "" {
		return ErrProductNameRequired
	}

	return nil
}

func ValidateUpdatePrice(price float64) error {
	if price <= 0 {
		return ErrProductPriceLessOrZero
	}

	return nil
}
