package database_test

import (
	"testing"

	_ "github.com/mattn/go-sqlite3"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/sunnygosdk/go-chi-fullcycle-api/database"
)

func TestSetupTestDB(t *testing.T) {
	db := database.SetupTestDB()
	require.NotNil(t, db, "DB connection should not be nil")
	assert.NoError(t, db.Ping(), "DB connection should be valid")

	row := db.QueryRow("SELECT name FROM sqlite_master WHERE type='table' AND name='products'")
	var tableName string
	err := row.Scan(&tableName)
	assert.NoError(t, err, "No error should be returned")
	assert.Equal(t, "products", tableName, "Table name should be 'products'")

	row = db.QueryRow("SELECT name FROM sqlite_master WHERE type='table' AND name='users'")
	err = row.Scan(&tableName)
	assert.NoError(t, err, "No error should be returned")
	assert.Equal(t, "users", tableName, "Table name should be 'users'")
}
