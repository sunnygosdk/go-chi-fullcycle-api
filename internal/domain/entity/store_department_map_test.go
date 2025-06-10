package entity_test

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/sunnygosdk/go-chi-fullcycle-api/internal/domain/entity"
	pkgEntity "github.com/sunnygosdk/go-chi-fullcycle-api/pkg/entity"
)

// TestNewStoreDepartmentMap tests the NewStoreDepartmentMap function.
func TestNewStoreDepartmentMap(t *testing.T) {
	storeID, _ := pkgEntity.ParseID("1")
	departmentID, _ := pkgEntity.ParseID("1")
	storeDepartmentMap, err := entity.NewStoreDepartmentMap(storeID.String(), departmentID.String())
	assert.NoError(t, err, "NewStoreDepartmentMap should return no error")
	assert.NotNil(t, storeDepartmentMap, "NewStoreDepartmentMap should return a valid store department map")
	assert.Equal(t, storeID, storeDepartmentMap.StoreID, "StoreID should be 1")
	assert.Equal(t, departmentID, storeDepartmentMap.DepartmentID, "DepartmentID should be 1")
	assert.WithinDuration(t, time.Now(), storeDepartmentMap.CreatedAt, 1*time.Second, "CreatedAt should be close to now")
	assert.WithinDuration(t, time.Now(), storeDepartmentMap.UpdatedAt, 1*time.Second, "UpdatedAt should be close to now")
	assert.Nil(t, storeDepartmentMap.DeletedAt, "DeletedAt should be nil")
}

// TestValidateNewStoreDepartmentMap tests the ValidateNewStoreDepartmentMap function.
func TestValidateNewStoreDepartmentMap(t *testing.T) {
	storeID, _ := pkgEntity.ParseID("1")
	departmentID, _ := pkgEntity.ParseID("1")

	storeDepartmentMap1, err := entity.NewStoreDepartmentMap(storeID.String(), "1")
	assert.Nil(t, storeDepartmentMap1, "ValidateNewStoreDepartmentMap should return no entity")
	assert.NotNil(t, err, "ValidateNewStoreDepartmentMap should return an error")
	assert.ErrorIs(t, entity.ErrorSDMInvalidDepartmentID, err, "ValidateNewStoreDepartmentMap should return an error")

	storeDepartmentMap2, err := entity.NewStoreDepartmentMap("1", departmentID.String())
	assert.Nil(t, storeDepartmentMap2, "ValidateNewStoreDepartmentMap should return no entity")
	assert.NotNil(t, err, "ValidateNewStoreDepartmentMap should return an error")
	assert.ErrorIs(t, entity.ErrorSDMInvalidStoreID, err, "ValidateNewStoreDepartmentMap should return an error")
}

// TestUpdateStoreDepartmentMap tests the UpdateStoreDepartmentMap function.
func TestUpdateStoreDepartments(t *testing.T) {
	storeID, _ := pkgEntity.ParseID("1")
	departmentID, _ := pkgEntity.ParseID("1")
	storeDepartments, _ := entity.NewStoreDepartmentMap(storeID.String(), departmentID.String())

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
	storeDepartments, _ := entity.NewStoreDepartmentMap(storeID.String(), departmentID.String())

	err := storeDepartments.Update(nil, nil)
	assert.ErrorIs(t, entity.ErrorSDMAtLeastOneField, err, "ValidateUpdateStoreDepartments should return an error")

	invalidStoreID := "invalid"
	err = storeDepartments.Update(&invalidStoreID, nil)
	assert.ErrorIs(t, entity.ErrorSDMInvalidStoreID, err, "ValidateUpdateStoreDepartments should return an error")

	invalidDepartmentID := "invalid"
	err = storeDepartments.Update(nil, &invalidDepartmentID)
	assert.ErrorIs(t, entity.ErrorSDMInvalidDepartmentID, err, "ValidateUpdateStoreDepartments should return an error")
}

// TestDeleteStoreDepartments tests the DeleteStoreDepartments function.
func TestDeleteStoreDepartments(t *testing.T) {
	storeID, _ := pkgEntity.ParseID("1")
	departmentID, _ := pkgEntity.ParseID("1")
	storeDepartments, _ := entity.NewStoreDepartmentMap(storeID.String(), departmentID.String())

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
	storeDepartments, _ := entity.NewStoreDepartmentMap(storeID.String(), departmentID.String())

	storeDepartments.Delete()
	assert.NotNil(t, storeDepartments.DeletedAt, "DeletedAt should not be nil")
	assert.WithinDuration(t, time.Now(), *storeDepartments.DeletedAt, 1*time.Second, "DeletedAt should be close to now")

	err := storeDepartments.Delete()
	assert.Error(t, err, "DeleteStoreDepartments should return an error")
	assert.ErrorIs(t, entity.ErrorSDMIsDeleted, err, "Error should be ErrorSDMIsDeleted")
}
