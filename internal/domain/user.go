package domain

import (
	"github.com/google/uuid"
	"time"
)

type User struct {
	ID        string
	Name      string
	LocalName string
	Email     string
	Password  string
	CreatedAt time.Time
	UpdatedAt time.Time
}

func NewUser(name, localName, email, password string) *User {
	user := &User{
		ID:        uuid.New().String(),
		LocalName: localName,
		Name:      name,
		Email:     email,
		Password:  password,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	return user
}
