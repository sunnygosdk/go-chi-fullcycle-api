package product

import (
	"github.com/sunnygosdk/go-chi-fullcycle-api/internal/domain/product"
	"github.com/sunnygosdk/go-chi-fullcycle-api/pkg/entity"
)

// CreateProductUseCase is the use case for creating a product.
type CreateProductUseCase struct {
	productRepository product.ProductRepository
}

// CreateProductUseCaseInput is the input for creating a product.
type CreateProductUseCaseInput struct {
	Name         string
	Price        float64
	DepartmentID entity.ID
}

// Execute creates a new product.
//
// Parameters:
//   - input: The input for creating a product.
//
// Returns:
//   - error: An error if the product creation fails.
func (u *CreateProductUseCase) Execute(input *CreateProductUseCaseInput) error {
	product, err := product.NewProduct(input.Name, input.Price, input.DepartmentID)
	if err != nil {
		return err
	}

	return u.productRepository.Create(product)
}
