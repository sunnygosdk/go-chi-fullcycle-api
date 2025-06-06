package employee

import (
	"regexp"
	"time"
	"unicode"

	"github.com/sunnygosdk/go-chi-fullcycle-api/pkg/entity"
	"golang.org/x/crypto/bcrypt"
)

// NewEmployee creates a new employee instance with the provided ID, username, first name, last name, email, and password.
// It initializes the employee with the given ID, username, first name, last name, email, and password,
// and sets the creation and update timestamps to the current time.
//
// The function also validates the employee before returning it. If validation fails,
// it returns an error.
//
// Parameters:
//   - Username: Username of the employee.
//   - FirstName: First name of the employee.
//   - LastName: Last name of the employee.
//   - Email: Email address of the employee.
//   - Password: Password of the employee.
//
// Returns:
//   - *employee: A pointer to the newly created and validated employee.
//   - error: An error if the employee validation fails.
func NewEmployee(Username string, FirstName string, LastName string, Email string, Password string) (*employee, error) {
	password, err := validateCreatePassword(Password)
	if err != nil {
		return nil, err
	}

	if !isValidEmail(Email) {
		return nil, ErrEmployeeInvalidEmail
	}

	employee := &employee{
		ID:        entity.NewID(),
		Username:  Username,
		FirstName: FirstName,
		LastName:  LastName,
		Email:     Email,
		Password:  password,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	employee, err = employee.validate()
	if err != nil {
		return nil, err
	}

	return employee, nil
}

// validate validates the employee instance.
// It checks if the employee ID is required and valid,
// and if the employee username, first name, last name, email, and password are required.
//
// Parameters:
//   - e: The employee instance to validate.
//
// Returns:
//   - *employee: A pointer to the validated employee instance.
//   - error: An error if the employee validation fails.
func (e *employee) validate() (*employee, error) {
	if e.ID.String() == "" {
		return nil, ErrEmployeeIDisRequired
	}

	_, err := entity.ParseID(e.ID.String())
	if err != nil {
		return nil, ErrEmployeeInvalidID
	}

	if e.Username == "" {
		return nil, ErrEmployeeUsernameRequired
	}

	if e.FirstName == "" {
		return nil, ErrEmployeeFirstNameRequired
	}

	if e.LastName == "" {
		return nil, ErrEmployeeLastNameRequired
	}

	if e.Email == "" {
		return nil, ErrEmployeeEmailRequired
	}

	if e.Password == "" {
		return nil, ErrEmployeePasswordRequired
	}

	return e, nil
}

// validateCreatePassword validates the employee password.
// It checks if the employee password is required and strong.
//
// Parameters:
//   - password: The employee password to validate.
//
// Returns:
//   - string: The validated employee password.
//   - error: An error if the employee password validation fails.
func validateCreatePassword(password string) (string, error) {
	err := isStrongPassword(password)
	if err != nil {
		return "", ErrEmployeeWeakPassword
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hash), nil
}

// isStrongPassword checks if the employee password is strong.
// It checks if the employee password is required and strong.
//
// Parameters:
//   - password: The employee password to validate.
//
// Returns:
//   - error: An error if the employee password validation fails.
func isStrongPassword(password string) error {
	if password == "" {
		return ErrEmployeePasswordRequired
	}
	if len(password) < 8 {
		return ErrEmployeeWeakPassword
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
		return ErrEmployeeWeakPassword
	}

	return nil
}

// ValidatePassword validates the employee password.
// It checks if the employee password is required and strong.
//
// Parameters:
//   - password: The employee password to validate.
//
// Returns:
//   - bool: True if the employee password is valid, false otherwise.
func (e *employee) ValidatePassword(password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(e.Password), []byte(password))
	return err == nil
}

// isValidEmail checks if the employee email is valid.
// It checks if the employee email is required and valid.
//
// Parameters:
//   - email: The employee email to validate.
//
// Returns:
//   - bool: True if the employee email is valid, false otherwise.
func isValidEmail(email string) bool {
	emailRegex := regexp.MustCompile(`^[a-zA-Z0-9._%+\-]+@[a-zA-Z0-9.\-]+\.[a-zA-Z]{2,}$`)
	return emailRegex.MatchString(email)
}
