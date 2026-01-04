package models

import (
	"time"

	"github.com/google/uuid"
)

type WatchList struct {
	ID         uuid.UUID `json:"id" db:"id"`
	UserID     uuid.UUID `json:"user_id" db:"user_id"`
	TmdbID     int       `json:"tmdb_id" db:"tmdb_id"`
	MediaType  string    `json:"media_type" db:"media_type"`
	Title      string    `json:"title" db:"title"`
	PosterPath *string   `json:"poster_path" db:"poster_path"`
	Status     string    `json:"status" db:"status"`
	Rating     *int32    `json:"rating" db:"rating"`
	Notes      *string   `json:"notes" db:"notes"`
	AddedAt    time.Time `json:"added_at" db:"added_at"`
	UpdatedAt  time.Time `json:"updated_at" db:"updated_at"`

	User *UserModel `json:"user,omitempty"`
}
