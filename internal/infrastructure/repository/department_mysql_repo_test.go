package repository_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/sunnygosdk/go-chi-fullcycle-api/internal/domain/entity"
	"github.com/sunnygosdk/go-chi-fullcycle-api/internal/infrastructure/repository"
)

// TestCreateDepartmentMySQLRepository tests the CreateDepartmentMySQLRepository function
func TestCreateDepartmentMySQLRepository(t *testing.T) {
	truncateTables(db)
	department, _ := entity.NewDepartment("Department 1", "Description 1")
	repo := repository.NewDepartmentMySQLRepository(db)

	err := repo.Create(department)
	assert.NoError(t, err, "CreateDepartmentMySQLRepository should return no error")
}
