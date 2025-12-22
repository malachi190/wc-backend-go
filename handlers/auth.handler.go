package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/malachi190/watchcircle/repository"
)

type AuthHandler struct {
	Repo *repository.Repo
}

func NewAuthHandler(repo *repository.Repo) *AuthHandler {
	return &AuthHandler{
		Repo: repo,
	}
}

func (h *AuthHandler) Register(ctx *gin.Context) {
	
} 
