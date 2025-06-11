package repository

import (
	"database/sql"

	"github.com/sunnygosdk/go-chi-fullcycle-api/internal/domain/entity"
)

// StockMySQLRepository is the repository for the stock entity.
type StockMySQLRepository struct {
	db *sql.DB
}

// NewStockMySQLRepository creates a new stock repository.
//
// Parameters:
//   - db: The database connection.
//
// Returns:
//   - *StockMySQLRepository: The new stock repository.
func NewStockMySQLRepository(db *sql.DB) *StockMySQLRepository {
	return &StockMySQLRepository{db: db}
}

// Create creates a new stock.
//
// Parameters:
//   - stock: The stock to create.
//
// Returns:
//   - error: An error if the stock creation fails.
func (s *StockMySQLRepository) Create(stock *entity.Stock) error {
	query := "INSERT INTO stocks (id, quantity, product_id, store_id, created_at, updated_at, deleted_at) VALUES (?, ?, ?, ?, ?, ?, ?)"
	_, err := s.db.Exec(query, stock.ID.String(), stock.Quantity, stock.ProductID.String(), stock.StoreID.String(), stock.CreatedAt, stock.UpdatedAt, stock.DeletedAt)
	return MapMySQLError(err)
}
