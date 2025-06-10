package entity_test

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/sunnygosdk/go-chi-fullcycle-api/internal/domain/entity"
)

// TestNewDepartment tests the NewDepartment function.
func TestNewDepartment(t *testing.T) {
	department, err := entity.NewDepartment("Department", "Description")
	assert.NoError(t, err, "NewDepartment should return no error")
	assert.NotNil(t, department, "NewDepartment should return a valid department")
	assert.Equal(t, "Department", department.Name, "Name should be Department")
	assert.Equal(t, "Description", department.Description, "Description should be Description")
	assert.NotEmpty(t, department.ID, "ID should not be empty")
	assert.WithinDuration(t, time.Now(), department.CreatedAt, 1*time.Second, "CreatedAt should be close to now")
	assert.WithinDuration(t, time.Now(), department.UpdatedAt, 1*time.Second, "UpdatedAt should be close to now")
	assert.Nil(t, department.DeletedAt, "DeletedAt should be nil")
}

// TestValidateNewDepartment tests the NewDepartment function.
func TestValidateNewDepartment(t *testing.T) {
	dep1, err := entity.NewDepartment("", "Description")
	assert.Error(t, err, "NewDepartment should return an error")
	assert.Nil(t, dep1, "NewDepartment should return nil")
	assert.ErrorIs(t, entity.ErrorDepartmentNameRequired, err)

	dep2, err := entity.NewDepartment("Department", "")
	assert.Error(t, err, "NewDepartment should return an error")
	assert.Nil(t, dep2, "NewDepartment should return nil")
	assert.ErrorIs(t, entity.ErrorDepartmentDescriptionRequired, err)

	dep3, err := entity.NewDepartment("D", "Description")
	assert.Error(t, err, "NewDepartment should return an error")
	assert.Nil(t, dep3, "NewDepartment should return nil")
	assert.ErrorIs(t, entity.ErrorDepartmentNameMinLength, err)

	dep4, err := entity.NewDepartment("Department", "D")
	assert.Error(t, err, "NewDepartment should return an error")
	assert.Nil(t, dep4, "NewDepartment should return nil")
	assert.ErrorIs(t, entity.ErrorDepartmentDescriptionMinLength, err)
}

// TestUpdateDepartment tests the Update function.
func TestUpdateDepartment(t *testing.T) {
	department, _ := entity.NewDepartment("Department", "Description")
	newName := "Department Updated"
	newDescription := "Description Updated"

	createdAt := department.CreatedAt
	err := department.Update(&newName, &newDescription)
	assert.NoError(t, err, "Update should return no error")
	assert.NotNil(t, department, "Update should return a valid department")
	assert.Equal(t, "Department Updated", department.Name, "Name should be Department Updated")
	assert.Equal(t, "Description Updated", department.Description, "Description should be Description Updated")
	assert.Equal(t, createdAt, department.CreatedAt, "CreatedAt should be same")
	assert.WithinDuration(t, time.Now(), department.UpdatedAt, 1*time.Second, "UpdatedAt should be close to now")
	assert.Nil(t, department.DeletedAt, "DeletedAt should be nil")
}

// TestValidateUpdateDepartment tests the Update function.
func TestValidateUpdateDepartment(t *testing.T) {
	department, _ := entity.NewDepartment("Department", "Description")

	err := department.Update(nil, nil)
	assert.Error(t, err, "Update should return an error")
	assert.ErrorIs(t, entity.ErrorDepartmentAtLeastOneField, err, "Update should return an error")

	newShortDescription := "D"
	err = department.Update(nil, &newShortDescription)
	assert.Error(t, err, "Update should return an error")
	assert.ErrorIs(t, entity.ErrorDepartmentDescriptionMinLength, err, "Update should return an error")

	newShortName := "D"
	err = department.Update(&newShortName, nil)
	assert.Error(t, err, "Update should return an error")
	assert.ErrorIs(t, entity.ErrorDepartmentNameMinLength, err, "Update should return an error")
}

// TestDeleteDepartment tests the Delete function.
func TestDeleteDepartment(t *testing.T) {
	department, _ := entity.NewDepartment("Department", "Description")
	assert.Nil(t, department.DeletedAt, "DeletedAt should be nil")
	err := department.Delete()
	assert.NoError(t, err, "Delete should return no error")
	assert.NotNil(t, department.DeletedAt, "DeletedAt should not be nil")
	assert.WithinDuration(t, time.Now(), *department.DeletedAt, 1*time.Second, "DeletedAt should be close to now")
}

// TestValidateDeleteDepartment tests the Delete function.
func TestValidateDeleteDepartment(t *testing.T) {
	department, _ := entity.NewDepartment("Department", "Description")
	department.Delete()
	assert.NotNil(t, department.DeletedAt, "DeletedAt should not be nil")
	err := department.Delete()
	assert.Error(t, err, "Delete should return an error")
	assert.ErrorIs(t, entity.ErrorDepartmentIsDeleted, err, "Delete should return an error")
}
