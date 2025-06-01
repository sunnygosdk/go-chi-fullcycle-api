package product

import "github.com/sunnygosdk/go-chi-fullcycle-api/pkg/entity"

func (product *Model) ValidateNewProduct() error {
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
