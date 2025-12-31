package repository

import (
	"github.com/malachi190/watchcircle/config"
	"github.com/malachi190/watchcircle/database"
	"github.com/malachi190/watchcircle/mailer"
	"github.com/malachi190/watchcircle/models"
	"github.com/malachi190/watchcircle/service"
	"github.com/redis/go-redis/v9"
)

type Repo struct {
	Config  *config.Config
	DB      *database.DB
	Model   *models.Models
	Service *service.Service
	Redis   *redis.Client
	Mailer  *mailer.Mailer
}

func New(
	cfg *config.Config,
	db *database.DB,
	model *models.Models,
	service *service.Service,
	redis *redis.Client,
	mailer *mailer.Mailer) *Repo {
	return &Repo{
		Config:  cfg,
		DB:      db,
		Model:   model,
		Service: service,
		Redis:   redis,
		Mailer:  mailer,
	}
}
