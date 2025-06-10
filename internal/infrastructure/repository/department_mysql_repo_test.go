package repository_test

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/sunnygosdk/go-chi-fullcycle-api/internal/domain/entity"
	"github.com/sunnygosdk/go-chi-fullcycle-api/internal/infrastructure/repository"
	pkgEntity "github.com/sunnygosdk/go-chi-fullcycle-api/pkg/entity"
)

// TestCreateDepartmentMySQLRepository tests the CreateDepartmentMySQLRepository function
func TestCreateDepartmentMySQLRepository(t *testing.T) {
	department, _ := entity.NewDepartment("Department 1", "Description 1")
	repo := repository.NewDepartmentMySQLRepository(db)

	err := repo.Create(department)
	assert.NoError(t, err, "CreateDepartmentMySQLRepository should return no error")
}

// TestCreateDepartmentMySQLRepositoryError tests the CreateDepartmentMySQLRepository function with invalid data
func TestCreateDepartmentMySQLRepositoryError(t *testing.T) {
	repo := repository.NewDepartmentMySQLRepository(db)
	department := &entity.Department{
		ID:          pkgEntity.NewID(),
		Name:        "",
		Description: "Description 2",
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
		DeletedAt:   nil,
	}

	err := repo.Create(department)
	assert.Error(t, err, "CreateDepartmentMySQLRepository should return an error")
	assert.ErrorIs(t, err, repository.ErrorMySQLCheckConstraintViolated, "CreateDepartmentMySQLRepository should return an error")

	department.Name = "Department 2"
	department.Description = ""
	err = repo.Create(department)
	assert.Error(t, err, "CreateDepartmentMySQLRepository should return an error")
	assert.ErrorIs(t, err, repository.ErrorMySQLCheckConstraintViolated, "CreateDepartmentMySQLRepository should return an error")

	department.Description = "Description 2"
	err = repo.Create(department)
	assert.Nil(t, err, "CreateDepartmentMySQLRepository should return no error")

	department.Name = "Department 2"
	err = repo.Create(department)
	assert.Error(t, err, "CreateDepartmentMySQLRepository should return an error")
	assert.ErrorIs(t, err, repository.ErrorMySQLDuplicateEntry, "CreateDepartmentMySQLRepository should return an error")
}
