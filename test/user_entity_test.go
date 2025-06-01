package test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/sunnygosdk/go-chi-fullcycle-api/internal/entity"
)

// TestNewUser tests the NewUser function.
// It tests if the function returns a valid user and no error.
func TestNewUser(t *testing.T) {
	user, err := entity.NewUser("John Doe", "john.doe@example.com", "123456")
	assert.Nil(t, err, "NewUser should return no error")
	assert.NotNil(t, user, "NewUser should return a valid user")
	assert.Equal(t, "John Doe", user.Name, "Name should be John Doe")
	assert.Equal(t, "john.doe@example.com", user.Email, "Email should be john.doe@example.com")
	assert.NotEmpty(t, user.ID, "ID should not be empty")
	assert.NotEmpty(t, user.Password, "Password should not be empty")
}

// TestUserValidatePassword tests the ValidatePassword function.
// It tests if the function returns true if the password matches and false if it does not match.
func TestUserValidatePassword(t *testing.T) {
	user, err := entity.NewUser("John Doe", "john.doe@example.com", "123456")
	assert.Nil(t, err, "NewUser should return no error")
	assert.NotNil(t, user, "NewUser should return a valid user")
	assert.True(t, user.ValidatePassword("123456"), "Password should be valid if it matches")
	assert.False(t, user.ValidatePassword("1234567"), "Password should be invalid if it does not match")
	assert.NotEqual(t, user.Password, "123456", "Password should not be equal")
}
