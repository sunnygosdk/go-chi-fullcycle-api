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

type CreateUserDTO struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type UpdateUserDTO struct {
	Name     *string `json:"name,omitempty"`
	Email    *string `json:"email,omitempty"`
	Password *string `json:"password,omitempty"`
}

func ToCreate(userDTO CreateUserDTO) (*Model, error) {
	hash, err := ValidateCreatePassword(userDTO.Password)
	if err != nil {
		return nil, err
	}

	user := &Model{
		ID:        entity.NewID(),
		Name:      userDTO.Name,
		Email:     userDTO.Email,
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

func (user *Model) ToUpdate(userDTO UpdateUserDTO) (*Model, error) {
	if userDTO.Password != nil {
		hash, err := ValidateUpdatePassword(*userDTO.Password)
		if err != nil {
			return nil, err
		}

		user.Password = hash
	}

	if userDTO.Name != nil {
		err := ValidateUpdateName(*userDTO.Name)
		if err != nil {
			return nil, err
		}

		user.Name = *userDTO.Name
	}

	if userDTO.Email != nil {
		err := ValidateUpdateEmail(*userDTO.Email)
		if err != nil {
			return nil, err
		}

		user.Email = *userDTO.Email
	}

	user.UpdatedAt = time.Now()
	return user, nil
}
