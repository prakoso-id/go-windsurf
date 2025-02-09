package persistence

import (
	"database/sql"
	"errors"
	"time"

	"github.com/prakoso-id/go-windsurf/internal/domain/models"
	"github.com/prakoso-id/go-windsurf/internal/domain/repositories"
)

type productRepository struct {
	db *sql.DB
}

func NewProductRepository(db *sql.DB) repositories.ProductRepository {
	return &productRepository{db: db}
}

func (r *productRepository) Create(product *models.Product) error {
	query := `
		INSERT INTO products (name, description, price, created_at, updated_at)
		VALUES ($1, $2, $3, $4, $5)
		RETURNING id
	`
	return r.db.QueryRow(
		query,
		product.Name,
		product.Description,
		product.Price,
		time.Now(),
		time.Now(),
	).Scan(&product.ID)
}

func (r *productRepository) FindByID(id uint) (*models.Product, error) {
	var product models.Product
	query := `
		SELECT id, name, description, price, created_at, updated_at
		FROM products
		WHERE id = $1
	`
	err := r.db.QueryRow(query, id).Scan(
		&product.ID,
		&product.Name,
		&product.Description,
		&product.Price,
		&product.CreatedAt,
		&product.UpdatedAt,
	)
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	return &product, nil
}

func (r *productRepository) FindAll() ([]models.Product, error) {
	query := `
		SELECT id, name, description, price, created_at, updated_at
		FROM products
		ORDER BY created_at DESC
	`
	rows, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var products []models.Product
	for rows.Next() {
		var product models.Product
		err := rows.Scan(
			&product.ID,
			&product.Name,
			&product.Description,
			&product.Price,
			&product.CreatedAt,
			&product.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}
		products = append(products, product)
	}
	return products, nil
}

func (r *productRepository) Update(product *models.Product) error {
	query := `
		UPDATE products
		SET name = $1, description = $2, price = $3, updated_at = $4
		WHERE id = $5
	`
	result, err := r.db.Exec(
		query,
		product.Name,
		product.Description,
		product.Price,
		time.Now(),
		product.ID,
	)
	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if rowsAffected == 0 {
		return errors.New("product not found")
	}

	return nil
}

func (r *productRepository) Delete(id uint) error {
	query := `DELETE FROM products WHERE id = $1`
	result, err := r.db.Exec(query, id)
	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if rowsAffected == 0 {
		return errors.New("product not found")
	}

	return nil
}
