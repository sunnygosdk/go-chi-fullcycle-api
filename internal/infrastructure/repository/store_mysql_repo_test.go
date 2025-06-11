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

	t.Log("TestCreateStoreMySQLRepository")
	repo := repository.NewStoreMySQLRepository(db)
	store, _ := entity.NewStore("Store 1", "Address 1")

	t.Log("Starting Store Repository Test - Create function")
	err := repo.Create(store)
	assert.NoError(t, err, "CreateStoreMySQLRepository should return no error")
	t.Log("Finished Store Repository Test - Create function")
}
