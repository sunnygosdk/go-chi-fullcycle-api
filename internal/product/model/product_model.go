package model

import (
	"time"

	"github.com/sunnygosdk/go-chi-fullcycle-api/internal/product/request"
	"github.com/sunnygosdk/go-chi-fullcycle-api/pkg/entity"
)

type ProductModel struct {
	ID        entity.ID `json:"id"`
	Name      string    `json:"name"`
	Price     float64   `json:"price"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (product *ProductModel) ProductToUpdate(productDTO request.UpdateProductRequest) (*ProductModel, error) {
	if productDTO.Name != nil {
		err := ValidateUpdateName(*productDTO.Name)
		if err != nil {
			return nil, err
		}

		product.Name = *productDTO.Name
	}

	if productDTO.Price != nil {
		err := ValidateUpdatePrice(*productDTO.Price)
		if err != nil {
			return nil, err
		}

		product.Price = *productDTO.Price
	}

	product.UpdatedAt = time.Now()
	return product, nil
}
