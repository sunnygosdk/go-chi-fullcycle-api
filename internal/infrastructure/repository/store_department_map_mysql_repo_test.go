package repository_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/sunnygosdk/go-chi-fullcycle-api/internal/domain/entity"
	"github.com/sunnygosdk/go-chi-fullcycle-api/internal/infrastructure/repository"
	"github.com/sunnygosdk/go-chi-fullcycle-api/test/fixtures"
)

// TestCreateStoreDepartmentMapMySQLRepository tests the CreateStoreDepartmentMapMySQLRepository function
func TestCreateStoreDepartmentMapMySQLRepository(t *testing.T) {
	truncateTables(db)

	t.Log("TestCreateStoreDepartmentMapMySQLRepository")
	storeRepo := repository.NewStoreMySQLRepository(db)
	departmentRepo := repository.NewDepartmentMySQLRepository(db)
	department := fixtures.CreateDepartmentFixture(t, departmentRepo)
	store := fixtures.CreateStoreFixture(t, storeRepo)

	sdmRepo := repository.NewStoreDepartmentMapMySQLRepository(db)
	sdm, _ := entity.NewStoreDepartmentMap(store.ID.String(), department.ID.String())

	t.Log("Starting Store Department Map Repository Test - Create function")
	err := sdmRepo.Create(sdm)
	assert.NoError(t, err, "CreateStoreDepartmentMapMySQLRepository should return no error")
	t.Log("Finished Store Department Map Repository Test - Create function")
}
