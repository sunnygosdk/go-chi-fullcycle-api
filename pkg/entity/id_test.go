package entity_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/sunnygosdk/go-chi-fullcycle-api/pkg/entity"
)

// TestNewID tests the NewID function.
func TestNewID(t *testing.T) {
	id := entity.NewID()
	assert.NotEmpty(t, id)
}

// TestParseID tests the ParseID function.
func TestParseID(t *testing.T) {
	id, err := entity.ParseID("12345678-1234-5678-1234-567812345678")
	assert.NoError(t, err)
	assert.Equal(t, "12345678-1234-5678-1234-567812345678", id.String())
}

// TestParseIDInvalid tests the ParseID function with an invalid UUID.
func TestParseIDInvalid(t *testing.T) {
	_, err := entity.ParseID("invalid")
	assert.Error(t, err)
}
