package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/malachi190/watchcircle/config"
	"github.com/malachi190/watchcircle/database"
	"github.com/malachi190/watchcircle/handlers"
	"github.com/malachi190/watchcircle/models"
	"github.com/malachi190/watchcircle/repository"
	"github.com/malachi190/watchcircle/routes"
	"github.com/malachi190/watchcircle/server"
	"github.com/malachi190/watchcircle/service"
)

func main() {
	// Load Config
	cfg, err := config.LoadConfig()

	if err != nil {
		log.Fatalf("config error: %v", err)
	}

	// Initialize DB
	db, err := database.New(cfg.DatabaseUrl, 25, 25, 5*time.Minute)

	if err != nil {
		log.Fatalf("database error: %v", err)
	}

	// Service
	service := service.NewService(db.DB)

	// Repo
	repo := repository.New(cfg, db, models.New(), service)

	// Routes
	r := routes.Routes(handlers.New(repo))

	// Server
	srv := server.New(fmt.Sprintf(":%s", cfg.App.Port), r)

	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()

	// Start server
	if err := srv.Start(); err != nil {
		log.Fatalf("server start: %v", err)
	}

	<-ctx.Done()
	log.Println("shutting down server...")

	if err := srv.Stop(); err != nil {
		log.Fatalf("server stop: %v", err)
	}

	log.Println("done")
}
