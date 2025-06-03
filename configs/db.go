package configs

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
	_ "github.com/mattn/go-sqlite3"
	"github.com/sunnygosdk/go-chi-fullcycle-api/test/database"
)

func (c *conf) configureDB() (*sql.DB, error) {
	test, dsn := c.GetConnectionInfo()
	if test {
		log.Println("Connecting to SQLite Test Database...")
		return database.SetupTestDB(), nil
	}

	log.Println("Connecting to MySQL Database...")
	return sql.Open("mysql", dsn)
}

func (c *conf) ConnectDB() (*sql.DB, error) {
	db, err := c.configureDB()
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
