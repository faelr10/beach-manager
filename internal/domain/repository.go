package domain

type UserRepository interface {
	Create(user *User) error
	GetByID(id string) (*User, error)
	GetByEmail(email string) (*User, error)
}
