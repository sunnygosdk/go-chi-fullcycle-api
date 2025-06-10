package fixtures

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/sunnygosdk/go-chi-fullcycle-api/internal/domain/entity"
	"github.com/sunnygosdk/go-chi-fullcycle-api/internal/domain/repository"
)

// CreateDepartmentFixture creates a department fixture.
//
// Parameters:
//   - t: The test instance.
//   - repo: The department repository.
//
// Returns:
//   - *entity.Department: The created department.
func CreateDepartmentFixture(t *testing.T, repo repository.DepartmentRepository) *entity.Department {
	department, err := entity.NewDepartment("Fixture Department", "Fixture Department Description")
	assert.NoError(t, err, "Error creating department fixture")

	err = repo.Create(department)
	assert.NoError(t, err, "Error creating department fixture")

	return department
}
