package product

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/sunnygosdk/go-chi-fullcycle-api/pkg/entity"
)

func TestNewProduct(t *testing.T) {
	product, err := New("Product", 10)
	assert.NoError(t, err)
	assert.NotNil(t, product)
	assert.Equal(t, "Product", product.Name)
	assert.Equal(t, 10.0, product.Price)
	assert.NotEmpty(t, product.ID)
	assert.NotEmpty(t, product.CreatedAt)
}

func TestValidateNewProduct(t *testing.T) {
	product := &Model{
		ID:    entity.NewID(),
		Name:  "Product",
		Price: 10,
	}

	err := product.ValidateNewProduct()
	assert.NoError(t, err)
}

func TestValidateNameIsRequired(t *testing.T) {
	product, err := New("", 10)
	assert.Nil(t, product)
	assert.Error(t, err)
	assert.Equal(t, ErrNameRequired, err)
}

func TestValidatePriceIsLessOrZero(t *testing.T) {
	product, err := New("Product", -1)
	assert.Nil(t, product)
	assert.Error(t, err)
	assert.Equal(t, ErrPriceLessOrZero, err)
}
