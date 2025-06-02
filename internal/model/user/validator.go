package user

import (
	"regexp"
	"unicode"

	"golang.org/x/crypto/bcrypt"
)

func (user *Model) ValidateNewUser() error {
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

	if user.Password == "" {
		return ErrPasswordRequired
	}

	if !isStrongPassword(user.Password) {
		return ErrWeakPassword
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

func isStrongPassword(password string) bool {
	if len(password) < 8 {
		return false
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

	return hasUpper && hasLower && hasNumber && hasSpecial
}
