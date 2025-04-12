package domain

type UserRepository interface {
	Create(user *User) error
	GetByID(id string) (*User, error)
	GetByEmail(email string) (*User, error)
}

type AgendaRepository interface {
	Create(agenda *Agenda) error
	GetByID(id string) (*Agenda, error)
	GetAll() ([]*Agenda, error)
}
