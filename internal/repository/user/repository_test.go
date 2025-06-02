package user

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/sunnygosdk/go-chi-fullcycle-api/internal/model/user"
	"github.com/sunnygosdk/go-chi-fullcycle-api/test/database"
)

func TestCreateUser(t *testing.T) {
	db := database.SetupTestDB()
	defer db.Close()

	userRepo := NewRepository(db)

	userOptions := user.CreateUserOptions{
		Name:     "User",
		Email:    "user@example.com",
		Password: "Test123!",
	}

	user, _ := user.ToCreate(userOptions)
	err := userRepo.Create(user)
	assert.NotNil(t, user)
	assert.NoError(t, err)
}

func TestGetUsers(t *testing.T) {
	db := database.SetupTestDB()
	defer db.Close()

	userRepo := NewRepository(db)

	userOptions := user.CreateUserOptions{
		Name:     "User",
		Email:    "user@example.com",
		Password: "Test123!",
	}

	user, _ := user.ToCreate(userOptions)

	err := userRepo.Create(user)
	assert.NotNil(t, user)
	assert.NoError(t, err)

	users, err := userRepo.GetUsers()
	assert.Equal(t, user.Name, users[0].Name)
	assert.Equal(t, user.Email, users[0].Email)
	assert.NoError(t, err)
	assert.Len(t, users, 1)
}
