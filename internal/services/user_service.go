package services

import (
	"errors"
	"time"

	"github.com/google/uuid"
	"github.com/prakoso-id/go-windsurf/internal/application/services"
	"github.com/prakoso-id/go-windsurf/internal/domain/models"
	"github.com/prakoso-id/go-windsurf/internal/domain/repositories"
)

type userService struct {
	userRepo    repositories.UserRepository
	authService services.AuthService
}

func NewUserService(userRepo repositories.UserRepository, authService services.AuthService) services.UserService {
	return &userService{
		userRepo:    userRepo,
		authService: authService,
	}
}

func (s *userService) Register(email, password, name string) (*models.User, error) {
	// Check if user already exists
	existingUser, err := s.userRepo.FindByEmail(email)
	if err != nil {
		return nil, err
	}
	if existingUser != nil {
		return nil, errors.New("email already registered")
	}

	// Password will be hashed by the database using crypt() function
	user := &models.User{
		ID:        uuid.New().String(),
		Email:     email,
		Password:  password, // Raw password will be hashed by the database
		Name:      name,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	err = s.userRepo.Create(user)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (s *userService) Login(email, password string) (string, error) {
	// The password verification will be done by the database using crypt()
	user, err := s.userRepo.FindByEmailAndPassword(email, password)
	if err != nil {
		return "", err
	}
	if user == nil {
		return "", errors.New("invalid email or password")
	}

	// Generate JWT token
	token, err := s.authService.GenerateToken(user.ID)
	if err != nil {
		return "", err
	}

	return token, nil
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

	user.Password = newPassword // The repository will hash the password
	return s.userRepo.UpdatePassword(id, newPassword)
}

func (s *userService) DeleteUser(id string) error {
	return s.userRepo.Delete(id)
}
