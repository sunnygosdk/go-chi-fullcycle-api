package user

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/sunnygosdk/go-chi-fullcycle-api/pkg/helper"
)

func TestToCreateUser(t *testing.T) {
	user, err := ToCreate(CreateUserDTO{
		Name:     "John Doe",
		Email:    "john.doe@example.com",
		Password: "Test@123",
	})

	assert.Nil(t, err, "NewUser should return no error")
	assert.NotNil(t, user, "NewUser should return a valid user")
	assert.Equal(t, "John Doe", user.Name, "Name should be John Doe")
	assert.Equal(t, "john.doe@example.com", user.Email, "Email should be john.doe@example.com")
	assert.NotEmpty(t, user.ID, "ID should not be empty")
	assert.NotEmpty(t, user.Password, "Password should not be empty")
}

func TestUserValidatePassword(t *testing.T) {
	user, err := ToCreate(CreateUserDTO{
		Name:     "John Doe",
		Email:    "john.doe@example.com",
		Password: "Test@123",
	})
	assert.Nil(t, err, "NewUser should return no error")
	assert.NotNil(t, user, "NewUser should return a valid user")
	assert.True(t, user.ValidatePassword("Test@123"), "Password should be valid if it matches")
	assert.False(t, user.ValidatePassword("Test@1234"), "Password should be invalid if it does not match")
	assert.NotEqual(t, user.Password, "Test@123", "Password should not be equal")
}

func TestValidateNewUser(t *testing.T) {
	user, err := ToCreate(CreateUserDTO{
		Name:     "John Doe",
		Email:    "john.doe@example.com",
		Password: "Test@123",
	})
	assert.Nil(t, err, "NewUser should return no error")
	assert.NotNil(t, user, "NewUser should return a valid user")
	assert.Equal(t, "John Doe", user.Name, "Name should be John Doe")
	assert.Equal(t, "john.doe@example.com", user.Email, "Email should be john.doe@example.com")
	assert.NotEmpty(t, user.ID, "ID should not be empty")
	assert.NotEmpty(t, user.Password, "Password should not be empty")
}

func TestValidateNameIsRequired(t *testing.T) {
	_, err := ToCreate(CreateUserDTO{
		Name:     "",
		Email:    "john.doe@example.com",
		Password: "Test@123",
	})

	assert.Error(t, err, ErrNameRequired, "ValidateNameIsRequired should return error")
}

func TestValidateEmailIsRequired(t *testing.T) {
	_, err := ToCreate(CreateUserDTO{
		Name:     "John Doe",
		Email:    "",
		Password: "Test@123",
	})

	assert.Error(t, err, ErrEmailRequired, "ValidateEmailIsRequired should return error")
}

func TestValidateInvalidEmail(t *testing.T) {
	_, err := ToCreate(CreateUserDTO{
		Name:     "John Doe",
		Email:    "john.doeexample.com",
		Password: "Test@123",
	})

	assert.Error(t, err, ErrInvalidEmail, "ValidateInvalidEmail should return error")
}

func TestValidatePasswordIsRequired(t *testing.T) {
	_, err := ToCreate(CreateUserDTO{
		Name:     "John Doe",
		Email:    "john.doe@example.com",
		Password: "",
	})

	assert.Error(t, err, ErrPasswordRequired, "ValidatePasswordIsRequired should return error")
}

func TestValidateWeakPassword(t *testing.T) {
	_, err := ToCreate(CreateUserDTO{
		Name:     "John Doe",
		Email:    "john.doe@example.com",
		Password: "test",
	})

	assert.Error(t, err, ErrWeakPassword, "ValidateWeakPassword should return error")
}

func TestToUpdateUser(t *testing.T) {
	user, _ := ToCreate(CreateUserDTO{
		Name:     "John Doe",
		Email:    "john.doe@example.com",
		Password: "Test@123",
	})

	updateOptions := UpdateUserDTO{
		Name:     helper.StrPtr("John Doe Updated"),
		Email:    helper.StrPtr("john.doe.updated@example.com"),
		Password: helper.StrPtr("Test@1234"),
	}

	user, err := user.ToUpdate(updateOptions)
	assert.Nil(t, err, "ValidateNewUser should return no error")
	assert.Equal(t, "John Doe Updated", user.Name, "Name should be John Doe Updated")
	assert.Equal(t, "john.doe.updated@example.com", user.Email, "Email should be john.doe.updated@example.com")
	assert.Equal(t, true, user.ValidatePassword("Test@1234"), "Password should be valid if it matches")
}

func TestValidateUpdateName(t *testing.T) {
	user, _ := ToCreate(CreateUserDTO{
		Name:     "John Doe",
		Email:    "john.doe@example.com",
		Password: "Test@123",
	})

	userDTO := UpdateUserDTO{
		Name:     helper.StrPtr("John Doe Updated"),
		Email:    nil,
		Password: nil,
	}

	user, err := user.ToUpdate(userDTO)
	assert.Nil(t, err, "ValidateUpdateName should return no error")
	assert.Equal(t, "John Doe Updated", user.Name, "Name should be John Doe Updated")
	assert.Equal(t, "john.doe@example.com", user.Email, "Email should be john.doe@example.com")
	assert.Equal(t, true, user.ValidatePassword("Test@123"), "Password should be valid if it matches")
}

func TestValidateUpdateEmail(t *testing.T) {
	user, _ := ToCreate(CreateUserDTO{
		Name:     "John Doe",
		Email:    "john.doe@example.com",
		Password: "Test@123",
	})

	userDTO := UpdateUserDTO{
		Name:     nil,
		Email:    helper.StrPtr("john.doe.updated@example.com"),
		Password: nil,
	}

	user, err := user.ToUpdate(userDTO)
	assert.Nil(t, err, "ValidateUpdateEmail should return no error")
	assert.Equal(t, "john.doe.updated@example.com", user.Email, "Email should be john.doe.updated@example.com")
	assert.Equal(t, "John Doe", user.Name, "Name should be John Doe")
	assert.Equal(t, true, user.ValidatePassword("Test@123"), "Password should be valid if it matches")
}

func TestValidateUpdatePassword(t *testing.T) {
	user, _ := ToCreate(CreateUserDTO{
		Name:     "John Doe",
		Email:    "john.doe@example.com",
		Password: "Test@123",
	})

	userDTO := UpdateUserDTO{
		Name:     nil,
		Email:    nil,
		Password: helper.StrPtr("teste"),
	}

	_, err := user.ToUpdate(userDTO)
	assert.Equal(t, ErrWeakPassword, err, "ValidateUpdatePassword should return ErrWeakPassword")

	userDTO2 := UpdateUserDTO{
		Name:     nil,
		Email:    nil,
		Password: helper.StrPtr("Test@1234"),
	}

	user, err = user.ToUpdate(userDTO2)
	assert.Nil(t, err, "ValidateUpdatePassword should return no error")
	assert.Equal(t, user.ValidatePassword("Test@1234"), true, "Password should be valid if it matches")
	assert.Equal(t, "John Doe", user.Name, "Name should be John Doe")
	assert.Equal(t, "john.doe@example.com", user.Email, "Email should be john.doe@example.com")
}
