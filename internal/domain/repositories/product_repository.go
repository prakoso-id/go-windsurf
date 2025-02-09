package repositories

import "github.com/prakoso-id/go-windsurf/internal/domain/models"

type ProductRepository interface {
	Create(product *models.Product) error
	FindByID(id uint) (*models.Product, error)
	FindAll() ([]models.Product, error)
	Update(product *models.Product) error
	Delete(id uint) error
}
