package store

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	_ "github.com/jackc/pgx/v4/stdlib"
)

func Open() (*sql.DB, error) {
	db, err := sql.Open("pgx", "host=localhost user=pulse password=pulse_pwd dbname=pulse_db port=5432 sslmode=disable")

	if err != nil {
		return nil, fmt.Errorf("failed to open database connection: %w", err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	pinging := db.PingContext(ctx)

	if pinging != nil {
		return nil, fmt.Errorf("failed to ping database: %w", err)
	}

	fmt.Println("Connecting to Database ....")

	return db, nil
}
