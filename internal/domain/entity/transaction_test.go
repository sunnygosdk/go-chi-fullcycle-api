package entity_test

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/sunnygosdk/go-chi-fullcycle-api/internal/domain/entity"
	pkgEntity "github.com/sunnygosdk/go-chi-fullcycle-api/pkg/entity"
)

// TestNewTransaction tests the NewTransaction function.
func TestNewTransaction(t *testing.T) {
	productID, _ := pkgEntity.ParseID("1")
	stockID, _ := pkgEntity.ParseID("1")
	transaction, err := entity.NewTransaction(10, entity.TransactionTypeIn, productID.String(), stockID.String())
	assert.NoError(t, err, "NewTransaction should return no error")
	assert.NotNil(t, transaction, "NewTransaction should return a valid transaction")
	assert.Equal(t, 10, transaction.Quantity, "Quantity should be 10")
	assert.Equal(t, entity.TransactionTypeIn, transaction.TransactionType, "TransactionType should be TransactionTypeIn")
	assert.Equal(t, productID, transaction.ProductID, "ProductID should be 1")
	assert.Equal(t, stockID, transaction.StockID, "StockID should be 1")
	assert.WithinDuration(t, time.Now(), transaction.CreatedAt, 1*time.Second, "CreatedAt should be close to now")
	assert.WithinDuration(t, time.Now(), transaction.UpdatedAt, 1*time.Second, "UpdatedAt should be close to now")
	assert.Nil(t, transaction.DeletedAt, "DeletedAt should be nil")
}

// TestValidateNewTransaction tests the ValidateNewTransaction function.
func TestValidateNewTransaction(t *testing.T) {
	productID, _ := pkgEntity.ParseID("1")
	stockID, _ := pkgEntity.ParseID("1")
	transaction, err := entity.NewTransaction(0, entity.TransactionTypeIn, productID.String(), stockID.String())
	assert.Nil(t, transaction, "ValidateNewTransaction should return no entity")
	assert.NotNil(t, err, "ValidateNewTransaction should return an error")
	assert.Equal(t, entity.ErrorTransactionQuantityIsZero, err, "ValidateNewTransaction should return an error")
}

// TestUpdateTransaction tests the UpdateTransaction function.
func TestUpdateTransaction(t *testing.T) {
	productID, _ := pkgEntity.ParseID("1")
	stockID, _ := pkgEntity.ParseID("1")
	transaction, _ := entity.NewTransaction(10, entity.TransactionTypeIn, productID.String(), stockID.String())
	quantityToUpdate := 20
	transactionTypeToUpdate := entity.TransactionTypeOut
	productIDToUpdate, _ := pkgEntity.ParseID("2")
	stockIDToUpdate, _ := pkgEntity.ParseID("2")
	p2 := productIDToUpdate.String()
	s2 := stockIDToUpdate.String()
	createdAt := transaction.CreatedAt
	err := transaction.Update(&quantityToUpdate, &transactionTypeToUpdate, &p2, &s2)
	assert.NoError(t, err, "UpdateTransaction should return no error")
	assert.Equal(t, quantityToUpdate, transaction.Quantity, "Quantity should be 20")
	assert.Equal(t, transactionTypeToUpdate, transaction.TransactionType, "TransactionType should be TransactionTypeOut")
	assert.Equal(t, productIDToUpdate, transaction.ProductID, "ProductID should be 1")
	assert.Equal(t, stockIDToUpdate, transaction.StockID, "StockID should be 1")
	assert.Equal(t, createdAt, transaction.CreatedAt, "CreatedAt should be the same")
	assert.WithinDuration(t, time.Now(), transaction.UpdatedAt, 1*time.Second, "UpdatedAt should be close to now")
	assert.Nil(t, transaction.DeletedAt, "DeletedAt should be nil")
}

// TestValidateUpdateTransaction tests the ValidateUpdateTransaction function.
func TestValidateUpdateTransaction(t *testing.T) {
	productID, _ := pkgEntity.ParseID("1")
	stockID, _ := pkgEntity.ParseID("1")
	transaction1, _ := entity.NewTransaction(10, entity.TransactionTypeIn, productID.String(), stockID.String())
	err := transaction1.Update(nil, nil, nil, nil)
	assert.Equal(t, entity.ErrorTransactionAtLeastOneField, err, "ValidateUpdateTransaction should return an error")

	quantityToUpdate := 0
	err = transaction1.Update(&quantityToUpdate, nil, nil, nil)
	assert.Equal(t, entity.ErrorTransactionQuantityIsZero, err, "ValidateUpdateTransaction should return an error")

	invalidProductID := "invalid"
	err = transaction1.Update(nil, nil, &invalidProductID, nil)
	assert.Equal(t, entity.ErrorTransactionInvalidProductID, err, "ValidateUpdateTransaction should return an error")

	invalidStockID := "invalid"
	err = transaction1.Update(nil, nil, nil, &invalidStockID)
	assert.Equal(t, entity.ErrorTransactionInvalidStockID, err, "ValidateUpdateTransaction should return an error")
}

// TestDeleteTransaction tests the DeleteTransaction function.
func TestDeleteTransaction(t *testing.T) {
	productID, _ := pkgEntity.ParseID("1")
	stockID, _ := pkgEntity.ParseID("1")
	transaction, _ := entity.NewTransaction(10, entity.TransactionTypeIn, productID.String(), stockID.String())
	assert.Nil(t, transaction.DeletedAt, "DeletedAt should be nil")
	err := transaction.Delete()
	assert.NoError(t, err, "DeleteTransaction should return no error")
	assert.NotNil(t, transaction.DeletedAt, "DeletedAt should not be nil")
	assert.WithinDuration(t, time.Now(), *transaction.DeletedAt, 1*time.Second, "DeletedAt should be close to now")
}

// TestValidateDeleteTransaction tests the ValidateDeleteTransaction function.
func TestValidateDeleteTransaction(t *testing.T) {
	productID, _ := pkgEntity.ParseID("1")
	stockID, _ := pkgEntity.ParseID("1")
	transaction, _ := entity.NewTransaction(10, entity.TransactionTypeIn, productID.String(), stockID.String())
	transaction.Delete()
	assert.NotNil(t, transaction.DeletedAt, "DeletedAt should not be nil")
	assert.WithinDuration(t, time.Now(), *transaction.DeletedAt, 1*time.Second, "DeletedAt should be close to now")

	err := transaction.Delete()
	assert.Equal(t, entity.ErrorTransactionIsDeleted, err, "Error should be ErrorTransactionIsDeleted")
}
