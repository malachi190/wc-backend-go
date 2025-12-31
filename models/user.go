package models

import (
	"time"

	"github.com/google/uuid"
)

type UserModel struct {
	Id              uuid.UUID  `json:"id" db:"id"`
	Email           string     `json:"email" db:"email"`
	DisplayName     string     `json:"display_name" db:"display_name"`
	Username        string     `json:"username" db:"username"`
	Password        string     `json:"password" db:"password"`
	EmailVerifiedAt *time.Time `json:"email_verified_at,omitempty" db:"email_verified_at"`
	Avatar          *string    `json:"avatar,omitempty" db:"avatar"`
	Bio             *string    `json:"bio,omitempty" db:"bio"`
	FcmToken        *string    `json:"fcm_token,omitempty" db:"fcm_token"`
	CreatedAt       time.Time  `json:"created_at" db:"created_at"`
	UpdatedAt       time.Time  `json:"updated_at" db:"updated_at"`
}
