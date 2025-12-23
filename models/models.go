package models

import (
	"time"

	"github.com/google/uuid"
)

type BaseWithId struct {
	Id uuid.UUID `json:"id" db:"id"`
}

type BaseWithCreatedAt struct {
	CreatedAt time.Time `json:"created_at" db:"created_at"`
}

type BaseWithUpdatedAt struct {
	UpdatedAt time.Time `json:"updated_at" db:"created_at"`
}

type Base struct {
	BaseWithId
	BaseWithCreatedAt
	BaseWithUpdatedAt
}

type Models struct {
	User *UserModel
}

func New() *Models {
	return &Models{
		User: &UserModel{},
	}
}
