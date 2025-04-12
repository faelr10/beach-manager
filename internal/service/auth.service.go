package service

import (
	"go/beach-manager/internal/domain"
	"go/beach-manager/internal/dto"
	"go/beach-manager/internal/provider"
)

type AuthService struct {
	userRepository domain.UserRepository
	jwt            provider.JWTProvider
}

func NewAuthService(userRepository domain.UserRepository, jwt provider.JWTProvider) *AuthService {
	return &AuthService{
		userRepository: userRepository,
		jwt:            jwt,
	}
}

func (s *AuthService) Login(input dto.LoginInput) (*dto.AuthOutput, error) {
	user, err := s.userRepository.GetByEmail(input.Email)
	if err != nil {
		return nil, err
	}

	if user.Password != input.Password {
		return nil, domain.ErrAuthInvalidCredentials
	}

	token, err := s.jwt.GenerateToken(user.ID)
	if err != nil {
		return nil, err
	}

	auth := domain.NewAuth(
		token,
		user.ID,
		user.Email,
	)

	output := dto.FromAuth(auth)

	return &output, nil

}
