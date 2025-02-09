package repositories

import (
	"github.com/prakoso-id/go-windsurf/internal/domain/models"
	"gorm.io/gorm"
)

type ProductRepository interface {
	CreateProduct(product *models.Product) error
	GetAllProducts() ([]models.Product, error)
	GetProductByID(id uint) (*models.Product, error)
	UpdateProduct(product *models.Product) error
	DeleteProduct(id uint) error
}

type productRepository struct {
	db *gorm.DB
}

func NewProductRepository(db *gorm.DB) ProductRepository {
	return &productRepository{db: db}
}

func (r *productRepository) CreateProduct(product *models.Product) error {
	return r.db.Create(product).Error
}

func (r *productRepository) GetAllProducts() ([]models.Product, error) {
	var products []models.Product
	err := r.db.Find(&products).Error
	return products, err
}

func (r *productRepository) GetProductByID(id uint) (*models.Product, error) {
	var product models.Product
	err := r.db.First(&product, id).Error
	if err != nil {
		return nil, err
	}
	return &product, nil
}

func (r *productRepository) UpdateProduct(product *models.Product) error {
	return r.db.Save(product).Error
}

func (r *productRepository) DeleteProduct(id uint) error {
	return r.db.Delete(&models.Product{}, id).Error
}
