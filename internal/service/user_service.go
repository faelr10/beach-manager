package service

import (
	"fmt"
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

	fmt.Println("User to be created:", user)

	existing, _ := s.repository.GetByEmail(user.Email)
	if existing != nil {
		return nil, domain.ErrUserAlreadyExists
	}

	err := s.repository.Create(user)
	if err != nil {
		return nil, err
	}

	output := dto.FromUser(user)

	return &output, nil
}

func (s *UserService) GetById(id string) (*dto.UserOutput, error) {
	user, err := s.repository.GetByID(id)
	if err != nil {
		return nil, err
	}

	output := dto.FromUser(user)
	return &output, nil
}
