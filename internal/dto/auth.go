package dto

import "go/beach-manager/internal/domain"

type LoginInput struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type AuthOutput struct {
	Token  string `json:"token"`
	UserID string `json:"user_id"`
	Email  string `json:"email"`
}

func FromAuth(auth *domain.Auth) AuthOutput {
	return AuthOutput{
		Token:  auth.Token,
		UserID: auth.UserID,
		Email:  auth.Email,
	}
}
