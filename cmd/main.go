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
	"github.com/malachi190/watchcircle/logger"
	"github.com/malachi190/watchcircle/mailer"
	"github.com/malachi190/watchcircle/models"
	"github.com/malachi190/watchcircle/repository"
	"github.com/malachi190/watchcircle/routes"
	"github.com/malachi190/watchcircle/server"
	"github.com/malachi190/watchcircle/service"
	"github.com/redis/go-redis/v9"
)

func main() {
	// Init Logger
	logger.Init("./logs")

	// Create goroutine to handle file change in background
	go func() {
		last := time.Now().Format("2006-01-02")
		for {
			now := time.Now()
			if now.Format("2006-01-02") != last {
				logger.Init("logs")
				last = now.Format("2006-01-02")
				log.Println("switched to new daily log file")
			}
			time.Sleep(time.Minute)
		}
	}()

	// select {} ---> replace select{} with a job so the file change goroutine can run

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

	// Redis
	var RDB = redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%s", cfg.Redis.Host, cfg.Redis.Port),
		Password: "",
		DB:       0,
	})

	defer RDB.Close()

	// Mailer
	newMlr, err := mailer.NewMailer(cfg.Resend.ApiKey)

	if err != nil {
		log.Fatalf("mailer error: %v", err)
	}

	// Repo
	repo := repository.New(cfg, db, models.New(), service, RDB, newMlr)

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
