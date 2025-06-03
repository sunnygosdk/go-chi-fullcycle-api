package user

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/sunnygosdk/go-chi-fullcycle-api/internal/model/user"
	"github.com/sunnygosdk/go-chi-fullcycle-api/pkg/helper"
	"github.com/sunnygosdk/go-chi-fullcycle-api/test/database"
)

func TestCreateUser(t *testing.T) {
	db := database.SetupTestDB()
	defer db.Close()

	userRepo := NewRepository(db)

	user, _ := user.ToCreate(
		user.CreateUserDTO{
			Name:     "John Doe",
			Email:    "john.doe@example.com",
			Password: "Test@123",
		},
	)
	err := userRepo.Create(user)
	assert.NotNil(t, user, "User should not be nil")
	assert.NoError(t, err, "CreateUser should return no error")
	assert.Equal(t, user.Name, "John Doe", "Name should be John Doe")
	assert.Equal(t, user.Email, "john.doe@example.com", "Email should be john.doe@example.com")
	assert.Equal(t, true, user.ValidatePassword("Test@123"), "Password should be valid if it matches")
}

func TestGetUsers(t *testing.T) {
	db := database.SetupTestDB()
	defer db.Close()

	userRepo := NewRepository(db)

	user, _ := user.ToCreate(
		user.CreateUserDTO{
			Name:     "John Doe",
			Email:    "john.doe@example.com",
			Password: "Test@123",
		},
	)

	err := userRepo.Create(user)
	assert.NotNil(t, user, "User should not be nil")
	assert.NoError(t, err, "CreateUser should return no error")

	users, err := userRepo.GetUsers()
	assert.Equal(t, user.Name, users[0].Name, "Name should be John Doe")
	assert.Equal(t, user.Email, users[0].Email, "Email should be john.doe@example.com")
	assert.Equal(t, true, users[0].ValidatePassword("Test@123"), "Password should be valid if it matches")
	assert.NoError(t, err, "GetUsers should return no error")
	assert.Len(t, users, 1, "GetUsers should return 1 user")
}

func TestGetUserByID(t *testing.T) {
	db := database.SetupTestDB()
	defer db.Close()

	userRepo := NewRepository(db)

	user, _ := user.ToCreate(
		user.CreateUserDTO{
			Name:     "John Doe",
			Email:    "john.doe@example.com",
			Password: "Test@123",
		},
	)

	err := userRepo.Create(user)
	assert.NotNil(t, user, "User should not be nil")
	assert.NoError(t, err, "CreateUser should return no error")

	userByID, err := userRepo.GetUserByID(user.ID.String())
	assert.Equal(t, user.Name, userByID.Name, "Name should be John Doe")
	assert.Equal(t, user.Email, userByID.Email, "Email should be john.doe@example.com")
	assert.Equal(t, true, userByID.ValidatePassword("Test@123"), "Password should be valid if it matches")
	assert.NoError(t, err, "GetUserByID should return no error")
}

func TestGetUserByEmail(t *testing.T) {
	db := database.SetupTestDB()
	defer db.Close()

	userRepo := NewRepository(db)

	user, _ := user.ToCreate(
		user.CreateUserDTO{
			Name:     "John Doe",
			Email:    "john.doe@example.com",
			Password: "Test@123",
		},
	)

	err := userRepo.Create(user)
	assert.NotNil(t, user, "User should not be nil")
	assert.NoError(t, err, "CreateUser should return no error")

	userByEmail, err := userRepo.GetUserByEmail(user.Email)
	assert.Equal(t, user.Name, userByEmail.Name, "Name should be John Doe")
	assert.Equal(t, user.Email, userByEmail.Email, "Email should be john.doe@example.com")
	assert.Equal(t, true, userByEmail.ValidatePassword("Test@123"), "Password should be valid if it matches")
	assert.NoError(t, err, "GetUserByEmail should return no error")
}

func TestUpdateUser(t *testing.T) {
	db := database.SetupTestDB()
	defer db.Close()

	userRepo := NewRepository(db)

	userCreated, _ := user.ToCreate(
		user.CreateUserDTO{
			Name:     "John Doe",
			Email:    "john.doe@example.com",
			Password: "Test@123",
		},
	)

	err := userRepo.Create(userCreated)
	assert.NotNil(t, userCreated, "User should not be nil")
	assert.NoError(t, err, "CreateUser should return no error")

	userUpdate, _ := userCreated.ToUpdate(
		user.UpdateUserDTO{
			Name:     helper.StrPtr("John Doe Updated"),
			Email:    helper.StrPtr("john.doe.updated@example.com"),
			Password: helper.StrPtr("Test@1234"),
		},
	)

	err = userRepo.Update(userCreated.ID, userUpdate)
	assert.Equal(t, userUpdate.Name, "John Doe Updated", "Name should be John Doe Updated")
	assert.Equal(t, userUpdate.Email, "john.doe.updated@example.com", "Email should be john.doe.updated@example.com")
	assert.Equal(t, true, userUpdate.ValidatePassword("Test@1234"), "Password should be valid if it matches")
	assert.NoError(t, err, "UpdateUser should return no error")
}

func TestDeleteUser(t *testing.T) {
	db := database.SetupTestDB()
	defer db.Close()

	userRepo := NewRepository(db)

	userCreated, _ := user.ToCreate(
		user.CreateUserDTO{
			Name:     "John Doe",
			Email:    "john.doe@example.com",
			Password: "Test@123",
		},
	)

	err := userRepo.Create(userCreated)
	assert.NotNil(t, userCreated, "User should not be nil")
	assert.NoError(t, err, "CreateUser should return no error")

	err = userRepo.Delete(userCreated.ID)
	assert.NoError(t, err, "DeleteUser should return no error")

	userByID, err := userRepo.GetUserByID(userCreated.ID.String())
	assert.Error(t, err, "GetUserByID should return error")
	assert.Nil(t, userByID, "User should be nil")
}
