package user

import (
	"regexp"
	"unicode"

	"golang.org/x/crypto/bcrypt"
)

func ValidateCreatePassword(password string) (string, error) {
	err := isStrongPassword(password)
	if err != nil {
		return "", ErrWeakPassword
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hash), nil
}

func ValidateUpdatePassword(userPassword, newPassword string) (string, error) {
	err := isStrongPassword(newPassword)
	if err != nil {
		return "", ErrWeakPassword
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(newPassword), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}

	if userPassword == string(hash) {
		return "", ErrSamePassword
	}

	return string(hash), nil
}

func ValidateUpdateEmail(userEmail, newEmail string) error {
	if userEmail == newEmail {
		return ErrSameEmail
	}

	if !isValidEmail(newEmail) {
		return ErrInvalidEmail
	}

	return nil
}

func ValidateUpdateName(userName, newName string) error {
	if userName == newName {
		return ErrSameName
	}

	if newName == "" {
		return ErrNameRequired
	}

	return nil
}

func (user *Model) ValidateCreateUser() error {
	if user.ID.String() == "" {
		return ErrInvalidID
	}

	if user.Name == "" {
		return ErrNameRequired
	}

	if user.Email == "" {
		return ErrEmailRequired
	}

	if !isValidEmail(user.Email) {
		return ErrInvalidEmail
	}

	return nil
}

func (user *Model) ValidatePassword(password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	return err == nil
}

func isValidEmail(email string) bool {
	emailRegex := regexp.MustCompile(`^[a-zA-Z0-9._%+\-]+@[a-zA-Z0-9.\-]+\.[a-zA-Z]{2,}$`)
	return emailRegex.MatchString(email)
}

func isStrongPassword(password string) error {
	if password == "" {
		return ErrPasswordRequired
	}
	if len(password) < 8 {
		return ErrWeakPassword
	}

	var hasUpper, hasLower, hasNumber, hasSpecial bool

	for _, ch := range password {
		switch {
		case unicode.IsUpper(ch):
			hasUpper = true
		case unicode.IsLower(ch):
			hasLower = true
		case unicode.IsNumber(ch):
			hasNumber = true
		case unicode.IsPunct(ch) || unicode.IsSymbol(ch):
			hasSpecial = true
		}
	}

	if !hasUpper || !hasLower || !hasNumber || !hasSpecial {
		return ErrWeakPassword
	}

	return nil
}
