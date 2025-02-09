package services

import (
	"github.com/prakoso-id/go-windsurf/internal/domain/models"
	"github.com/prakoso-id/go-windsurf/internal/domain/repositories"
)

type ProductService interface {
	CreateProduct(product *models.Product) error
	GetProduct(id uint) (*models.Product, error)
	ListProducts() ([]models.Product, error)
	UpdateProduct(product *models.Product) error
	DeleteProduct(id uint) error
}

type productService struct {
	repo repositories.ProductRepository
}

func NewProductService(repo repositories.ProductRepository) ProductService {
	return &productService{repo: repo}
}

func (s *productService) CreateProduct(product *models.Product) error {
	return s.repo.Create(product)
}

func (s *productService) GetProduct(id uint) (*models.Product, error) {
	return s.repo.FindByID(id)
}

func (s *productService) ListProducts() ([]models.Product, error) {
	return s.repo.FindAll()
}

func (s *productService) UpdateProduct(product *models.Product) error {
	return s.repo.Update(product)
}

func (s *productService) DeleteProduct(id uint) error {
	return s.repo.Delete(id)
}
