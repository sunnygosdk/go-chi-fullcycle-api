package stock

import (
	"github.com/sunnygosdk/go-chi-fullcycle-api/internal/domain/entity"
	"github.com/sunnygosdk/go-chi-fullcycle-api/internal/domain/repository"
)

// CreateStockUseCase is the use case for creating a stock.
type CreateStockUseCase struct {
	stockRepository repository.StockRepository
}

// CreateStockUseCaseInput is the input for creating a stock.
type CreateStockUseCaseInput struct {
	Quantity  int
	ProductID string
	StoreID   string
}

// Execute creates a new stock.
//
// Parameters:
//   - input: The input for creating a stock.
//
// Returns:
//   - error: An error if the stock creation fails.
func (u *CreateStockUseCase) Execute(input *CreateStockUseCaseInput) error {
	newStock, err := entity.NewStock(input.Quantity, input.ProductID, input.StoreID)
	if err != nil {
		return err
	}
	return u.stockRepository.Create(newStock)
}
