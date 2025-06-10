package repository

import (
	"errors"
	"fmt"

	"github.com/go-sql-driver/mysql"
)

// Error messages for MySQL errors
var (
	ErrorMySQLDuplicateEntry          = errors.New("mysql duplicate entry error")
	ErrorMySQLFieldNotNull            = errors.New("mysql field not null error")
	ErrorMySQLUnknown                 = errors.New("mysql unknown error")
	ErrorMySQLInvalidType             = errors.New("mysql invalid type error")
	ErrorMySQLForeignKeyConstraint    = errors.New("mysql foreign key constraint error")
	ErrorMySQLForeignKeyMissing       = errors.New("mysql foreign key missing error")
	ErrorMySQLCheckConstraintViolated = errors.New("mysql check constraint violated error")
)

// MySQLError is a wrapper for MySQL errors
type MySQLError struct {
	*mysql.MySQLError
}

// newMySQLError creates a new MySQLError
//
// Parameters:
//   - mysqlErr: The MySQL error to wrap.
//
// Returns:
//   - MySQLError: The new MySQLError.
func newMySQLError(mysqlErr *mysql.MySQLError) *MySQLError {
	return &MySQLError{mysqlErr}
}

// MapMySQLError maps a MySQL error to a custom error
//
// Parameters:
//   - err: The error to map.
//
// Returns:
//   - error: The mapped error.
func MapMySQLError(err error) error {
	if err == nil {
		return nil
	}

	mysqlErr, isMySQLErr := err.(*mysql.MySQLError)
	if !isMySQLErr {
		return err
	}

	e := newMySQLError(mysqlErr)
	return e.getError()
}

// getError returns the custom error for the MySQL error
//
// Parameters:
//   - e: The MySQLError to get the error for.
//
// Returns:
//   - error: The custom error.
func (e *MySQLError) getError() error {
	switch e.Number {
	case 1062:
		return fmt.Errorf("%w: %v", ErrorMySQLDuplicateEntry, e.Message)
	case 1048:
		return fmt.Errorf("%w: %v", ErrorMySQLFieldNotNull, e.Message)
	case 1265:
		return fmt.Errorf("%w: %v", ErrorMySQLInvalidType, e.Message)
	case 1451:
		return fmt.Errorf("%w: %v", ErrorMySQLForeignKeyConstraint, e.Message)
	case 1452:
		return fmt.Errorf("%w: %v", ErrorMySQLForeignKeyMissing, e.Message)
	case 3819:
		return fmt.Errorf("%w: %v", ErrorMySQLCheckConstraintViolated, e.Message)
	default:
		return fmt.Errorf("%w: %v", ErrorMySQLUnknown, e.Message)
	}
}
