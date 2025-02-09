package postgres

import (
	"database/sql"
	"errors"

	"github.com/prakoso-id/go-windsurf/internal/domain/models"
	"github.com/prakoso-id/go-windsurf/internal/domain/repositories"
)

type userRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) repositories.UserRepository {
	return &userRepository{db: db}
}

func (r *userRepository) Create(user *models.User) error {
	query := `
		INSERT INTO users (id, email, password, name, created_at, updated_at)
		VALUES ($1, $2, crypt($3, gen_salt('bf')), $4, $5, $6)
	`
	_, err := r.db.Exec(query, user.ID, user.Email, user.Password, user.Name, user.CreatedAt, user.UpdatedAt)
	return err
}

func (r *userRepository) FindByID(id string) (*models.User, error) {
	var user models.User
	query := `
		SELECT id, email, name, created_at, updated_at
		FROM users
		WHERE id = $1
	`
	err := r.db.QueryRow(query, id).Scan(&user.ID, &user.Email, &user.Name, &user.CreatedAt, &user.UpdatedAt)
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *userRepository) FindByEmail(email string) (*models.User, error) {
	var user models.User
	query := `
		SELECT id, email, name, created_at, updated_at
		FROM users
		WHERE email = $1
	`
	err := r.db.QueryRow(query, email).Scan(&user.ID, &user.Email, &user.Name, &user.CreatedAt, &user.UpdatedAt)
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *userRepository) FindByEmailAndPassword(email, password string) (*models.User, error) {
	var user models.User
	query := `
		SELECT id, email, name, created_at, updated_at
		FROM users
		WHERE email = $1 AND password = crypt($2, password)
	`
	err := r.db.QueryRow(query, email, password).Scan(&user.ID, &user.Email, &user.Name, &user.CreatedAt, &user.UpdatedAt)
	if err == sql.ErrNoRows {
		return nil, errors.New("invalid credentials")
	}
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *userRepository) Update(user *models.User) error {
	query := `
		UPDATE users
		SET email = $1, name = $2, updated_at = $3
		WHERE id = $4
	`
	result, err := r.db.Exec(query, user.Email, user.Name, user.UpdatedAt, user.ID)
	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if rowsAffected == 0 {
		return errors.New("user not found")
	}

	return nil
}

func (r *userRepository) UpdatePassword(id string, newPassword string) error {
	query := `
		UPDATE users
		SET password = crypt($1, gen_salt('bf')), updated_at = NOW()
		WHERE id = $2
	`
	result, err := r.db.Exec(query, newPassword, id)
	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if rowsAffected == 0 {
		return errors.New("user not found")
	}

	return nil
}

func (r *userRepository) Delete(id string) error {
	query := `DELETE FROM users WHERE id = $1`
	result, err := r.db.Exec(query, id)
	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if rowsAffected == 0 {
		return errors.New("user not found")
	}

	return nil
}
