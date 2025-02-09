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
	productRepo repositories.ProductRepository
}

func NewProductService(productRepo repositories.ProductRepository) ProductService {
	return &productService{productRepo: productRepo}
}

func (s *productService) CreateProduct(product *models.Product) error {
	return s.productRepo.Create(product)
}

func (s *productService) GetProduct(id uint) (*models.Product, error) {
	return s.productRepo.FindByID(id)
}

func (s *productService) ListProducts() ([]models.Product, error) {
	return s.productRepo.FindAll()
}

func (s *productService) UpdateProduct(product *models.Product) error {
	return s.productRepo.Update(product)
}

func (s *productService) DeleteProduct(id uint) error {
	return s.productRepo.Delete(id)
}
