package repository

import (
	"database/sql"

	"github.com/sunnygosdk/go-chi-fullcycle-api/internal/product/model"
	"github.com/sunnygosdk/go-chi-fullcycle-api/pkg/entity"
)

type ProductRepository struct {
	db *sql.DB
}

func NewProductRepository(db *sql.DB) *ProductRepository {
	return &ProductRepository{db: db}
}

func (r *ProductRepository) GetProducts(page int, limit int) ([]model.ProductModel, error) {
	query := "SELECT id, name, price, created_at FROM products LIMIT ? OFFSET ?"
	rows, err := r.db.Query(query, limit, (page-1)*limit)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var products []model.ProductModel
	for rows.Next() {
		var p model.ProductModel
		if err := rows.Scan(&p.ID, &p.Name, &p.Price, &p.CreatedAt); err != nil {
			return nil, err
		}
		products = append(products, p)
	}
	return products, nil
}

func (r *ProductRepository) GetProductByID(id entity.ID) (*model.ProductModel, error) {
	query := "SELECT id, name, price, created_at FROM products WHERE id = ?"
	row := r.db.QueryRow(query, id.String())
	var p model.ProductModel
	if err := row.Scan(&p.ID, &p.Name, &p.Price, &p.CreatedAt); err != nil {
		return nil, err
	}
	return &p, nil
}

func (r *ProductRepository) GetProductByName(name string) (*model.ProductModel, error) {
	query := "SELECT id, name, price, created_at FROM products WHERE name = ?"
	row := r.db.QueryRow(query, name)
	var p model.ProductModel
	if err := row.Scan(&p.ID, &p.Name, &p.Price, &p.CreatedAt); err != nil {
		return nil, err
	}
	return &p, nil
}

func (r *ProductRepository) GetTotalProducts() (int, error) {
	query := "SELECT COUNT(*) FROM products"
	row := r.db.QueryRow(query)
	var count int
	if err := row.Scan(&count); err != nil {
		return 0, err
	}
	return count, nil
}

func (r *ProductRepository) GetTotalProductPages(limit int) (int, error) {
	totalProducts, err := r.GetTotalProducts()
	if err != nil {
		return 0, err
	}
	return (totalProducts + limit - 1) / limit, nil
}

func (r *ProductRepository) Create(product *model.ProductModel) (*model.ProductModel, error) {
	query := "INSERT INTO products (id, name, price, created_at) VALUES (?, ?, ?, ?)"
	_, err := r.db.Exec(query, product.ID, product.Name, product.Price, product.CreatedAt)
	if err != nil {
		return nil, err
	}
	return product, nil
}

func (r *ProductRepository) CreateBatch(products []model.ProductModel) error {
	query := "INSERT INTO products (id, name, price, created_at) VALUES (?, ?, ?, ?)"
	for _, product := range products {
		_, err := r.db.Exec(query, product.ID, product.Name, product.Price, product.CreatedAt)
		if err != nil {
			return err
		}
	}
	return nil
}

func (r *ProductRepository) Update(id entity.ID, product *model.ProductModel) error {
	query := "UPDATE products SET name = ?, price = ?, updated_at = ? WHERE id = ?"
	_, err := r.db.Exec(query, product.Name, product.Price, product.UpdatedAt, id.String())
	if err != nil {
		return err
	}
	return nil
}

func (r *ProductRepository) Delete(id entity.ID) error {
	query := "DELETE FROM products WHERE id = ?"
	_, err := r.db.Exec(query, id.String())
	if err != nil {
		return err
	}
	return nil
}
