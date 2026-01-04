package service

import "github.com/jmoiron/sqlx"

type WatchListService struct {
	db *sqlx.DB
}

func NewWatchListService(db *sqlx.DB) *WatchListService {
	return &WatchListService{
		db: db,
	}
}


