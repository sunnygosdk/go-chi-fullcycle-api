package repository

import (
	"database/sql"

	"github.com/sunnygosdk/go-chi-fullcycle-api/internal/domain/entity"
)

// TransactionMySQLRepository is the repository for the transaction entity.
type TransactionMySQLRepository struct {
	db *sql.DB
}

// NewTransactionMySQLRepository creates a new transaction repository.
//
// Parameters:
//   - db: The database connection.
//
// Returns:
//   - *TransactionMySQLRepository: The new transaction repository.
func NewTransactionMySQLRepository(db *sql.DB) *TransactionMySQLRepository {
	return &TransactionMySQLRepository{db: db}
}

// Create creates a new transaction.
//
// Parameters:
//   - transaction: The transaction to create.
//
// Returns:
//   - error: An error if the transaction creation fails.
func (s *TransactionMySQLRepository) Create(transaction *entity.Transaction) error {
	query := "INSERT INTO transactions (id, quantity, transaction_type, stock_id, created_at, updated_at, deleted_at) VALUES (?, ?, ?, ?, ?, ?, ?, ?)"
	_, err := s.db.Exec(query, transaction.ID.String(), transaction.Quantity, transaction.TransactionType, transaction.StockID.String(), transaction.CreatedAt, transaction.UpdatedAt, transaction.DeletedAt)
	return MapMySQLError(err)
}
