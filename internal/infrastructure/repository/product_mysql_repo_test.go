package repository_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/sunnygosdk/go-chi-fullcycle-api/internal/domain/entity"
	"github.com/sunnygosdk/go-chi-fullcycle-api/internal/infrastructure/repository"
	"github.com/sunnygosdk/go-chi-fullcycle-api/test/fixtures"
)

// TestCreateProductMySQLRepository tests the CreateProductMySQLRepository function
func TestCreateProductMySQLRepository(t *testing.T) {
	truncateTables(db)

	t.Log("TestCreateProductMySQLRepository")
	productRepo := repository.NewProductMySQLRepository(db)
	departmentRepo := repository.NewDepartmentMySQLRepository(db)
	department := fixtures.CreateDepartmentFixture(t, departmentRepo)
	product, _ := entity.NewProduct("Product 1", "Description 1", 1.0, department.ID.String())

	t.Log("Starting Product Repository Test - Create function")
	err := productRepo.Create(product)
	assert.NoError(t, err, "CreateProductMySQLRepository should return no error")
	t.Log("Finished Product Repository Test - Create function")
}
