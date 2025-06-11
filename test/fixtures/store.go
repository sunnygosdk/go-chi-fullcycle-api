package fixtures

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/sunnygosdk/go-chi-fullcycle-api/internal/domain/entity"
	"github.com/sunnygosdk/go-chi-fullcycle-api/internal/domain/repository"
)

// CreateStoreFixture creates a store fixture.
//
// Parameters:
//   - t: The test instance.
//   - repo: The store repository.
//
// Returns:
//   - *entity.Store: The created store.
func CreateStoreFixture(t *testing.T, repo repository.StoreRepository) *entity.Store {
	t.Log("Creating Store Fixture")
	store, err := entity.NewStore("Fixture Store", "Fixture Store Address")
	assert.NoError(t, err, "Error creating store fixture")

	err = repo.Create(store)
	assert.NoError(t, err, "Error creating store fixture")

	return store
}
