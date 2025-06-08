package repository

import "github.com/sunnygosdk/go-chi-fullcycle-api/internal/domain/entity"

// ProductRepository is the interface for product repository.
type ProductRepository interface {
	// Create creates a new product.
	Create(product *entity.Product) error
}
