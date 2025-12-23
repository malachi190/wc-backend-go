package service

import "github.com/jmoiron/sqlx"

type Service struct {
	Auth *AuthService
}

func NewService(db *sqlx.DB) *Service {
	return &Service{
		Auth: NewAuthService(db),
	}
}
