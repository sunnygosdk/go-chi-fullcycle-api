package request

import (
	"time"

	"github.com/sunnygosdk/go-chi-fullcycle-api/internal/product/model"
	"github.com/sunnygosdk/go-chi-fullcycle-api/pkg/entity"
)

type CreateProductRequest struct {
	Name  string  `json:"name"`
	Price float64 `json:"price"`
}

func ProductToCreate(productToCreate CreateProductRequest) (*model.ProductModel, error) {
	product := &model.ProductModel{
		ID:        entity.NewID(),
		Name:      productToCreate.Name,
		Price:     productToCreate.Price,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	err := product.ValidateCreateProduct()
	if err != nil {
		return nil, err
	}

	return product, nil
}
