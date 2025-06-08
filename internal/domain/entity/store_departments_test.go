package entity_test

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/sunnygosdk/go-chi-fullcycle-api/internal/domain/entity"
	pkgEntity "github.com/sunnygosdk/go-chi-fullcycle-api/pkg/entity"
)

// TestNewStoreDepartments tests the NewStoreDepartments function.
func TestNewStoreDepartments(t *testing.T) {
	storeID, _ := pkgEntity.ParseID("1")
	departmentID, _ := pkgEntity.ParseID("1")
	storeDepartments, err := entity.NewStoreDepartments(storeID.String(), departmentID.String())
	assert.NoError(t, err, "NewStoreDepartments should return no error")
	assert.NotNil(t, storeDepartments, "NewStoreDepartments should return a valid store departments")
	assert.Equal(t, storeID, storeDepartments.StoreID, "StoreID should be 1")
	assert.Equal(t, departmentID, storeDepartments.DepartmentID, "DepartmentID should be 1")
	assert.WithinDuration(t, time.Now(), storeDepartments.CreatedAt, 1*time.Second, "CreatedAt should be close to now")
	assert.WithinDuration(t, time.Now(), storeDepartments.UpdatedAt, 1*time.Second, "UpdatedAt should be close to now")
	assert.Nil(t, storeDepartments.DeletedAt, "DeletedAt should be nil")
}

// TestValidateNewStoreDepartments tests the ValidateNewStoreDepartments function.
func TestValidateNewStoreDepartments(t *testing.T) {
	storeID, _ := pkgEntity.ParseID("1")
	departmentID, _ := pkgEntity.ParseID("1")

	storeDepartments1, err := entity.NewStoreDepartments(storeID.String(), "1")
	assert.Nil(t, storeDepartments1, "ValidateNewStoreDepartments should return no entity")
	assert.NotNil(t, err, "ValidateNewStoreDepartments should return an error")
	assert.Equal(t, entity.ErrorSDInvalidDepartmentID, err, "ValidateNewStoreDepartments should return an error")

	storeDepartments2, err := entity.NewStoreDepartments("1", departmentID.String())
	assert.Nil(t, storeDepartments2, "ValidateNewStoreDepartments should return no entity")
	assert.NotNil(t, err, "ValidateNewStoreDepartments should return an error")
	assert.Equal(t, entity.ErrorSDInvalidStoreID, err, "ValidateNewStoreDepartments should return an error")
}

// TestUpdateStoreDepartments tests the UpdateStoreDepartments function.
func TestUpdateStoreDepartments(t *testing.T) {
	storeID, _ := pkgEntity.ParseID("1")
	departmentID, _ := pkgEntity.ParseID("1")
	storeDepartments, _ := entity.NewStoreDepartments(storeID.String(), departmentID.String())

	storeID2, _ := pkgEntity.ParseID("2")
	departmentID2, _ := pkgEntity.ParseID("2")
	s2 := storeID2.String()
	d2 := departmentID2.String()
	createdAt := storeDepartments.CreatedAt
	storeDepartments.Update(&s2, &d2)
	assert.Equal(t, storeID2, storeDepartments.StoreID, "StoreID should be 2")
	assert.Equal(t, departmentID2, storeDepartments.DepartmentID, "DepartmentID should be 2")
	assert.Equal(t, createdAt, storeDepartments.CreatedAt, "CreatedAt should be the same")
	assert.WithinDuration(t, time.Now(), storeDepartments.UpdatedAt, 1*time.Second, "UpdatedAt should be close to now")
	assert.Nil(t, storeDepartments.DeletedAt, "DeletedAt should be nil")
}

// TestValidateUpdateStoreDepartments tests the ValidateUpdateStoreDepartments function.
func TestValidateUpdateStoreDepartments(t *testing.T) {
	storeID, _ := pkgEntity.ParseID("1")
	departmentID, _ := pkgEntity.ParseID("1")
	storeDepartments, _ := entity.NewStoreDepartments(storeID.String(), departmentID.String())

	err := storeDepartments.Update(nil, nil)
	assert.Equal(t, entity.ErrorSDAtLeastOneField, err, "ValidateUpdateStoreDepartments should return an error")

	invalidStoreID := "invalid"
	err = storeDepartments.Update(&invalidStoreID, nil)
	assert.Equal(t, entity.ErrorSDInvalidStoreID, err, "ValidateUpdateStoreDepartments should return an error")

	invalidDepartmentID := "invalid"
	err = storeDepartments.Update(nil, &invalidDepartmentID)
	assert.Equal(t, entity.ErrorSDInvalidDepartmentID, err, "ValidateUpdateStoreDepartments should return an error")
}

// TestDeleteStoreDepartments tests the DeleteStoreDepartments function.
func TestDeleteStoreDepartments(t *testing.T) {
	storeID, _ := pkgEntity.ParseID("1")
	departmentID, _ := pkgEntity.ParseID("1")
	storeDepartments, _ := entity.NewStoreDepartments(storeID.String(), departmentID.String())

	assert.Nil(t, storeDepartments.DeletedAt, "DeletedAt should be nil")
	err := storeDepartments.Delete()
	assert.NoError(t, err, "DeleteStoreDepartments should return no error")
	assert.NotNil(t, storeDepartments.DeletedAt, "DeletedAt should not be nil")
	assert.WithinDuration(t, time.Now(), *storeDepartments.DeletedAt, 1*time.Second, "DeletedAt should be close to now")
}

// TestValidateDeleteStoreDepartments tests the ValidateDeleteStoreDepartments function.
func TestValidateDeleteStoreDepartments(t *testing.T) {
	storeID, _ := pkgEntity.ParseID("1")
	departmentID, _ := pkgEntity.ParseID("1")
	storeDepartments, _ := entity.NewStoreDepartments(storeID.String(), departmentID.String())

	storeDepartments.Delete()
	assert.NotNil(t, storeDepartments.DeletedAt, "DeletedAt should not be nil")
	assert.WithinDuration(t, time.Now(), *storeDepartments.DeletedAt, 1*time.Second, "DeletedAt should be close to now")

	err := storeDepartments.Delete()
	assert.Error(t, err, "DeleteStoreDepartments should return an error")
	assert.Equal(t, entity.ErrorSDIsDeleted, err, "Error should be ErrorSDIsDeleted")
}
