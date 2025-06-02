package product

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/sunnygosdk/go-chi-fullcycle-api/internal/model/product"
	"github.com/sunnygosdk/go-chi-fullcycle-api/test/database"
)

func TestCreateProduct(t *testing.T) {
	db := database.SetupTestDB()
	defer db.Close()

	productRepo := NewRepository(db)

	product, _ := product.New("Product", 10.0)

	err := productRepo.Create(product)
	assert.NotNil(t, product)
	assert.NoError(t, err)
}

func TestGetProducts(t *testing.T) {
	db := database.SetupTestDB()
	defer db.Close()

	productRepo := NewRepository(db)

	product, _ := product.New("Product", 10.0)

	err := productRepo.Create(product)
	assert.NotNil(t, product)
	assert.NoError(t, err)

	products, err := productRepo.GetProducts()
	assert.Equal(t, product.Name, products[0].Name)
	assert.Equal(t, product.Price, products[0].Price)
	assert.NoError(t, err)
	assert.Len(t, products, 1)
}
