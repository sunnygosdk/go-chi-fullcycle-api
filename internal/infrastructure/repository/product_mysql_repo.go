package repository

import (
	"database/sql"

	"github.com/sunnygosdk/go-chi-fullcycle-api/internal/domain/entity"
)

// ProductMySQLRepository is the repository for the product entity.
type ProductMySQLRepository struct {
	db *sql.DB
}

// NewProductMySQLRepository creates a new product repository.
//
// Parameters:
//   - db: The database connection.
//
// Returns:
//   - ProductMySQLRepository: The new product repository.
func NewProductMySQLRepository(db *sql.DB) ProductMySQLRepository {
	return ProductMySQLRepository{db: db}
}

// Create creates a new product.
//
// Parameters:
//   - product: The product to create.
//
// Returns:
//   - error: An error if the product creation fails.
func (s *ProductMySQLRepository) Create(product *entity.Product) error {
	query := "INSERT INTO products (id, name, description, price, department_id, created_at, updated_at, deleted_at) VALUES (?, ?, ?, ?, ?, ?, ?, ?)"
	_, err := s.db.Exec(query, product.ID.String(), product.Name, product.Description, product.Price, product.DepartmentID.String(), product.CreatedAt, product.UpdatedAt, product.DeletedAt)
	return MapMySQLError(err)
}
