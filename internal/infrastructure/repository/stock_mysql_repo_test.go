package repository_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/sunnygosdk/go-chi-fullcycle-api/internal/domain/entity"
	"github.com/sunnygosdk/go-chi-fullcycle-api/internal/infrastructure/repository"
	"github.com/sunnygosdk/go-chi-fullcycle-api/test/fixtures"
)

// TestCreateStockMySQLRepository tests the CreateStockMySQLRepository function
func TestCreateStockMySQLRepository(t *testing.T) {
	truncateTables(db)
	productRepo := repository.NewProductMySQLRepository(db)
	departmentRepo := repository.NewDepartmentMySQLRepository(db)
	department := fixtures.CreateDepartmentFixture(t, departmentRepo)
	product := fixtures.CreateProductFixture(t, productRepo, department.ID)
	storeRepo := repository.NewStoreMySQLRepository(db)
	store := fixtures.CreateStoreFixture(t, storeRepo)

	stock, _ := entity.NewStock(10, product.ID.String(), store.ID.String())
	repo := repository.NewStockMySQLRepository(db)

	err := repo.Create(stock)
	assert.NoError(t, err, "CreateStockMySQLRepository should return no error")
}
