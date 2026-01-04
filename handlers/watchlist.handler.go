package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/malachi190/watchcircle/repository"
)

type WatchListHandler struct {
	Repo *repository.Repo
}

func NewWatchListHandler(repo *repository.Repo) *WatchListHandler {
	return &WatchListHandler{Repo: repo}
}


func (w *WatchListHandler) AddToList(ctx *gin.Context) {
	
}
