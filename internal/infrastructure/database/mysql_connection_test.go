package database_test

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/sunnygosdk/go-chi-fullcycle-api/internal/infrastructure/database"
)

// TestLoadDBConfig tests the LoadDBConfig function
func TestLoadDBConfig(t *testing.T) {
	os.Setenv("DB_USER", "root")
	os.Setenv("DB_PASSWORD", "root")
	os.Setenv("DB_HOST", "localhost")
	os.Setenv("DB_PORT", "3306")
	os.Setenv("DB_NAME", "test")

	config := database.LoadDBConfig()
	assert.Equal(t, "root", config.DBUser, "Return DB_USER")
	assert.Equal(t, "root", config.DBPassword, "Return DB_PASSWORD")
	assert.Equal(t, "localhost", config.DBHost, "Return DB_HOST")
	assert.Equal(t, "3306", config.DBPort, "Return DB_PORT")
	assert.Equal(t, "test", config.DBName, "Return DB_NAME")
}
