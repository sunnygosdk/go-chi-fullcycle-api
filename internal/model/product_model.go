package model

import (
	"time"

	"github.com/sunnygosdk/go-chi-fullcycle-api/pkg/entity"
)

type ProductModel struct {
	ID        entity.ID `json:"id"`
	Name      string    `json:"name"`
	Price     float64   `json:"price"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type CreateProductDTO struct {
	Name  string  `json:"name"`
	Price float64 `json:"price"`
}

type UpdateProductDTO struct {
	Name  *string  `json:"name"`
	Price *float64 `json:"price"`
}

func ProductToCreate(productDTO CreateProductDTO) (*ProductModel, error) {
	product := &ProductModel{
		ID:        entity.NewID(),
		Name:      productDTO.Name,
		Price:     productDTO.Price,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	err := product.ValidateCreateProduct()
	if err != nil {
		return nil, err
	}

	return product, nil
}

func (product *ProductModel) ProductToUpdate(productDTO UpdateProductDTO) (*ProductModel, error) {
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
