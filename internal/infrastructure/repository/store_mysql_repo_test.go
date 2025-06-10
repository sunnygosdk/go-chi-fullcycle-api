package repository_test

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/sunnygosdk/go-chi-fullcycle-api/internal/domain/entity"
	"github.com/sunnygosdk/go-chi-fullcycle-api/internal/infrastructure/repository"
	pkgEntity "github.com/sunnygosdk/go-chi-fullcycle-api/pkg/entity"
)

// TestCreateStoreMySQLRepository tests the CreateStoreMySQLRepository function
func TestCreateStoreMySQLRepository(t *testing.T) {
	repo := repository.NewStoreMySQLRepository(db)
	store, _ := entity.NewStore("Store 1", "Address 1")

	err := repo.Create(store)
	assert.NoError(t, err, "CreateStoreMySQLRepository should return no error")
}

// TestCreateStoreMySQLRepositoryError tests the CreateStoreMySQLRepository function with invalid data
func TestCreateStoreMySQLRepositoryError(t *testing.T) {
	repo := repository.NewStoreMySQLRepository(db)
	store := &entity.Store{
		ID:        pkgEntity.NewID(),
		Name:      "",
		Address:   "Address 2",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		DeletedAt: nil,
	}

	err := repo.Create(store)
	assert.Error(t, err, "CreateStoreMySQLRepository should return an error")
	assert.ErrorIs(t, err, repository.ErrorMySQLCheckConstraintViolated, "CreateStoreMySQLRepository should return an error")

	store.Name = "Store 2"
	store.Address = ""
	err = repo.Create(store)
	assert.Error(t, err, "CreateStoreMySQLRepository should return an error")
	assert.ErrorIs(t, err, repository.ErrorMySQLCheckConstraintViolated, "CreateStoreMySQLRepository should return an error")

	store.Address = "Address 2"
	err = repo.Create(store)
	assert.Nil(t, err, "CreateStoreMySQLRepository should return an error")

	store.Name = "Store 2"
	err = repo.Create(store)
	assert.Error(t, err, "CreateStoreMySQLRepository should return an error")
	assert.ErrorIs(t, err, repository.ErrorMySQLDuplicateEntry, "CreateStoreMySQLRepository should return an error")
}
