package dto

import "go/beach-manager/internal/domain"

type CreateUserInput struct {
	Name      string `json:"name"`
	LocalName string `json:"local_name"`
	Email     string `json:"email"`
	Password  string `json:"password"`
}

type UserOutput struct {
	ID        string `json:"id"`
	Name      string `json:"name"`
	LocalName string `json:"local_name"`
	Email     string `json:"email"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

func ToUser(input CreateUserInput) *domain.User {
	return domain.NewUser(input.Name, input.LocalName, input.Email, input.Password)
}

func FromUser(user *domain.User) UserOutput {
	return UserOutput{
		ID:        user.ID,
		Name:      user.Name,
		LocalName: user.LocalName,
		Email:     user.Email,
		CreatedAt: user.CreatedAt.Format("2006-01-02T15:04:05Z07:00"),
		UpdatedAt: user.UpdatedAt.Format("2006-01-02T15:04:05Z07:00"),
	}
}
