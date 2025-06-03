package database

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

func SetupTestDB() *sql.DB {
	db, err := sql.Open("sqlite3", "file::memory:")
	if err != nil {
		panic(err)
	}
	createTestTables(db)
	return db
}

func createTestTables(db *sql.DB) {
	schema := `
	CREATE TABLE products (
		id TEXT PRIMARY KEY,
		name TEXT UNIQUE,
		price REAL,
		created_at DATETIME,
		updated_at DATETIME
	);

	CREATE TABLE users (
		id TEXT PRIMARY KEY,
		name TEXT,
		email TEXT UNIQUE,
		password TEXT,
		created_at DATETIME,
		updated_at DATETIME
	);
	`
	_, err := db.Exec(schema)
	if err != nil {
		panic(err)
	}
}
