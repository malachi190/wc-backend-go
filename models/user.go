package models

import "time"

type UserModel struct {
	Base
	Email           string     `json:"email" db:"email"`
	Username        string     `json:"username" db:"username"`
	Password        string     `json:"password" db:"password"`
	EmailVerifiedAt *time.Time `json:"email_verified_at,omitempty" db:"email_verified_at"`
	Avatar          *string    `json:"avatar,omitempty" db:"avatar"`
	Bio             *string    `json:"bio,omitempty" db:"bio"`
	FcmToken        *string    `json:"fcm_token,omitempty" db:"fcm_token"`
}
