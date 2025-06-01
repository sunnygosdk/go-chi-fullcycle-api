package user

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/sunnygosdk/go-chi-fullcycle-api/pkg/entity"
)

func TestNewUser(t *testing.T) {
	user, err := New("John Doe", "john.doe@example.com", "Test@123")
	assert.Nil(t, err, "NewUser should return no error")
	assert.NotNil(t, user, "NewUser should return a valid user")
	assert.Equal(t, "John Doe", user.Name, "Name should be John Doe")
	assert.Equal(t, "john.doe@example.com", user.Email, "Email should be john.doe@example.com")
	assert.NotEmpty(t, user.ID, "ID should not be empty")
	assert.NotEmpty(t, user.Password, "Password should not be empty")
}

func TestUserValidatePassword(t *testing.T) {
	user, err := New("John Doe", "john.doe@example.com", "Test@123")
	assert.Nil(t, err, "NewUser should return no error")
	assert.NotNil(t, user, "NewUser should return a valid user")
	assert.True(t, user.ValidatePassword("Test@123"), "Password should be valid if it matches")
	assert.False(t, user.ValidatePassword("Test@1234"), "Password should be invalid if it does not match")
	assert.NotEqual(t, user.Password, "Test@123", "Password should not be equal")
}

func TestValidateNewUser(t *testing.T) {
	user := &Model{
		ID:       entity.NewID(),
		Name:     "John Doe",
		Email:    "john.doe@example.com",
		Password: "Test@123",
	}

	err := user.ValidateNewUser()
	assert.Nil(t, err, "ValidateNewUser should return no error")
}

func TestValidateNameIsRequired(t *testing.T) {
	user := &Model{
		ID:       entity.NewID(),
		Name:     "",
		Email:    "john.doe@example.com",
		Password: "Test@123",
	}

	err := user.ValidateNewUser()
	assert.Error(t, err, ErrNameRequired)
}

func TestValidateEmailIsRequired(t *testing.T) {
	user := &Model{
		ID:       entity.NewID(),
		Name:     "John Doe",
		Email:    "",
		Password: "Test@123",
	}

	err := user.ValidateNewUser()
	assert.Error(t, err, ErrEmailRequired)
}

func TestValidateInvalidEmail(t *testing.T) {
	user := &Model{
		ID:       entity.NewID(),
		Name:     "John Doe",
		Email:    "john.doeexample.com",
		Password: "Test@123",
	}

	err := user.ValidateNewUser()
	assert.Error(t, err, ErrInvalidEmail)
}

func TestValidatePasswordIsRequired(t *testing.T) {
	user := &Model{
		ID:       entity.NewID(),
		Name:     "John Doe",
		Email:    "john.doe@example.com",
		Password: "",
	}

	err := user.ValidateNewUser()
	assert.Error(t, err, ErrPasswordRequired)
}

func TestValidateWeakPassword(t *testing.T) {
	user := &Model{
		ID:       entity.NewID(),
		Name:     "John Doe",
		Email:    "john.doe@example.com",
		Password: "test",
	}

	err := user.ValidateNewUser()
	assert.Error(t, err, ErrWeakPassword)
}
