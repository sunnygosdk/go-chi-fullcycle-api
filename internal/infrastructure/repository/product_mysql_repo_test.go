package repository_test

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/sunnygosdk/go-chi-fullcycle-api/internal/domain/entity"
	"github.com/sunnygosdk/go-chi-fullcycle-api/internal/infrastructure/repository"
	pkgEntity "github.com/sunnygosdk/go-chi-fullcycle-api/pkg/entity"
)

// TestCreateProductMySQLRepository tests the CreateProductMySQLRepository function
func TestCreateProductMySQLRepository(t *testing.T) {
	truncateTables(db)
	product, _ := entity.NewProduct("Product 1", "Description 1", 1.0, "Department 1")
	repo := repository.NewProductMySQLRepository(db)

	err := repo.Create(product)
	assert.NoError(t, err, "CreateProductMySQLRepository should return no error")
}

// TestCreateProductMySQLRepositoryError tests the CreateProductMySQLRepository function with invalid data
func TestCreateProductMySQLRepositoryError(t *testing.T) {
	truncateTables(db)
	repo := repository.NewProductMySQLRepository(db)
	product := &entity.Product{
		ID:          pkgEntity.NewID(),
		Name:        "",
		Description: "Description 2",
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
		DeletedAt:   nil,
	}

	err := repo.Create(product)
	assert.Error(t, err, "CreateProductMySQLRepository should return an error")
	assert.ErrorIs(t, err, repository.ErrorMySQLCheckConstraintViolated, "CreateProductMySQLRepository should return an error")

	product.Name = "Product 2"
	product.Description = ""
	err = repo.Create(product)
	assert.Error(t, err, "CreateProductMySQLRepository should return an error")
	assert.ErrorIs(t, err, repository.ErrorMySQLCheckConstraintViolated, "CreateProductMySQLRepository should return an error")

	product.Description = "Description 2"
	err = repo.Create(product)
	assert.Nil(t, err, "CreateProductMySQLRepository should return no error")

	product.Name = "Product 2"
	err = repo.Create(product)
	assert.Error(t, err, "CreateProductMySQLRepository should return an error")
	assert.ErrorIs(t, err, repository.ErrorMySQLDuplicateEntry, "CreateProductMySQLRepository should return an error")
}
