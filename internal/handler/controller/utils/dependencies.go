package utils

import "database/sql"

func InjectController(db *sql.DB) *Controller {
	return NewController(db)
}
