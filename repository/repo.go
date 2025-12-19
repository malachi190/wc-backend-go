package repository

import (
	"github.com/malachi190/watchcircle/config"
	"github.com/malachi190/watchcircle/database"
	"github.com/malachi190/watchcircle/server"
)

type App struct {
	Config *config.Config
	DB     *database.DB
	Server *server.Server
}
