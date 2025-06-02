package user

import (
	"time"

	"github.com/sunnygosdk/go-chi-fullcycle-api/pkg/entity"
)

type Model struct {
	ID        entity.ID `json:"id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	Password  string    `json:"-"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type CreateUserOptions struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type UpdateUserOptions struct {
	Name     *string `json:"name,omitempty"`
	Email    *string `json:"email,omitempty"`
	Password *string `json:"password,omitempty"`
}

func ToCreate(options CreateUserOptions) (*Model, error) {
	hash, err := ValidateCreatePassword(options.Password)
	if err != nil {
		return nil, err
	}

	user := &Model{
		ID:        entity.NewID(),
		Name:      options.Name,
		Email:     options.Email,
		Password:  hash,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	err = user.ValidateCreateUser()
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (user *Model) ToUpdate(options UpdateUserOptions) (*Model, error) {
	if options.Password != nil {
		hash, err := ValidateUpdatePassword(user.Password, *options.Password)
		if err != nil {
			return nil, err
		}

		user.Password = hash
	}

	if options.Name != nil {
		err := ValidateUpdateName(user.Name, *options.Name)
		if err != nil {
			return nil, err
		}

		user.Name = *options.Name
	}

	if options.Email != nil {
		err := ValidateUpdateEmail(user.Email, *options.Email)
		if err != nil {
			return nil, err
		}

		user.Email = *options.Email
	}

	user.UpdatedAt = time.Now()
	return user, nil
}
