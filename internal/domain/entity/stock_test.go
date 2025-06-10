package entity_test

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/sunnygosdk/go-chi-fullcycle-api/internal/domain/entity"
	pkgEntity "github.com/sunnygosdk/go-chi-fullcycle-api/pkg/entity"
)

// TestNewStock tests the NewStock function.
func TestNewStock(t *testing.T) {
	productID, _ := pkgEntity.ParseID("1")
	storeID, _ := pkgEntity.ParseID("1")
	stock, err := entity.NewStock(10, productID.String(), storeID.String())
	assert.NoError(t, err, "NewStock should return no error")
	assert.NotNil(t, stock, "NewStock should return a valid stock")
	assert.Equal(t, 10, stock.Quantity, "Quantity should be 10")
	assert.Equal(t, productID, stock.ProductID, "ProductID should be 1")
	assert.Equal(t, storeID, stock.StoreID, "StoreID should be 1")
	assert.WithinDuration(t, time.Now(), stock.CreatedAt, 1*time.Second, "CreatedAt should be close to now")
	assert.WithinDuration(t, time.Now(), stock.UpdatedAt, 1*time.Second, "UpdatedAt should be close to now")
	assert.Nil(t, stock.DeletedAt, "DeletedAt should be nil")
}

func TestValidateNewStock(t *testing.T) {
	productID, _ := pkgEntity.ParseID("1")
	storeID, _ := pkgEntity.ParseID("1")
	stock1, err := entity.NewStock(-1, productID.String(), storeID.String())
	assert.Error(t, err, "NewStock should return an error")
	assert.Nil(t, stock1, "NewStock should return a valid stock")
	assert.ErrorIs(t, entity.ErrorStockQuantityLessOfZero, err, "Error should be ErrorStockQuantityLessOfZero")

	stock2, err := entity.NewStock(10, "invalid", storeID.String())
	assert.Error(t, err, "NewStock should return an error")
	assert.Nil(t, stock2, "NewStock should return a valid stock")
	assert.ErrorIs(t, entity.ErrorStockInvalidProductID, err, "Error should be ErrorStockInvalidProductID")

	stock3, err := entity.NewStock(10, productID.String(), "invalid")
	assert.Error(t, err, "NewStock should return an error")
	assert.Nil(t, stock3, "NewStock should return a valid stock")
	assert.ErrorIs(t, entity.ErrorStockInvalidStoreID, err, "Error should be ErrorStockInvalidStoreID")
}

// TestUpdateStock tests the UpdateStock function.
func TestUpdateStock(t *testing.T) {
	productID, _ := pkgEntity.ParseID("1")
	storeID, _ := pkgEntity.ParseID("1")
	stock, _ := entity.NewStock(10, productID.String(), storeID.String())

	productID2, _ := pkgEntity.ParseID("2")
	storeID2, _ := pkgEntity.ParseID("2")
	p2 := productID2.String()
	s2 := storeID2.String()
	quantity := 20
	createdAt := stock.CreatedAt
	err := stock.Update(&quantity, &p2, &s2)
	assert.NoError(t, err, "UpdateStock should return no error")
	assert.NotNil(t, stock, "UpdateStock should return a valid stock")
	assert.Equal(t, 20, stock.Quantity, "Quantity should be 20")
	assert.Equal(t, productID2, stock.ProductID, "ProductID should be 2")
	assert.Equal(t, storeID2, stock.StoreID, "StoreID should be 2")
	assert.Equal(t, createdAt, stock.CreatedAt, "CreatedAt should be the same")
	assert.WithinDuration(t, time.Now(), stock.UpdatedAt, 1*time.Second, "UpdatedAt should be close to now")
	assert.Nil(t, stock.DeletedAt, "DeletedAt should be nil")
}

// TestValidateUpdateStock tests the UpdateStock function.
func TestValidateUpdateStock(t *testing.T) {
	productID, _ := pkgEntity.ParseID("1")
	storeID, _ := pkgEntity.ParseID("1")
	stock, _ := entity.NewStock(10, productID.String(), storeID.String())

	err := stock.Update(nil, nil, nil)
	assert.Error(t, err, "UpdateStock should return an error")
	assert.ErrorIs(t, entity.ErrorStockAtLeastOneField, err, "Error should be ErrorStockAtLeastOneField")

	quantity := -1
	err = stock.Update(&quantity, nil, nil)
	assert.Error(t, err, "UpdateStock should return an error")
	assert.ErrorIs(t, entity.ErrorStockQuantityLessOfZero, err, "Error should be ErrorStockQuantityLessOfZero")

	invalidProductID := "invalid"
	err = stock.Update(nil, &invalidProductID, nil)
	assert.Error(t, err, "UpdateStock should return an error")
	assert.ErrorIs(t, entity.ErrorStockInvalidProductID, err, "Error should be ErrorStockInvalidProductID")

	invalidStoreID := "invalid"
	err = stock.Update(nil, nil, &invalidStoreID)
	assert.Error(t, err, "UpdateStock should return an error")
	assert.ErrorIs(t, entity.ErrorStockInvalidStoreID, err, "Error should be ErrorStockInvalidStoreID")
}

// TestDeleteStock tests the DeleteStock function.
func TestDeleteStock(t *testing.T) {
	productID, _ := pkgEntity.ParseID("1")
	storeID, _ := pkgEntity.ParseID("1")
	stock, _ := entity.NewStock(10, productID.String(), storeID.String())
	assert.Nil(t, stock.DeletedAt, "DeletedAt should be nil")
	err := stock.Delete()
	assert.NoError(t, err, "DeleteStock should return no error")
	assert.WithinDuration(t, time.Now(), *stock.DeletedAt, 1*time.Second, "DeletedAt should be close to now")
}

// TestValidateDeleteStock tests the DeleteStock function.
func TestValidateDeleteStock(t *testing.T) {
	productID, _ := pkgEntity.ParseID("1")
	storeID, _ := pkgEntity.ParseID("1")
	stock, _ := entity.NewStock(10, productID.String(), storeID.String())
	stock.Delete()
	assert.NotNil(t, stock.DeletedAt, "DeletedAt should not be nil")
	assert.WithinDuration(t, time.Now(), *stock.DeletedAt, 1*time.Second, "DeletedAt should be close to now")

	err := stock.Delete()
	assert.Error(t, err, "DeleteStock should return an error")
	assert.ErrorIs(t, entity.ErrorStockIsDeleted, err, "Error should be ErrorStockIsDeleted")
}
