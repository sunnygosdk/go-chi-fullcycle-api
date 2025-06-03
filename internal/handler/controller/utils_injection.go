package controller

import "database/sql"

func InjectUtilsController(db *sql.DB) *UtilsController {
	return NewUtilsController(db)
}
