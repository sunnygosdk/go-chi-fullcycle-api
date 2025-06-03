package repository

import (
	"database/sql"

	"github.com/sunnygosdk/go-chi-fullcycle-api/internal/model"
	"github.com/sunnygosdk/go-chi-fullcycle-api/pkg/entity"
)

type UserRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{db: db}
}

func (r *UserRepository) GetUsers(page int, limit int) ([]model.UserModel, error) {
	query := "SELECT id, name, email, password, created_at, updated_at FROM users LIMIT ? OFFSET ?"
	rows, err := r.db.Query(query, limit, (page-1)*limit)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []model.UserModel
	for rows.Next() {
		var u model.UserModel
		if err := rows.Scan(&u.ID, &u.Name, &u.Email, &u.Password, &u.CreatedAt, &u.UpdatedAt); err != nil {
			return nil, err
		}
		users = append(users, u)
	}
	return users, nil
}

func (r *UserRepository) GetUserByID(id entity.ID) (*model.UserModel, error) {
	query := "SELECT id, name, email, password, created_at, updated_at FROM users WHERE id = ?"
	row := r.db.QueryRow(query, id.String())
	var u model.UserModel
	if err := row.Scan(&u.ID, &u.Name, &u.Email, &u.Password, &u.CreatedAt, &u.UpdatedAt); err != nil {
		return nil, err
	}
	return &u, nil
}

func (r *UserRepository) GetUserByEmail(email string) (*model.UserModel, error) {
	query := "SELECT id, name, email, password, created_at, updated_at FROM users WHERE email = ?"
	row := r.db.QueryRow(query, email)
	var u model.UserModel
	if err := row.Scan(&u.ID, &u.Name, &u.Email, &u.Password, &u.CreatedAt, &u.UpdatedAt); err != nil {
		return nil, err
	}
	return &u, nil
}

func (r *UserRepository) GetTotalUsers() (int, error) {
	query := "SELECT COUNT(*) FROM users"
	row := r.db.QueryRow(query)
	var count int
	if err := row.Scan(&count); err != nil {
		return 0, err
	}
	return count, nil
}

func (r *UserRepository) GetTotalPages(limit int) (int, error) {
	totalUsers, err := r.GetTotalUsers()
	if err != nil {
		return 0, err
	}
	return (totalUsers + limit - 1) / limit, nil
}

func (r *UserRepository) Create(user *model.UserModel) error {
	query := "INSERT INTO users (id, name, email, password, created_at, updated_at) VALUES (?, ?, ?, ?, ?, ?)"
	_, err := r.db.Exec(query, user.ID, user.Name, user.Email, user.Password, user.CreatedAt, user.UpdatedAt)
	if err != nil {
		return err
	}
	return nil
}

func (r *UserRepository) CreateBatch(users []model.UserModel) error {
	query := "INSERT INTO users (id, name, email, password, created_at, updated_at) VALUES (?, ?, ?, ?, ?, ?)"
	for _, user := range users {
		_, err := r.db.Exec(query, user.ID, user.Name, user.Email, user.Password, user.CreatedAt, user.UpdatedAt)
		if err != nil {
			return err
		}
	}
	return nil
}

func (r *UserRepository) Update(id entity.ID, user *model.UserModel) error {
	query := "UPDATE users SET name = ?, email = ?, password = ?, updated_at = ? WHERE id = ?"
	_, err := r.db.Exec(query, user.Name, user.Email, user.Password, user.UpdatedAt, id.String())
	if err != nil {
		return err
	}
	return nil
}

func (r *UserRepository) Delete(id entity.ID) error {
	query := "DELETE FROM users WHERE id = ?"
	_, err := r.db.Exec(query, id.String())
	if err != nil {
		return err
	}
	return nil
}
