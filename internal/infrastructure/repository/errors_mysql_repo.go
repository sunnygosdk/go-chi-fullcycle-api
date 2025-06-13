package repository

import (
	"fmt"

	"github.com/go-sql-driver/mysql"
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
		return fmt.Errorf("%w: %v", ErrorRepoDuplicateEntry, e.Message)
	case 1048:
		return fmt.Errorf("%w: %v", ErrorRepoFieldNotNull, e.Message)
	case 1265:
		return fmt.Errorf("%w: %v", ErrorRepoInvalidType, e.Message)
	case 1451:
		return fmt.Errorf("%w: %v", ErrorRepoForeignKeyConstraint, e.Message)
	case 1452:
		return fmt.Errorf("%w: %v", ErrorRepoForeignKeyMissing, e.Message)
	case 3819:
		return fmt.Errorf("%w: %v", ErrorRepoCheckConstraintViolated, e.Message)
	default:
		return fmt.Errorf("%w: %v", ErrorRepoUnknown, e.Message)
	}
}
