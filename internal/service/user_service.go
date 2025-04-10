package service

import (
	"go/beach-manager/internal/domain"
	"go/beach-manager/internal/dto"
)

type UserService struct {
	repository domain.UserRepository
}

func NewUserService(repository domain.UserRepository) *UserService {
	return &UserService{repository: repository}
}

func (s *UserService) CreateUser(input dto.CreateUserInput) (*dto.UserOutput, error) {
	user := dto.ToUser(input)

	_, err := s.repository.GetByEmail(user.Email)
	if err != nil {
		return nil, err
	}

	err = s.repository.Create(user)
	if err != nil {
		return nil, err
	}

	output := dto.FromUser(user)

	return &output, nil
}
