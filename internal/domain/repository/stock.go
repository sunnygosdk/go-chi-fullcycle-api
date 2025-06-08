package repository

import "github.com/sunnygosdk/go-chi-fullcycle-api/internal/domain/entity"

// StockRepository is the interface for stock repository.
type StockRepository interface {
	// Create creates a new stock.
	Create(stock *entity.Stock) error
}
