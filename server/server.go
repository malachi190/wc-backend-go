package server

import (
	"context"
	"errors"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type Server struct {
	*gin.Engine
	srv *http.Server
}

func New(addr string, routes func(*gin.Engine)) *Server {
	router := gin.New()
	router.Use(gin.Recovery())

	if routes != nil {
		routes(router)
	}

	srv := &http.Server{
		Addr:    addr,
		Handler: router,
	}

	return &Server{Engine: router, srv: srv}
}

// Start runs the server in a goroutine and returns immediately
func (s *Server) Start() error {
	go func() {
		if err := s.srv.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			log.Printf("server listen error: %v", err)
		}
	}()
	return nil
}

func (s *Server) Stop() error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	return s.srv.Shutdown(ctx)
}
