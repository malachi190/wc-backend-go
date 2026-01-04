package types

type AddToListRequest struct {
	TmdbID    int    `json:"tmdb_id" binding:"required"`
	MediaType string `json:"media_type" binding:"required"`
	Title     string `json:"title" binding:"required"`
	PosterPath *string `json:"poster_path,omitempty"`

	// TODO: Make sure to add extra fields to watchlist table, i.e year of realease, genre, duration, etc
}
