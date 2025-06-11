package fixtures

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/sunnygosdk/go-chi-fullcycle-api/internal/domain/entity"
	"github.com/sunnygosdk/go-chi-fullcycle-api/internal/domain/repository"
	pkgEntity "github.com/sunnygosdk/go-chi-fullcycle-api/pkg/entity"
)

// CreateProductFixture creates a product fixture.
//
// Parameters:
//   - t: The test instance.
//   - repo: The product repository.
//   - departmentID: The department ID.
//
// Returns:
//   - *entity.Product: The created product.
func CreateProductFixture(t *testing.T, repo repository.ProductRepository, departmentID pkgEntity.ID) *entity.Product {
	t.Log("Creating Product Fixture")
	product, err := entity.NewProduct("Fixture Product", "Fixture Product Description", 100.0, departmentID.String())
	assert.NoError(t, err, "Error creating product fixture")

	err = repo.Create(product)
	assert.NoError(t, err, "Error creating product fixture")

	return product
}
