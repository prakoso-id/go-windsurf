package repositories

import "github.com/prakoso-id/go-windsurf/internal/domain/models"

type UserRepository interface {
	Create(user *models.User) error
	FindByID(id string) (*models.User, error)
	FindByEmail(email string) (*models.User, error)
	FindByEmailAndPassword(email, password string) (*models.User, error)
	Update(user *models.User) error
	UpdatePassword(id string, newPassword string) error
	Delete(id string) error
}
