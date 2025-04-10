package repository

import (
	"database/sql"
	"go/beach-manager/internal/domain"
)

// como se fosse a classe de implementacao do repository
type UserRepository struct {
	db *sql.DB
}

// como se fosse o construtor
func NewUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{db: db}
}

// como se fosse o metodo de criar um usuario
func (r *UserRepository) Create(user *domain.User) error {
	query := "INSERT INTO users (name, email, password) VALUES ($1, $2, $3)"
	_, err := r.db.Exec(query, user.Name, user.Email, user.Password)
	if err != nil {
		return err
	}
	return nil
}

func (r *UserRepository) GetByID(id string) (*domain.User, error) {
	query := "SELECT id, name, email, password FROM users WHERE id = $1"
	row := r.db.QueryRow(query, id)

	var user domain.User
	err := row.Scan(&user.ID, &user.Name, &user.Email, &user.Password)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}

	return &user, nil
}

func (r *UserRepository) GetByEmail(email string) (*domain.User, error) {
	query := "SELECT id, name, email, password FROM users WHERE email = $1"
	row := r.db.QueryRow(query, email)
	
	var user domain.User
	err := row.Scan(&user.ID, &user.Name, &user.Email, &user.Password)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}
	return &user, nil
}
