package repository

import (
	"database/sql"
	"fmt"

	"github.com/sunnygosdk/go-chi-fullcycle-api/internal/domain/entity"
)

// Table name for the store entity.
const tableName = "stores"

// StoreMySQLRepository is the repository for the store entity.
type StoreMySQLRepository struct {
	db *sql.DB
}

// NewStoreMySQLRepository creates a new store repository.
//
// Parameters:
//   - db: The database connection.
//
// Returns:
//   - StoreMySQLRepository: The new store repository.
func NewStoreMySQLRepository(db *sql.DB) StoreMySQLRepository {
	return StoreMySQLRepository{db: db}
}

// Create creates a new store.
//
// Parameters:
//   - store: The store to create.
//
// Returns:
//   - error: An error if the store creation fails.
func (s *StoreMySQLRepository) Create(store *entity.Store) error {
	query := fmt.Sprintf("INSERT INTO %s (id, name, address, created_at, updated_at, deleted_at) VALUES (?, ?, ?, ?, ?, ?)", tableName)
	_, err := s.db.Exec(query, store.ID.String(), store.Name, store.Address, store.CreatedAt, store.UpdatedAt, store.DeletedAt)
	if err != nil {
		return err
	}

	return nil
}
