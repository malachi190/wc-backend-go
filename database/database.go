package database

import (
	"context"
	"fmt"
	"time"

	_ "github.com/lib/pq"
	"github.com/jmoiron/sqlx"
)

type DB struct {
	*sqlx.DB
}

func New(dataSourceName string, maxOpen, maxIdle int, maxLifetime time.Duration) (*DB, error) {
	sqldb, err := sqlx.Connect("postgres", dataSourceName)

	if err != nil {
		return nil, fmt.Errorf("open: %w", err)
	}

	sqldb.SetMaxOpenConns(maxOpen)
	sqldb.SetMaxIdleConns(maxIdle)
	sqldb.SetConnMaxLifetime(maxLifetime)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := sqldb.PingContext(ctx); err != nil {
		return nil, fmt.Errorf("ping: %w", err)
	}

	return &DB{sqldb}, nil
}
