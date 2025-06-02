package configs

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
	_ "github.com/mattn/go-sqlite3"
	"github.com/sunnygosdk/go-chi-fullcycle-api/test/database"
)

func configureDB() (*sql.DB, error) {
	config := NewConfig()
	test, dsn := config.GetConnectionInfo()
	if test {
		return database.SetupTestDB(), nil
	}

	return sql.Open("mysql", dsn)
}

func ConnectDB() (*sql.DB, error) {
	db, err := configureDB()
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	log.Println("Database connected")
	return db, nil
}
