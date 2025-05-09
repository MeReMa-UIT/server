package repo

import (
	"context"
	"log"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
)

var (
	dbpool *pgxpool.Pool
)

func ConnectToDB(ctx context.Context, connString string) error {
	config, err := pgxpool.ParseConfig(connString)
	if err != nil {
		log.Println("Unable to connect to database:", err)
		return err
	}
	config.MaxConns = 100
	config.MinConns = 20
	config.MaxConnLifetime = 30 * time.Minute
	config.MaxConnIdleTime = 5 * time.Minute
	config.HealthCheckPeriod = 1 * time.Minute

	dbpool, err = pgxpool.NewWithConfig(ctx, config)
	if err != nil {
		log.Println("Unable to connect to database:", err)
		return err
	}
	return nil
}

func CloseDBConnect() {
	if dbpool != nil {
		dbpool.Close()
	}
}
