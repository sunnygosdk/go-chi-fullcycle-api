package repository

import (
	"database/sql"

	"github.com/sunnygosdk/go-chi-fullcycle-api/internal/domain/entity"
)

// StoreDepartmentMapMySQLRepository is the repository for the product entity.
type StoreDepartmentMapMySQLRepository struct {
	db *sql.DB
}

// NewStoreDepartmentMapMySQLRepository creates a new product repository.
//
// Parameters:
//   - db: The database connection.
//
// Returns:
//   - *StoreDepartmentMapMySQLRepository: The new product repository.
func NewStoreDepartmentMapMySQLRepository(db *sql.DB) *StoreDepartmentMapMySQLRepository {
	return &StoreDepartmentMapMySQLRepository{db: db}
}

// Create creates a new store department map.
//
// Parameters:
//   - sdm: The store department map to create.
//
// Returns:
//   - error: An error if the product creation fails.
func (s *StoreDepartmentMapMySQLRepository) Create(sdm *entity.StoreDepartmentMap) error {
	query := "INSERT INTO store_department_map (id, store_id, department_id, created_at, updated_at, deleted_at) VALUES (?, ?, ?, ?, ?, ?)"
	_, err := s.db.Exec(query, sdm.ID.String(), sdm.StoreID.String(), sdm.DepartmentID.String(), sdm.CreatedAt, sdm.UpdatedAt, sdm.DeletedAt)
	return MapMySQLError(err)
}
