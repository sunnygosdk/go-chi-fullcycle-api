package repository

import (
	"database/sql"

	"github.com/sunnygosdk/go-chi-fullcycle-api/internal/domain/entity"
)

// DepartmentMySQLRepository is the repository for the department entity.
type DepartmentMySQLRepository struct {
	db *sql.DB
}

// NewDepartmentMySQLRepository creates a new department repository.
//
// Parameters:
//   - db: The database connection.
//
// Returns:
//   - DepartmentMySQLRepository: The new department repository.
func NewDepartmentMySQLRepository(db *sql.DB) DepartmentMySQLRepository {
	return DepartmentMySQLRepository{db: db}
}

// Create creates a new department.
//
// Parameters:
//   - department: The department to create.
//
// Returns:
//   - error: An error if the department creation fails.
func (d *DepartmentMySQLRepository) Create(department *entity.Department) error {
	query := "INSERT INTO departments (id, name, description, created_at, updated_at, deleted_at) VALUES (?, ?, ?, ?, ?, ?)"
	_, err := d.db.Exec(query, department.ID.String(), department.Name, department.Description, department.CreatedAt, department.UpdatedAt, department.DeletedAt)
	if err != nil {
		return err
	}

	return nil
}
