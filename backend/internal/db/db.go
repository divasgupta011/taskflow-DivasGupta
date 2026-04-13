package db

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"time"

	_ "github.com/lib/pq"

	"taskflow/internal/config"
)

func New(cfg *config.Config) (*sql.DB, error) {
	dsn := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		cfg.DBHost,
		cfg.DBPort,
		cfg.DBUser,
		cfg.DBPassword,
		cfg.DBName,
	)

	var db *sql.DB
	var err error

	// Retry logic (important for Docker startup timing)
	for i := 0; i < 5; i++ {
		db, err = sql.Open("postgres", dsn)
		if err != nil {
			return nil, err
		}

		ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
		defer cancel()

		err = db.PingContext(ctx)
		if err == nil {
			log.Println("connected to database")
			return db, nil
		}

		log.Println("waiting for database...")
		time.Sleep(2 * time.Second)
	}

	return nil, fmt.Errorf("could not connect to database: %w", err)
}
