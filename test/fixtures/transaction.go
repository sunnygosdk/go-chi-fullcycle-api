package fixtures

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/sunnygosdk/go-chi-fullcycle-api/internal/domain/entity"
	"github.com/sunnygosdk/go-chi-fullcycle-api/internal/domain/repository"
	pkgEntity "github.com/sunnygosdk/go-chi-fullcycle-api/pkg/entity"
)

// CreateTransactionInFixture creates a transaction fixture.
//
// Parameters:
//   - t: The test instance.
//   - repo: The transaction repository.
//   - productID: The product ID.
//   - stockID: The stock ID.
//
// Returns:
//   - *entity.Transaction: The created transaction.
func CreateTransactionInFixture(t *testing.T, repo repository.TransactionRepository, productID pkgEntity.ID, stockID pkgEntity.ID) *entity.Transaction {
	t.Log("Creating Transaction Fixture")
	transaction, err := entity.NewTransaction(10, entity.TransactionTypeIn, productID.String(), stockID.String())
	assert.NoError(t, err, "Error creating transaction fixture")

	err = repo.Create(transaction)
	assert.NoError(t, err, "Error creating transaction fixture")

	return transaction
}

// CreateTransactionOutFixture creates a transaction fixture.
//
// Parameters:
//   - t: The test instance.
//   - repo: The transaction repository.
//   - productID: The product ID.
//   - stockID: The stock ID.
//
// Returns:
//   - *entity.Transaction: The created transaction.
func CreateTransactionOutFixture(t *testing.T, repo repository.TransactionRepository, productID pkgEntity.ID, stockID pkgEntity.ID) *entity.Transaction {
	transaction, err := entity.NewTransaction(10, entity.TransactionTypeOut, productID.String(), stockID.String())
	assert.NoError(t, err, "Error creating transaction fixture")

	err = repo.Create(transaction)
	assert.NoError(t, err, "Error creating transaction fixture")

	return transaction
}
