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
	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/wait"
)

// Global variables for the database connection and teardown function
var db *sql.DB
var teardown func()

// TestMain is the entry point for the test suite
func TestMain(m *testing.M) {
	var err error
	db, teardown, err = setupMySQLContainer()
	if err != nil {
		log.Fatalf("failed to setup container: %v", err)
	}
	defer teardown()

	err = runMigrations()
	if err != nil {
		log.Fatalf("failed to run migrations: %v", err)
	}

	code := m.Run()
	os.Exit(code)
}

// setupMySQLContainer sets up a MySQL container for testing
func setupMySQLContainer() (*sql.DB, func(), error) {
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
		WaitingFor: wait.ForLog("ready for connections").WithStartupTimeout(60 * time.Second),
	}

	mysqlContainer, err := testcontainers.GenericContainer(ctx, testcontainers.GenericContainerRequest{
		ContainerRequest: req,
		Started:          true,
	})
	if err != nil {
		return nil, nil, fmt.Errorf("failed to setup container: %w", err)
	}

	host, err := mysqlContainer.Host(ctx)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to get container host: %w", err)
	}

	port, err := mysqlContainer.MappedPort(ctx, "3306")
	if err != nil {
		return nil, nil, fmt.Errorf("failed to get container port: %w", err)
	}

	user := "test"
	password := "test"
	database := "test"

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true", user, password, host, port.Port(), database)

	var db *sql.DB
	var pingErr error
	for i := 0; i < 10; i++ {
		db, err = sql.Open("mysql", dsn)
		if err == nil {
			pingErr = db.Ping()
			if pingErr == nil {
				break
			}
		}
		time.Sleep(1 * time.Second)
	}
	if pingErr != nil {
		return nil, nil, fmt.Errorf("failed to connect to database: %w", pingErr)
	}

	teardown := func() {
		db.Close()
		_ = mysqlContainer.Terminate(ctx)
	}

	return db, teardown, nil
}

// runMigrations runs the database migrations
func runMigrations() error {
	statements := []string{
		`CREATE TABLE departments (
			id CHAR(36) PRIMARY KEY,
			name VARCHAR(255) UNIQUE NOT NULL,
			description TEXT NOT NULL,
			created_at DATETIME,
			updated_at DATETIME,
			deleted_at DATETIME NULL,
			CONSTRAINT chk_department_name_min_length CHECK (CHAR_LENGTH(name) >= 2),
			CONSTRAINT chk_department_description_min_length CHECK (CHAR_LENGTH(description) >= 2)
		);`,
		`CREATE TABLE products (
			id CHAR(36) PRIMARY KEY,
			name VARCHAR(255) UNIQUE NOT NULL,
			description TEXT NOT NULL,
			price DOUBLE NOT NULL,
			department_id CHAR(36),
			created_at DATETIME,
			updated_at DATETIME,
			deleted_at DATETIME NULL,
			FOREIGN KEY (department_id) REFERENCES departments(id) ON DELETE CASCADE,
			CONSTRAINT chk_product_name_is_not_empty CHECK (CHAR_LENGTH(name) > 0),
			CONSTRAINT chk_product_description_min_length CHECK (CHAR_LENGTH(description) > 2),
			CONSTRAINT chk_product_price_is_non_negative CHECK (price >= 0)
		);`,
		`CREATE TABLE stores (
			id CHAR(36) PRIMARY KEY,
			name VARCHAR(255) UNIQUE NOT NULL,
			address TEXT NOT NULL,
			created_at DATETIME,
			updated_at DATETIME,
			deleted_at DATETIME NULL,
			CONSTRAINT chk_store_name_min_length CHECK (CHAR_LENGTH(name) >= 2),
			CONSTRAINT chk_store_address_min_length CHECK (CHAR_LENGTH(address) >= 2)
		);`,
		`CREATE TABLE stocks (
			id CHAR(36) PRIMARY KEY,
			quantity INT NOT NULL,
			product_id CHAR(36),
			store_id CHAR(36),
			created_at DATETIME,
			updated_at DATETIME,
			deleted_at DATETIME NULL,
			FOREIGN KEY (product_id) REFERENCES products(id) ON DELETE CASCADE,
			FOREIGN KEY (store_id) REFERENCES stores(id) ON DELETE CASCADE,
			CONSTRAINT chk_stock_quantity_is_non_negative CHECK (quantity >= 0),
			CONSTRAINT uq_stock_product_store UNIQUE (product_id, store_id)
		);`,
		`CREATE TABLE store_department_map (
			id CHAR(36) PRIMARY KEY,
			store_id CHAR(36),
			department_id CHAR(36),
			created_at DATETIME,
			updated_at DATETIME,
			deleted_at DATETIME NULL,
			FOREIGN KEY (store_id) REFERENCES stores(id) ON DELETE CASCADE,
			FOREIGN KEY (department_id) REFERENCES departments(id) ON DELETE CASCADE,
			CONSTRAINT uq_store_department UNIQUE (store_id, department_id)
		);`,
		`CREATE TABLE transactions (
			id CHAR(36) PRIMARY KEY,
			quantity INT NOT NULL,
			transaction_type VARCHAR(50) NOT NULL,
			product_id CHAR(36),
			stock_id CHAR(36),
			created_at DATETIME,
			updated_at DATETIME,
			deleted_at DATETIME NULL,
			FOREIGN KEY (product_id) REFERENCES products(id) ON DELETE CASCADE,
			FOREIGN KEY (stock_id) REFERENCES stocks(id) ON DELETE CASCADE,
			CONSTRAINT chk_quantity_not_zero CHECK (quantity <> 0),
			CONSTRAINT chk_transaction_type CHECK (transaction_type IN ('IN', 'OUT'))
		);`,
	}

	for i, stmt := range statements {
		_, err := db.Exec(stmt)
		if err != nil {
			return fmt.Errorf("error on migration %d: %w", i+1, err)
		}
	}
	return nil
}

func truncateTables(db *sql.DB) error {
	statements := []string{
		`SET FOREIGN_KEY_CHECKS = 0;`,
		`TRUNCATE TABLE departments;`,
		`TRUNCATE TABLE products;`,
		`TRUNCATE TABLE stores;`,
		`TRUNCATE TABLE stocks;`,
		`TRUNCATE TABLE store_department_map;`,
		`TRUNCATE TABLE transactions;`,
		`SET FOREIGN_KEY_CHECKS = 1;`,
	}

	for i, stmt := range statements {
		_, err := db.Exec(stmt)
		if err != nil {
			return fmt.Errorf("error on truncate %d: %w", i+1, err)
		}
	}
	return nil
}
