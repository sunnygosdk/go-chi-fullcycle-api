package model_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/sunnygosdk/go-chi-fullcycle-api/internal/product/model"
	"github.com/sunnygosdk/go-chi-fullcycle-api/internal/product/request"
	"github.com/sunnygosdk/go-chi-fullcycle-api/pkg/helper"
)

func TestToCreateProduct(t *testing.T) {
	product, err := request.ProductToCreate(request.CreateProductRequest{
		Name:  "Product",
		Price: 10,
	})
	assert.NoError(t, err, "ProductToCreate should return no error")
	assert.NotNil(t, product, "ProductToCreate should return a valid product")
	assert.Equal(t, "Product", product.Name, "Name should be Product")
	assert.Equal(t, 10.0, product.Price, "Price should be 10")
	assert.NotEmpty(t, product.ID, "ID should not be empty")
	assert.NotEmpty(t, product.CreatedAt, "CreatedAt should not be empty")
}

func TestValidateProductNameIsRequired(t *testing.T) {
	product, err := request.ProductToCreate(request.CreateProductRequest{
		Name:  "",
		Price: 10,
	})

	assert.Nil(t, product, "ProductToCreate should return nil product")
	assert.Error(t, err, "ProductToCreate should return error")
	assert.Equal(t, model.ErrProductNameRequired, err, "ProductToCreate should return ErrProductNameRequired")
}

func TestValidateProductPriceIsLessOrZero(t *testing.T) {
	product, err := request.ProductToCreate(request.CreateProductRequest{
		Name:  "Product",
		Price: -1,
	})
	assert.Nil(t, product, "ProductToCreate should return nil product")
	assert.Error(t, err, "ProductToCreate should return error")
	assert.Equal(t, model.ErrProductPriceLessOrZero, err, "ProductToCreate should return ErrPriceLessOrZero")
}

func TestToUpdateProduct(t *testing.T) {
	product, err := request.ProductToCreate(request.CreateProductRequest{
		Name:  "Product",
		Price: 10,
	})
	assert.NoError(t, err, "ProductToCreate should return no error")
	assert.NotNil(t, product, "ProductToCreate should return a valid product")

	product, err = product.ProductToUpdate(request.UpdateProductRequest{
		Name:  helper.StrPtr("Product Updated"),
		Price: helper.Float64Ptr(20),
	})
	assert.NoError(t, err, "ProductToUpdate should return no error")
	assert.NotNil(t, product, "ProductToUpdate should return a valid product")
	assert.Equal(t, "Product Updated", product.Name, "Name should be Product Updated")
	assert.Equal(t, 20.0, product.Price, "Price should be 20")
}

func TestToUpdateProductWithInvalidName(t *testing.T) {
	product, err := model.ProductToCreate(model.CreateProductDTO{
		Name:  "Product",
		Price: 10,
	})
	assert.NoError(t, err, "ProductToCreate should return no error")
	assert.NotNil(t, product, "ProductToCreate should return a valid product")

	product, err = product.ProductToUpdate(model.UpdateProductDTO{
		Name: helper.StrPtr(""),
	})
	assert.Error(t, err, "ProductToUpdate should return error")
	assert.Nil(t, product, "ProductToUpdate should return nil product")
	assert.Equal(t, model.ErrProductNameRequired, err, "ProductToUpdate should return ErrNameRequired")
}

func TestToUpdateProductWithInvalidPrice(t *testing.T) {
	product, err := model.ProductToCreate(model.CreateProductDTO{
		Name:  "Product",
		Price: 10,
	})
	assert.NoError(t, err, "ProductToCreate should return no error")
	assert.NotNil(t, product, "ProductToCreate should return a valid product")

	product, err = product.ProductToUpdate(model.UpdateProductDTO{
		Price: helper.Float64Ptr(-1),
	})
	assert.Error(t, err, "ProductToUpdate should return error")
	assert.Nil(t, product, "ProductToUpdate should return nil product")
	assert.Equal(t, model.ErrProductPriceLessOrZero, err, "ProductToUpdate should return ErrPriceLessOrZero")
}
