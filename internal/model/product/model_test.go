package product

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/sunnygosdk/go-chi-fullcycle-api/pkg/helper"
)

func TestToCreateProduct(t *testing.T) {
	product, err := ToCreate(CreateProductDTO{
		Name:  "Product",
		Price: 10,
	})
	assert.NoError(t, err, "ToCreate should return no error")
	assert.NotNil(t, product, "ToCreate should return a valid product")
	assert.Equal(t, "Product", product.Name, "Name should be Product")
	assert.Equal(t, 10.0, product.Price, "Price should be 10")
	assert.NotEmpty(t, product.ID, "ID should not be empty")
	assert.NotEmpty(t, product.CreatedAt, "CreatedAt should not be empty")
}

func TestValidateNameIsRequired(t *testing.T) {
	product, err := ToCreate(CreateProductDTO{
		Name:  "",
		Price: 10,
	})

	assert.Nil(t, product, "ToCreate should return nil product")
	assert.Error(t, err, "ToCreate should return error")
	assert.Equal(t, ErrNameRequired, err, "ToCreate should return ErrNameRequired")
}

func TestValidatePriceIsLessOrZero(t *testing.T) {
	product, err := ToCreate(CreateProductDTO{
		Name:  "Product",
		Price: -1,
	})
	assert.Nil(t, product, "ToCreate should return nil product")
	assert.Error(t, err, "ToCreate should return error")
	assert.Equal(t, ErrPriceLessOrZero, err, "ToCreate should return ErrPriceLessOrZero")
}

func TestToUpdateProduct(t *testing.T) {
	product, err := ToCreate(CreateProductDTO{
		Name:  "Product",
		Price: 10,
	})
	assert.NoError(t, err, "ToCreate should return no error")
	assert.NotNil(t, product, "ToCreate should return a valid product")

	product, err = product.ToUpdate(UpdateProductDTO{
		Name:  helper.StrPtr("Product Updated"),
		Price: helper.Float64Ptr(20),
	})
	assert.NoError(t, err, "ToUpdate should return no error")
	assert.NotNil(t, product, "ToUpdate should return a valid product")
	assert.Equal(t, "Product Updated", product.Name, "Name should be Product Updated")
	assert.Equal(t, 20.0, product.Price, "Price should be 20")
}

func TestToUpdateProductWithInvalidName(t *testing.T) {
	product, err := ToCreate(CreateProductDTO{
		Name:  "Product",
		Price: 10,
	})
	assert.NoError(t, err, "ToCreate should return no error")
	assert.NotNil(t, product, "ToCreate should return a valid product")

	product, err = product.ToUpdate(UpdateProductDTO{
		Name: helper.StrPtr(""),
	})
	assert.Error(t, err, "ToUpdate should return error")
	assert.Nil(t, product, "ToUpdate should return nil product")
	assert.Equal(t, ErrNameRequired, err, "ToUpdate should return ErrNameRequired")
}

func TestToUpdateProductWithInvalidPrice(t *testing.T) {
	product, err := ToCreate(CreateProductDTO{
		Name:  "Product",
		Price: 10,
	})
	assert.NoError(t, err, "ToCreate should return no error")
	assert.NotNil(t, product, "ToCreate should return a valid product")

	product, err = product.ToUpdate(UpdateProductDTO{
		Price: helper.Float64Ptr(-1),
	})
	assert.Error(t, err, "ToUpdate should return error")
	assert.Nil(t, product, "ToUpdate should return nil product")
	assert.Equal(t, ErrPriceLessOrZero, err, "ToUpdate should return ErrPriceLessOrZero")
}
