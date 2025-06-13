package repository

import "errors"

// Error messages for MySQL errors
var (
	ErrorRepoDuplicateEntry          = errors.New("duplicate entry error")
	ErrorRepoFieldNotNull            = errors.New("field not null error")
	ErrorRepoUnknown                 = errors.New("unknown error")
	ErrorRepoInvalidType             = errors.New("invalid type error")
	ErrorRepoForeignKeyConstraint    = errors.New("foreign key constraint error")
	ErrorRepoForeignKeyMissing       = errors.New("foreign key missing error")
	ErrorRepoCheckConstraintViolated = errors.New("check constraint violated error")
)
