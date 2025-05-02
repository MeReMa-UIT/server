package repo

import (
	"context"

	"github.com/jackc/pgx/v5"
)

const DATABASE_URL = "postgres://postgres:pg@localhost:5432/merema"

func ConnectToDB(ctx context.Context, connString string) (*pgx.Conn, error) {
	conn, err := pgx.Connect(ctx, connString)
	if err != nil {
		return nil, err
	}
	return conn, nil
}
