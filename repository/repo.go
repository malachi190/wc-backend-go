package repository

import (
	"github.com/malachi190/watchcircle/config"
	"github.com/malachi190/watchcircle/database"
	"github.com/malachi190/watchcircle/models"
	"github.com/malachi190/watchcircle/service"
)

type Repo struct {
	Config  *config.Config
	DB      *database.DB
	Model   *models.Models
	Service *service.Service
}

func New(cfg *config.Config, db *database.DB, model *models.Models, service *service.Service) *Repo {
	return &Repo{
		Config:  cfg,
		DB:      db,
		Model:   model,
		Service: service,
	}
}
