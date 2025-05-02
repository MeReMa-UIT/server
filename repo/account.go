package repo

import (
	"context"

	"github.com/jackc/pgx/v5"
	"github.com/merema-uit/server/models"
)

type Credentials struct {
	Username     string `json:"username" db:"username"`
	PasswordHash string `json:"password_hash" db:"password_hash"`
	Role         string `json:"role" db:"role"`
}

func GetCredentialsByUsername(ctx context.Context, username string) (Credentials, error) {
	conn, err := ConnectToDB(ctx, DATABASE_URL)
	if err != nil {
		return Credentials{}, err
	}
	defer conn.Close(ctx)

	const query = `
		SELECT password_hash, role
		FROM accounts
		WHERE username = $1
	`

	var creds Credentials
	err = conn.QueryRow(ctx, query, username).Scan(&creds.PasswordHash, &creds.Role)
	creds.Username = username

	if err != nil {
		if err == pgx.ErrNoRows {
			return Credentials{}, models.ErrUsernameNotExists
		}
	}

	return creds, err
}
