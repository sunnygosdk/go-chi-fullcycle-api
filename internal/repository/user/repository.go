package user

import (
	"database/sql"

	"github.com/sunnygosdk/go-chi-fullcycle-api/internal/model/user"
	"github.com/sunnygosdk/go-chi-fullcycle-api/pkg/entity"
)

type Repository struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) *Repository {
	return &Repository{db: db}
}

func (r *Repository) GetUsers() ([]user.Model, error) {
	query := "SELECT id, name, email, password, created_at, updated_at FROM users"
	rows, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []user.Model
	for rows.Next() {
		var u user.Model
		if err := rows.Scan(&u.ID, &u.Name, &u.Email, &u.Password, &u.CreatedAt, &u.UpdatedAt); err != nil {
			return nil, err
		}
		users = append(users, u)
	}
	return users, nil
}

func (r *Repository) GetUserByID(id string) (*user.Model, error) {
	query := "SELECT id, name, email, password, created_at, updated_at FROM users WHERE id = ?"
	row := r.db.QueryRow(query, id)
	var u user.Model
	if err := row.Scan(&u.ID, &u.Name, &u.Email, &u.Password, &u.CreatedAt, &u.UpdatedAt); err != nil {
		return nil, err
	}
	return &u, nil
}

func (r *Repository) GetUserByEmail(email string) (*user.Model, error) {
	query := "SELECT id, name, email, password, created_at, updated_at FROM users WHERE email = ?"
	row := r.db.QueryRow(query, email)
	var u user.Model
	if err := row.Scan(&u.ID, &u.Name, &u.Email, &u.Password, &u.CreatedAt, &u.UpdatedAt); err != nil {
		return nil, err
	}
	return &u, nil
}

func (r *Repository) Create(user *user.Model) error {
	query := "INSERT INTO users (id, name, email, password, created_at, updated_at) VALUES (?, ?, ?, ?, ?, ?)"
	_, err := r.db.Exec(query, user.ID, user.Name, user.Email, user.Password, user.CreatedAt, user.UpdatedAt)
	if err != nil {
		return err
	}
	return nil
}

func (r *Repository) Update(id entity.ID, user *user.Model) error {
	query := "UPDATE users SET name = ?, email = ?, password = ?, updated_at = ? WHERE id = ?"
	_, err := r.db.Exec(query, user.Name, user.Email, user.Password, user.UpdatedAt, id)
	if err != nil {
		return err
	}
	return nil
}

func (r *Repository) Delete(id entity.ID) error {
	query := "DELETE FROM users WHERE id = ?"
	_, err := r.db.Exec(query, id)
	if err != nil {
		return err
	}
	return nil
}
