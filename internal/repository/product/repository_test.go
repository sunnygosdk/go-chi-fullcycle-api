package product

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/sunnygosdk/go-chi-fullcycle-api/internal/model/product"
	"github.com/sunnygosdk/go-chi-fullcycle-api/pkg/helper"
	"github.com/sunnygosdk/go-chi-fullcycle-api/test/database"
)

func TestCreateProduct(t *testing.T) {
	db := database.SetupTestDB()
	defer db.Close()

	productRepo := NewRepository(db)

	product, _ := product.ToCreate(
		product.CreateProductDTO{
			Name:  "Product",
			Price: 10.0,
		})

	err := productRepo.Create(product)
	assert.NotNil(t, product, "Create should return a valid product")
	assert.NoError(t, err, "Create should return no error")
}

func TestGetProducts(t *testing.T) {
	db := database.SetupTestDB()
	defer db.Close()

	productRepo := NewRepository(db)

	product, _ := product.ToCreate(
		product.CreateProductDTO{
			Name:  "Product",
			Price: 10.0,
		})

	err := productRepo.Create(product)
	assert.NotNil(t, product, "Create should return a valid product")
	assert.NoError(t, err, "Create should return no error")

	products, err := productRepo.GetProducts(1, 1)
	assert.Equal(t, product.Name, products[0].Name, "GetProducts should return the correct product name")
	assert.Equal(t, product.Price, products[0].Price, "GetProducts should return the correct product price")
	assert.NoError(t, err, "GetProducts should return no error")
	assert.Len(t, products, 1, "GetProducts should return 1 product")
}

func TestProductPagination(t *testing.T) {
	db := database.SetupTestDB()
	defer db.Close()

	productRepo := NewRepository(db)

	products := make([]product.Model, 0)
	for i := range 50 {
		product, _ := product.ToCreate(
			product.CreateProductDTO{
				Name:  fmt.Sprintf("Product-%d", i),
				Price: 10.0,
			})
		products = append(products, *product)
	}

	err := productRepo.CreateBatch(products)
	assert.NotNil(t, products, "CreateBatch should return a valid product")
	assert.NoError(t, err, "CreateBatch should return no error")

	products, err = productRepo.GetProducts(1, 1)
	assert.Equal(t, "Product-0", products[0].Name, "GetProducts should return the correct product name")
	assert.Equal(t, 10.0, products[0].Price, "GetProducts should return the correct product price")
	assert.NoError(t, err, "GetProducts should return no error")
	assert.Len(t, products, 1, "GetProducts should return 1 product")

	products, err = productRepo.GetProducts(2, 1)
	assert.Equal(t, "Product-1", products[0].Name, "GetProducts should return the correct product name")
	assert.Equal(t, 10.0, products[0].Price, "GetProducts should return the correct product price")
	assert.NoError(t, err, "GetProducts should return no error")
	assert.Len(t, products, 1, "GetProducts should return 1 product")

	totalProducts, err := productRepo.GetTotalProducts()
	assert.NoError(t, err, "GetTotalProducts should return no error")
	assert.Equal(t, 50, totalProducts, "GetTotalProducts should return 50 products")

	totalPages, err := productRepo.GetTotalPages(10)
	assert.NoError(t, err, "GetTotalPages should return no error")
	assert.Equal(t, 5, totalPages, "GetTotalPages should return 5 pages")
}

func TestGetProductByID(t *testing.T) {
	db := database.SetupTestDB()
	defer db.Close()

	productRepo := NewRepository(db)

	product, _ := product.ToCreate(
		product.CreateProductDTO{
			Name:  "Product",
			Price: 10.0,
		})

	err := productRepo.Create(product)
	assert.NotNil(t, product, "Create should return a valid product")
	assert.NoError(t, err, "Create should return no error")

	productByID, err := productRepo.GetProductByID(product.ID)
	assert.Equal(t, product.Name, productByID.Name, "GetProductByID should return the correct product name")
	assert.Equal(t, product.Price, productByID.Price, "GetProductByID should return the correct product price")
	assert.NoError(t, err, "GetProductByID should return no error")
}

func TestGetProductByName(t *testing.T) {
	db := database.SetupTestDB()
	defer db.Close()

	productRepo := NewRepository(db)

	product, _ := product.ToCreate(
		product.CreateProductDTO{
			Name:  "Product",
			Price: 10.0,
		})

	err := productRepo.Create(product)
	assert.NotNil(t, product, "Create should return a valid product")
	assert.NoError(t, err, "Create should return no error")

	productByName, err := productRepo.GetProductByName(product.Name)
	assert.Equal(t, product.Name, productByName.Name, "GetProductByName should return the correct product name")
	assert.Equal(t, product.Price, productByName.Price, "GetProductByName should return the correct product price")
	assert.NoError(t, err, "GetProductByName should return no error")
}

func TestUpdateProduct(t *testing.T) {
	db := database.SetupTestDB()
	defer db.Close()

	productRepo := NewRepository(db)

	productCreated, _ := product.ToCreate(
		product.CreateProductDTO{
			Name:  "Product",
			Price: 10.0,
		})

	err := productRepo.Create(productCreated)
	assert.NotNil(t, productCreated, "Create should return a valid product")
	assert.NoError(t, err, "Create should return no error")

	productToUpdate, _ := productCreated.ToUpdate(
		product.UpdateProductDTO{
			Name:  helper.StrPtr("Product Updated"),
			Price: helper.Float64Ptr(20),
		})

	err = productRepo.Update(productCreated.ID, productToUpdate)
	assert.NoError(t, err, "Update should return no error")

	productUpdated, err := productRepo.GetProductByID(productCreated.ID)
	assert.Equal(t, productToUpdate.Name, productUpdated.Name, "Update should update the product name")
	assert.Equal(t, productToUpdate.Price, productUpdated.Price, "Update should update the product price")
	assert.NoError(t, err, "Update should return no error")
}

func TestDeleteProduct(t *testing.T) {
	db := database.SetupTestDB()
	defer db.Close()

	productRepo := NewRepository(db)

	productCreated, _ := product.ToCreate(
		product.CreateProductDTO{
			Name:  "Product",
			Price: 10.0,
		})

	err := productRepo.Create(productCreated)
	assert.NotNil(t, productCreated, "Create should return a valid product")
	assert.NoError(t, err, "Create should return no error")

	err = productRepo.Delete(productCreated.ID)
	assert.NoError(t, err, "Delete should return no error")

	productDeleted, err := productRepo.GetProductByID(productCreated.ID)
	assert.Nil(t, productDeleted, "Delete should delete the product")
	assert.Error(t, err, "Delete should return error")
}
