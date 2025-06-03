package product

import (
	"database/sql"

	"github.com/sunnygosdk/go-chi-fullcycle-api/internal/model/product"
	"github.com/sunnygosdk/go-chi-fullcycle-api/pkg/entity"
)

type Repository struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) *Repository {
	return &Repository{db: db}
}

func (r *Repository) GetProducts() ([]product.Model, error) {
	query := "SELECT id, name, price, created_at FROM products"
	rows, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var products []product.Model
	for rows.Next() {
		var p product.Model
		if err := rows.Scan(&p.ID, &p.Name, &p.Price, &p.CreatedAt); err != nil {
			return nil, err
		}
		products = append(products, p)
	}
	return products, nil
}

func (r *Repository) GetProductByID(id entity.ID) (*product.Model, error) {
	query := "SELECT id, name, price, created_at FROM products WHERE id = ?"
	row := r.db.QueryRow(query, id.String())
	var p product.Model
	if err := row.Scan(&p.ID, &p.Name, &p.Price, &p.CreatedAt); err != nil {
		return nil, err
	}
	return &p, nil
}

func (r *Repository) GetProductByName(name string) (*product.Model, error) {
	query := "SELECT id, name, price, created_at FROM products WHERE name = ?"
	row := r.db.QueryRow(query, name)
	var p product.Model
	if err := row.Scan(&p.ID, &p.Name, &p.Price, &p.CreatedAt); err != nil {
		return nil, err
	}
	return &p, nil
}

func (r *Repository) Create(product *product.Model) error {
	query := "INSERT INTO products (id, name, price, created_at) VALUES (?, ?, ?, ?)"
	_, err := r.db.Exec(query, product.ID, product.Name, product.Price, product.CreatedAt)
	if err != nil {
		return err
	}
	return nil
}

func (r *Repository) Update(id entity.ID, product *product.Model) error {
	query := "UPDATE products SET name = ?, price = ?, updated_at = ? WHERE id = ?"
	_, err := r.db.Exec(query, product.Name, product.Price, product.UpdatedAt, id.String())
	if err != nil {
		return err
	}
	return nil
}

func (r *Repository) Delete(id entity.ID) error {
	query := "DELETE FROM products WHERE id = ?"
	_, err := r.db.Exec(query, id.String())
	if err != nil {
		return err
	}
	return nil
}
