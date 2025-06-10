package fixtures

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/sunnygosdk/go-chi-fullcycle-api/internal/domain/entity"
	"github.com/sunnygosdk/go-chi-fullcycle-api/internal/domain/repository"
	pkgEntity "github.com/sunnygosdk/go-chi-fullcycle-api/pkg/entity"
)

// CreateStoreDepartmentMapFixture creates a store department map fixture.
//
// Parameters:
//   - t: The test instance.
//   - repo: The store department map repository.
//   - storeID: The store ID.
//   - departmentID: The department ID.
//
// Returns:
//   - *entity.StoreDepartmentMap: The created store department map.
func CreateStoreDepartmentMapFixture(t *testing.T, repo repository.StoreDepartmentMapRepository, storeID pkgEntity.ID, departmentID pkgEntity.ID) *entity.StoreDepartmentMap {
	storeDepartmentMap, err := entity.NewStoreDepartmentMap(storeID.String(), departmentID.String())
	assert.NoError(t, err, "Error creating store department map fixture")

	err = repo.Create(storeDepartmentMap)
	assert.NoError(t, err, "Error creating store department map fixture")

	return storeDepartmentMap
}
