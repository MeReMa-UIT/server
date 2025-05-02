package repo

import (
	"context"
	"log"

	"github.com/jackc/pgx/v5/pgxpool"
)

var dbpool *pgxpool.Pool

func ConnectToDB(ctx context.Context, connString string) {
	var err error
	dbpool, err = pgxpool.New(ctx, connString)
	if err != nil {
		log.Println("Unable to connect to database:", err)
	}
}

func CloseDB() {
	if dbpool != nil {
		dbpool.Close()
	}
}
