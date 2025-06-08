package product

import (
	"github.com/sunnygosdk/go-chi-fullcycle-api/internal/domain/entity"
	"github.com/sunnygosdk/go-chi-fullcycle-api/internal/domain/repository"
)

// CreateProductUseCase is the use case for creating a product.
type CreateProductUseCase struct {
	productRepository repository.ProductRepository
}

// CreateProductUseCaseInput is the input for creating a product.
type CreateProductUseCaseInput struct {
	Name         string
	Description  string
	Price        float64
	DepartmentID string
}

// Execute creates a new product.
//
// Parameters:
//   - input: The input for creating a product.
//
// Returns:
//   - error: An error if the product creation fails.
func (u *CreateProductUseCase) Execute(input *CreateProductUseCaseInput) error {
	product, err := entity.NewProduct(input.Name, input.Description, input.Price, input.DepartmentID)
	if err != nil {
		return err
	}

	return u.productRepository.Create(product)
}
