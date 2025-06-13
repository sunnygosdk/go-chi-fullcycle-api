package entity_test

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/sunnygosdk/go-chi-fullcycle-api/internal/domain/entity"
)

// TestNewProduct tests the NewProduct function.
func TestNewProduct(t *testing.T) {
	departmentID := entity.NewID()
	product, err := entity.NewProduct("Product", "Description", 10, departmentID.String())
	assert.NoError(t, err, "NewProduct should return no error")
	assert.NotNil(t, product, "NewProduct should return a valid product")
	assert.Equal(t, "Product", product.Name, "Name should be Product")
	assert.Equal(t, "Description", product.Description, "Description should be Description")
	assert.Equal(t, 10.0, product.Price, "Price should be 10")
	assert.Equal(t, departmentID, product.DepartmentID, "DepartmentID should be 1")
	assert.WithinDuration(t, time.Now(), product.CreatedAt, 1*time.Second, "CreatedAt should be close to now")
	assert.WithinDuration(t, time.Now(), product.UpdatedAt, 1*time.Second, "UpdatedAt should be close to now")
	assert.Nil(t, product.DeletedAt, "DeletedAt should be nil")
}

// TestValidateNewProduct tests the NewProduct function validation.
func TestValidateNewProduct(t *testing.T) {
	departmentID := entity.NewID()
	product, err := entity.NewProduct("", "Description", 10, departmentID.String())
	assert.Error(t, err, "NewProduct should return an error")
	assert.Nil(t, product, "NewProduct should return nil")
	assert.ErrorIs(t, entity.ErrorProductNameRequired, err)

	product, err = entity.NewProduct("Product", "", 10, departmentID.String())
	assert.Error(t, err, "NewProduct should return an error")
	assert.Nil(t, product, "NewProduct should return nil")
	assert.ErrorIs(t, entity.ErrorProductDescriptionRequired, err)

	product, err = entity.NewProduct("Product", "Description", -1, departmentID.String())
	assert.Error(t, err, "NewProduct should return an error")
	assert.Nil(t, product, "NewProduct should return nil")
	assert.ErrorIs(t, entity.ErrorProductPriceLessOrZero, err)

	departmentIDInvalid := "invalid"
	product, err = entity.NewProduct("Product", "Description", 10, departmentIDInvalid)
	assert.Error(t, err, "NewProduct should return an error")
	assert.Nil(t, product, "NewProduct should return nil")
	assert.ErrorIs(t, entity.ErrorProductInvalidDepartmentID, err)
}

// TestUpdateProduct tests the Update function.
func TestUpdateProduct(t *testing.T) {
	departmentID := entity.NewID()
	product, _ := entity.NewProduct("Product", "Description", 10, departmentID.String())
	product.Name = "Product Updated"
	product.Description = "Description Updated"
	product.Price = 20
	depIDUpdated := entity.NewID()
	departmentIDUpdated := depIDUpdated.String()
	createdAt := product.CreatedAt
	err := product.Update(&product.Name, &product.Description, &product.Price, &departmentIDUpdated)
	assert.NoError(t, err, "Update should return no error")
	assert.Equal(t, "Product Updated", product.Name, "Name should be Product Updated")
	assert.Equal(t, "Description Updated", product.Description, "Description should be Description Updated")
	assert.Equal(t, 20.0, product.Price, "Price should be 20")
	assert.Equal(t, departmentIDUpdated, product.DepartmentID.String(), "DepartmentID should be 2")
	assert.Equal(t, createdAt, product.CreatedAt, "CreatedAt should be the same")
	assert.WithinDuration(t, time.Now(), product.UpdatedAt, 1*time.Second, "UpdatedAt should be close to now")
	assert.Nil(t, product.DeletedAt, "DeletedAt should be nil")
}

// TestValidateUpdateProduct tests the Update function validation.
func TestValidateUpdateProduct(t *testing.T) {
	departmentID := entity.NewID()
	prod1, _ := entity.NewProduct("Product", "Description", 10, departmentID.String())
	err := prod1.Update(nil, nil, nil, nil)
	assert.Error(t, err, "Update should return an error")
	assert.ErrorIs(t, entity.ErrorProductAtLeastOneField, err)

	prod2, _ := entity.NewProduct("Product", "Description", 10, departmentID.String())
	newName := ""
	err = prod2.Update(&newName, nil, nil, nil)
	assert.Error(t, err, "Update should return an error")
	assert.ErrorIs(t, entity.ErrorProductNameRequired, err)

	prod3, _ := entity.NewProduct("Product", "Description", 10, departmentID.String())
	newDescription := ""
	err = prod3.Update(nil, &newDescription, nil, nil)
	assert.Error(t, err, "Update should return an error")
	assert.ErrorIs(t, entity.ErrorProductDescriptionRequired, err)

	prod4, _ := entity.NewProduct("Product", "Description", 10, departmentID.String())
	newDescription = "D"
	err = prod4.Update(nil, &newDescription, nil, nil)
	assert.Error(t, err, "Update should return an error")
	assert.ErrorIs(t, entity.ErrorProductDescriptionMinLength, err)

	prod5, _ := entity.NewProduct("Product", "Description", 10, departmentID.String())
	newPrice := -1.0
	err = prod5.Update(nil, nil, &newPrice, nil)
	assert.Error(t, err, "Update should return an error")
	assert.ErrorIs(t, entity.ErrorProductPriceLessOrZero, err)

	prod6, _ := entity.NewProduct("Product", "Description", 10, departmentID.String())
	newDepartmentID := "invalid"
	err = prod6.Update(nil, nil, nil, &newDepartmentID)
	assert.Error(t, err, "Update should return an error")
	assert.ErrorIs(t, entity.ErrorProductInvalidDepartmentID, err)
}

// TestDeleteProduct tests the Delete function.
func TestDeleteProduct(t *testing.T) {
	departmentID := entity.NewID()
	product, _ := entity.NewProduct("Product", "Description", 10, departmentID.String())
	assert.Nil(t, product.DeletedAt, "DeletedAt should be nil")
	err := product.Delete()
	assert.NoError(t, err, "Delete should return no error")
	assert.NotNil(t, product.DeletedAt, "DeletedAt should not be nil")
	assert.WithinDuration(t, time.Now(), *product.DeletedAt, 1*time.Second, "DeletedAt should be close to now")
}

// TestValidateDeleteProduct tests the Delete function validation.
func TestValidateDeleteProduct(t *testing.T) {
	departmentID := entity.NewID()
	product, _ := entity.NewProduct("Product", "Description", 10, departmentID.String())
	assert.Nil(t, product.DeletedAt, "DeletedAt should be nil")
	product.Delete()
	assert.NotNil(t, product.DeletedAt, "DeletedAt should not be nil")
	err := product.Delete()
	assert.Error(t, err, "Delete should return an error")
	assert.ErrorIs(t, entity.ErrorProductIsDeleted, err)
}
