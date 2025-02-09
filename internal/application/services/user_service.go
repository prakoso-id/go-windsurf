package services

import (
	"errors"
	"time"

	"github.com/google/uuid"
	"github.com/prakoso-id/go-windsurf/internal/domain/models"
	"github.com/prakoso-id/go-windsurf/internal/domain/repositories"
)

type UserService interface {
	Register(email, password, name string) (*models.User, error)
	Login(email, password string) (string, error) // Returns JWT token
	GetUserByID(id string) (*models.User, error)
	UpdateUser(id, email, name string) error
	UpdatePassword(id, newPassword string) error
	DeleteUser(id string) error
}

type userService struct {
	userRepo repositories.UserRepository
}

func NewUserService(userRepo repositories.UserRepository) UserService {
	return &userService{userRepo: userRepo}
}

func (s *userService) Register(email, password, name string) (*models.User, error) {
	existingUser, err := s.userRepo.FindByEmail(email)
	if err != nil {
		return nil, err
	}
	if existingUser != nil {
		return nil, errors.New("email already registered")
	}

	user := &models.User{
		ID:        uuid.New().String(),
		Email:     email,
		Password:  password,
		Name:      name,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	if err := s.userRepo.Create(user); err != nil {
		return nil, err
	}

	return user, nil
}

func (s *userService) Login(email, password string) (string, error) {
	user, err := s.userRepo.FindByEmailAndPassword(email, password)
	if err != nil {
		return "", err
	}
	if user == nil {
		return "", errors.New("invalid credentials")
	}

	return user.ID, nil
}

func (s *userService) GetUserByID(id string) (*models.User, error) {
	return s.userRepo.FindByID(id)
}

func (s *userService) UpdateUser(id, email, name string) error {
	user, err := s.userRepo.FindByID(id)
	if err != nil {
		return err
	}
	if user == nil {
		return errors.New("user not found")
	}

	user.Email = email
	user.Name = name
	user.UpdatedAt = time.Now()

	return s.userRepo.Update(user)
}

func (s *userService) UpdatePassword(id, newPassword string) error {
	user, err := s.userRepo.FindByID(id)
	if err != nil {
		return err
	}
	if user == nil {
		return errors.New("user not found")
	}

	return s.userRepo.UpdatePassword(id, newPassword)
}

func (s *userService) DeleteUser(id string) error {
	return s.userRepo.Delete(id)
}
