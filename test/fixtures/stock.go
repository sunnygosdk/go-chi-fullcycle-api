package fixtures

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/sunnygosdk/go-chi-fullcycle-api/internal/domain/entity"
	"github.com/sunnygosdk/go-chi-fullcycle-api/internal/domain/repository"
	pkgEntity "github.com/sunnygosdk/go-chi-fullcycle-api/pkg/entity"
)

// CreateStockFixture creates a stock fixture.
//
// Parameters:
//   - t: The test instance.
//   - repo: The stock repository.
//   - productID: The product ID.
//   - storeID: The store ID.
//
// Returns:
//   - *entity.Stock: The created stock.
func CreateStockFixture(t *testing.T, repo repository.StockRepository, productID pkgEntity.ID, storeID pkgEntity.ID) *entity.Stock {
	t.Log("Creating Stock Fixture")
	stock, err := entity.NewStock(20, productID.String(), storeID.String())
	assert.NoError(t, err, "Error creating stock fixture")

	err = repo.Create(stock)
	assert.NoError(t, err, "Error creating stock fixture")

	return stock
}
