package handlers

import "github.com/malachi190/watchcircle/repository"

type Handler struct {
	Auth *AuthHandler
}

func New(repo *repository.Repo) *Handler {
	return &Handler{
		Auth: NewAuthHandler(repo),
	}
}
