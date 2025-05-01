package repo

import (
	"context"
	"fmt"

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

func GetCredentialsByUsername(ctx context.Context, username string) (string, error) {

	conn, err := ConnectToDB(ctx, DATABASE_URL)
	if err != nil {
		return "", err
	}
	defer conn.Close(ctx)

	const query = `
		SELECT password_hash
		FROM accounts
		WHERE username = $1
		LIMIT 1
	`

	var password_hash string
	err = conn.QueryRow(ctx, query, username).Scan(&password_hash)

	if err != nil {
		fmt.Println("Error querying database:", err)
		return "", err
	}

	return password_hash, nil
}
