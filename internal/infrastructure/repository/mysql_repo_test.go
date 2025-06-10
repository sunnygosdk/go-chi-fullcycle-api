package repository_test

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"os"
	"testing"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/stretchr/testify/assert"
	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/wait"
)

// Global variables for the database connection and teardown function
var db *sql.DB
var teardown func()

// TestMain is the entry point for the test suite
func TestMain(m *testing.M) {
	var err error
	t := &testing.T{}
	db, teardown = setupMySQLContainer(t)
	defer teardown()

	err = runMigrations(db, t)
	if err != nil {
		log.Fatalf("failed to run migrations: %v", err)
	}

	code := m.Run()
	os.Exit(code)
}

// setupMySQLContainer sets up a MySQL container for testing
func setupMySQLContainer(t *testing.T) (*sql.DB, func()) {
	ctx := context.Background()

	req := testcontainers.ContainerRequest{
		Image:        "mysql:8.0.23",
		ExposedPorts: []string{"3306/tcp"},
		Env: map[string]string{
			"MYSQL_ROOT_PASSWORD": "root",
			"MYSQL_DATABASE":      "test",
			"MYSQL_USER":          "test",
			"MYSQL_PASSWORD":      "test",
		},
		WaitingFor: wait.ForLog("port 3306 MySQL Community Server - GPL"),
	}

	mysqlContainer, err := testcontainers.GenericContainer(ctx, testcontainers.GenericContainerRequest{
		ContainerRequest: req,
		Started:          true,
	})
	assert.NoError(t, err)

	host, err := mysqlContainer.Host(ctx)
	assert.NoError(t, err)

	port, err := mysqlContainer.MappedPort(ctx, "3306")
	assert.NoError(t, err)

	user := "test"
	password := "test"
	database := "test"

	dns := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true", user, password, host, port.Port(), database)
	var db *sql.DB

	for i := 0; i < 10; i++ {
		db, err = sql.Open("mysql", dns)
		if err == nil && db.Ping() == nil {
			break
		}
		time.Sleep(1 * time.Second)
	}
	assert.NoError(t, err)

	teardown := func() {
		db.Close()
		_ = mysqlContainer.Terminate(ctx)
	}

	return db, teardown
}

// runMigrations runs the database migrations
func runMigrations(db *sql.DB, t *testing.T) error {
	statements := []string{
		`CREATE TABLE departments (
			id CHAR(36) PRIMARY KEY,
			name VARCHAR(255),
			description TEXT,
			created_at DATETIME,
			updated_at DATETIME,
			deleted_at DATETIME NULL
		);`,
		`CREATE TABLE products (
			id CHAR(36) PRIMARY KEY,
			name VARCHAR(255),
			description TEXT,
			price DOUBLE,
			department_id CHAR(36),
			created_at DATETIME,
			updated_at DATETIME,
			deleted_at DATETIME NULL,
			FOREIGN KEY (department_id) REFERENCES departments(id) ON DELETE CASCADE
		);`,
		`CREATE TABLE stores (
			id CHAR(36) PRIMARY KEY,
			name VARCHAR(255),
			address TEXT,
			created_at DATETIME,
			updated_at DATETIME,
			deleted_at DATETIME NULL
		);`,
		`CREATE TABLE stock (
			id CHAR(36) PRIMARY KEY,
			quantity INT,
			product_id CHAR(36),
			store_id CHAR(36),
			created_at DATETIME,
			updated_at DATETIME,
			deleted_at DATETIME NULL,
			FOREIGN KEY (product_id) REFERENCES products(id) ON DELETE CASCADE,
			FOREIGN KEY (store_id) REFERENCES stores(id) ON DELETE CASCADE
		);`,
		`CREATE TABLE store_departments (
			id CHAR(36) PRIMARY KEY,
			store_id CHAR(36),
			department_id CHAR(36),
			created_at DATETIME,
			updated_at DATETIME,
			deleted_at DATETIME NULL,
			FOREIGN KEY (store_id) REFERENCES stores(id) ON DELETE CASCADE,
			FOREIGN KEY (department_id) REFERENCES departments(id) ON DELETE CASCADE
		);`,
		`CREATE TABLE transactions (
			id CHAR(36) PRIMARY KEY,
			quantity INT,
			transaction_type VARCHAR(50),
			product_id CHAR(36),
			stock_id CHAR(36),
			created_at DATETIME,
			updated_at DATETIME,
			deleted_at DATETIME NULL,
			FOREIGN KEY (product_id) REFERENCES products(id) ON DELETE CASCADE,
			FOREIGN KEY (stock_id) REFERENCES stock(id) ON DELETE CASCADE
		);`,
	}

	for i, stmt := range statements {
		_, err := db.Exec(stmt)
		assert.NoError(t, err)
		if err != nil {
			return fmt.Errorf("error on migration %d: %w", i+1, err)
		}
	}
	return nil
}
