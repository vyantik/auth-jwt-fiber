package user

import (
	"golang.org/x/crypto/bcrypt"
)

type Service struct {
	repository *Repository
}

func NewService(repository *Repository) *Service {
	return &Service{
		repository: repository,
	}
}

func (s *Service) Create(user User) (*User, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}
	user.Password = string(hashedPassword)
	createdUser, err := s.repository.Create(user)
	if err != nil {
		return nil, err
	}
	return createdUser, nil
}

func (s *Service) GetByEmail(email string) (*User, error) {
	user, err := s.repository.GetByEmail(email)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (s *Service) GetByID(id uint) (*User, error) {
	user, err := s.repository.GetByID(id)
	if err != nil {
		return nil, err
	}
	return user, nil
}
