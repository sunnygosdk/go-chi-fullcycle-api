package database

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
	_ "github.com/mattn/go-sqlite3"
	"github.com/sunnygosdk/go-chi-fullcycle-api/config"
)

func configureDB(config *config.Config) (*sql.DB, error) {
	test, dsn := config.GetConnectionInfo()
	if test {
		log.Println("Connecting to SQLite Test Database...")
		return SetupTestDB(), nil
	}

	log.Println("Connecting to MySQL Database...")
	return sql.Open("mysql", dsn)
}

func ConnectDB(config *config.Config) (*sql.DB, error) {
	db, err := configureDB(config)
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	log.Println("Database Connected")
	return db, nil
}
