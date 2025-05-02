package repo

import (
	"context"
	"log"

	"github.com/jackc/pgx/v5/pgxpool"
)

const DATABASE_URL = "postgres://postgres:pg@localhost:5432/merema"

var dbpool *pgxpool.Pool
var err error

func ConnectToDB(ctx context.Context, connString string) {
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
