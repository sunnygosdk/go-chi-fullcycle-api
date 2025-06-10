package entity_test

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/sunnygosdk/go-chi-fullcycle-api/internal/domain/entity"
)

func TestNewStore(t *testing.T) {
	store, err := entity.NewStore("store", "address")
	assert.NoError(t, err, "NewStore should return no error")
	assert.NotNil(t, store, "NewStore should return a valid store")
	assert.Equal(t, "store", store.Name, "NewStore should return a valid store")
	assert.Equal(t, "address", store.Address, "NewStore should return a valid store")
	assert.WithinDuration(t, time.Now(), store.CreatedAt, 1*time.Second, "NewStore should return a valid store")
	assert.WithinDuration(t, time.Now(), store.UpdatedAt, 1*time.Second, "NewStore should return a valid store")
	assert.Nil(t, store.DeletedAt, "NewStore should return a valid store")
}

func TestValidateNewStore(t *testing.T) {
	store, err := entity.NewStore("", "address")
	assert.Nil(t, store, "ValidateNewStore should return no entity")
	assert.NotNil(t, err, "ValidateNewStore should return an error")
	assert.ErrorIs(t, entity.ErrorStoreInvalidName, err, "ValidateNewStore should return an error")

	store, err = entity.NewStore("store", "")
	assert.Nil(t, store, "ValidateNewStore should return no entity")
	assert.NotNil(t, err, "ValidateNewStore should return an error")
	assert.ErrorIs(t, entity.ErrorStoreInvalidAddress, err, "ValidateNewStore should return an error")

	store, err = entity.NewStore("s", "address")
	assert.Nil(t, store, "ValidateNewStore should return no entity")
	assert.NotNil(t, err, "ValidateNewStore should return an error")
	assert.ErrorIs(t, entity.ErrorStoreMinLengthName, err, "ValidateNewStore should return an error")

	store, err = entity.NewStore("store", "a")
	assert.Nil(t, store, "ValidateNewStore should return no entity")
	assert.NotNil(t, err, "ValidateNewStore should return an error")
	assert.ErrorIs(t, entity.ErrorStoreMinLengthAddress, err, "ValidateNewStore should return an error")
}

func TestUpdateStore(t *testing.T) {
	store, _ := entity.NewStore("store", "address")

	createdAt := store.CreatedAt
	newName := "new store"
	newAddress := "new address"
	store.Update(&newName, &newAddress)

	assert.Equal(t, newName, store.Name, "UpdateStore should change the name")
	assert.Equal(t, newAddress, store.Address, "UpdateStore should change the address")
	assert.Equal(t, createdAt, store.CreatedAt, "UpdateStore should not change the creation timestamp")
	assert.WithinDuration(t, time.Now(), store.UpdatedAt, 1*time.Second, "UpdateStore should change the update timestamp")
}

func TestValidateUpdateStore(t *testing.T) {
	store, _ := entity.NewStore("store", "address")

	err := store.Update(nil, nil)
	assert.NotNil(t, err, "ValidateUpdateStore should return an error")
	assert.ErrorIs(t, entity.ErrorStoreAtLeastOneField, err, "ValidateUpdateStore should return an error")

	invalidName := ""
	err = store.Update(&invalidName, nil)
	assert.NotNil(t, err, "ValidateUpdateStore should return an error")
	assert.ErrorIs(t, entity.ErrorStoreInvalidName, err, "ValidateUpdateStore should return an error")

	invalidAddress := ""
	err = store.Update(nil, &invalidAddress)
	assert.NotNil(t, err, "ValidateUpdateStore should return an error")
	assert.ErrorIs(t, entity.ErrorStoreInvalidAddress, err, "ValidateUpdateStore should return an error")

	invalidName = "s"
	err = store.Update(&invalidName, nil)
	assert.NotNil(t, err, "ValidateUpdateStore should return an error")
	assert.ErrorIs(t, entity.ErrorStoreMinLengthName, err, "ValidateUpdateStore should return an error")

	invalidAddress = "a"
	err = store.Update(nil, &invalidAddress)
	assert.NotNil(t, err, "ValidateUpdateStore should return an error")
	assert.ErrorIs(t, entity.ErrorStoreMinLengthAddress, err, "ValidateUpdateStore should return an error")
}

func TestDeleteStore(t *testing.T) {
	store, _ := entity.NewStore("store", "address")
	store.Delete()
	assert.NotNil(t, store.DeletedAt, "DeleteStore should mark the store as deleted")
	assert.WithinDuration(t, time.Now(), *store.DeletedAt, 1*time.Second, "DeleteStore should mark the store as deleted")
}

func TestValidateDeleteStore(t *testing.T) {
	store, _ := entity.NewStore("store", "address")
	assert.Nil(t, store.DeletedAt, "ValidateDeleteStore should mark the store as deleted")
	store.Delete()
	assert.NotNil(t, store.DeletedAt, "ValidateDeleteStore should mark the store as deleted")

	err := store.Delete()
	assert.NotNil(t, err, "ValidateDeleteStore should return an error")
	assert.ErrorIs(t, entity.ErrorStoreIsDeleted, err, "ValidateDeleteStore should return an error")
}
