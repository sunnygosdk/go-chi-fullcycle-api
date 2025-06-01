package user

import (
	"golang.org/x/crypto/bcrypt"

	"github.com/sunnygosdk/go-chi-fullcycle-api/pkg/entity"
)

type Model struct {
	ID       entity.ID `json:"id"`
	Name     string    `json:"name"`
	Email    string    `json:"email"`
	Password string    `json:"-"`
}

func New(name, email, password string) (*Model, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}
	user := &Model{
		ID:       entity.NewID(),
		Name:     name,
		Email:    email,
		Password: string(hash),
	}

	err = user.ValidateNewUser()
	if err != nil {
		return nil, err
	}

	return user, nil
}
