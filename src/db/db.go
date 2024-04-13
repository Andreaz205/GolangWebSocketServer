package db

import (
	"database/sql"
	_ "github.com/lib/pq"
	"golang-websocket-chat-server/internal/config"
)

type Database struct {
	db *sql.DB
}

func NewDatabase(cfg *config.Config) (*Database, error) {
	db, err := sql.Open(cfg.DatabaseDriver, cfg.DatabaseConnection+
		"://"+cfg.DatabaseUser+":"+cfg.DatabasePassword+
		"@"+cfg.DatabaseHost+":"+cfg.DatabasePort+
		"/"+cfg.DatabaseDb+"?sslmode=disable")

	if err != nil {
		return nil, err
	}

	return &Database{db: db}, nil
}

func (d *Database) Close() error {
	return d.db.Close()
}

func (d *Database) GetDB() *sql.DB {
	return d.db
}
