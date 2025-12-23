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
        INSERT INTO users (email, username, password, avatar, bio, fcm_token)
        VALUES (:email, :username, :password, :avatar, :bio, :fcm_token)
        RETURNING id, email, username, created_at, updated_at`

	rows, err := s.db.NamedQueryContext(ctx, query, user)

	if err != nil {
		return err
	}

	defer rows.Close()

	if rows.Next() {
		return rows.Scan(&user.Id, &user.Email, &user.Username, &user.CreatedAt, &user.UpdatedAt)
	}

	return sql.ErrNoRows
}
