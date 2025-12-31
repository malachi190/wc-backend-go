package service

import (
	"context"
	"database/sql"

	"github.com/jmoiron/sqlx"
	"github.com/malachi190/watchcircle/models"
)

type AuthService struct {
	db *sqlx.DB
}

func NewAuthService(db *sqlx.DB) *AuthService {
	return &AuthService{db: db}
}

// CreateUser registers a new user
func (s *AuthService) CreateUser(ctx context.Context, user *models.UserModel) error {
	query := `
        INSERT INTO users (email, display_name, username, password, fcm_token)
        VALUES (:email, :display_name, :username, :password, :fcm_token)
        RETURNING id, email, display_name, username, created_at, updated_at`

	rows, err := s.db.NamedQueryContext(ctx, query, user)

	if err != nil {
		return err
	}

	defer rows.Close()

	if rows.Next() {
		return rows.Scan(&user.Id, &user.Email, &user.DisplayName, &user.Username, &user.CreatedAt, &user.UpdatedAt)
	}

	return sql.ErrNoRows
}

func (s *AuthService) UpdateUserVerification(ctx context.Context, user *models.UserModel) error {
	query := `
		UPDATE users
		SET email_verified_at = :email_verified_at
		WHERE email = :email
	`

	res, err := s.db.NamedExecContext(ctx, query, user)

	if err != nil {
		return err
	}

	rows, err := res.RowsAffected()

	if err != nil {
		return err
	}

	if rows == 0 {
		return sql.ErrNoRows
	}

	return nil
}

func (s *AuthService) GetUser(ctx context.Context, user *models.UserModel) error {
	query := `
		SELECT * FROM users
		WHERE email = :email
	`
	rows, err := s.db.NamedQueryContext(ctx, query, user)

	if err != nil {
		return err
	}

	defer rows.Close()

	if !rows.Next() {
		return sql.ErrNoRows
	}

	return rows.StructScan(user)
}

func (s *AuthService) GetUserById(ctx context.Context, user *models.UserModel) error {
	query := `
		SELECT * FROM users
		WHERE id = :id
	`
	rows, err := s.db.NamedQueryContext(ctx, query, user)

	if err != nil {
		return err
	}

	defer rows.Close()

	if !rows.Next() {
		return sql.ErrNoRows
	}

	return rows.StructScan(user)
}
