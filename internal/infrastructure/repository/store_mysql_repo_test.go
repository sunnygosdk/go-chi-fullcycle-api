package repository_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/sunnygosdk/go-chi-fullcycle-api/internal/domain/entity"
	"github.com/sunnygosdk/go-chi-fullcycle-api/internal/infrastructure/repository"
)

// TestCreateStoreMySQLRepository tests the CreateStoreMySQLRepository function
func TestCreateStoreMySQLRepository(t *testing.T) {
	truncateTables(db)
	repo := repository.NewStoreMySQLRepository(db)
	store, _ := entity.NewStore("Store 1", "Address 1")

	err := repo.Create(store)
	assert.NoError(t, err, "CreateStoreMySQLRepository should return no error")
}
