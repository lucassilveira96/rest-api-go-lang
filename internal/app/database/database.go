package database

import (
	"context"
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"log"
	"time"
)

type Config struct {
	Host                  string
	Port                  string
	Database              string
	Username              string
	Password              string
	MinConnections        int
	MaxConnections        int
	ConnectionMaxLifetime time.Duration // 0, connections are reused forever
	ConnectionMaxIdleTime time.Duration
}

func NewPostgres(cfg *Config) (*sql.DB, error) {
	db, err := sql.Open("postgres", fmt.Sprintf(postgresConnectionString, cfg.Username, cfg.Password, cfg.Host, cfg.Port, cfg.Database))
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	timeout, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err = db.PingContext(timeout)
	if err != nil {
		return nil, err
	}

	db.SetMaxIdleConns(cfg.MinConnections)
	db.SetMaxOpenConns(cfg.MaxConnections)
	db.SetConnMaxLifetime(cfg.ConnectionMaxLifetime)
	db.SetConnMaxIdleTime(cfg.ConnectionMaxIdleTime)

	return db, err
}

const postgresConnectionString = "postgres://%s:%s@%s:%s/%s?sslmode=disable"
