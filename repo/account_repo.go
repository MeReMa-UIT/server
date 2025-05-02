package repo

import (
	"context"

	"github.com/jackc/pgx/v5"
	"github.com/merema-uit/server/models"
)

type Credentials struct {
	Username     string `json:"citizen_id" db:"citizen_id"`
	PasswordHash string `json:"password_hash" db:"password_hash"`
	Role         string `json:"role" db:"role"`
}

func GetCredentialsByCitizenID(ctx context.Context, username string) (Credentials, error) {
	conn, err := ConnectToDB(ctx, DATABASE_URL)
	if err != nil {
		return Credentials{}, err
	}
	defer conn.Close(ctx)

	const query = `
		SELECT password_hash, role
		FROM accounts
		WHERE citizen_id = $1
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

func CheckCitizenIDExists(ctx context.Context, citizenID string) error {
	conn, err := ConnectToDB(ctx, DATABASE_URL)
	if err != nil {
		return err
	}
	defer conn.Close(ctx)
	var citizenIDExists bool
	err = conn.QueryRow(ctx,
		"SELECT EXISTS(SELECT 1 FROM accounts WHERE citizen_id = $1)",
		citizenID,
	).Scan(&citizenIDExists)

	if citizenIDExists {
		return models.ErrCitizenIDExists
	}
	return err
}

func StoreAccountInfo(ctx context.Context, req models.AccountRegisterRequest) error {
	conn, err := ConnectToDB(ctx, DATABASE_URL)
	if err != nil {
		return err
	}
	defer conn.Close(ctx)

	const query = `
		INSERT INTO accounts (citizen_id, password_hash, phone, email, role)
		VALUES ($1, $2, $3, $4, $5)
		ON CONFLICT (citizen_id) DO NOTHING
		RETURNING citizen_id
	`
	var createdUsername string
	err = conn.QueryRow(ctx, query, req.CitizenID, req.Password, req.Phone, req.Email, req.Role).Scan(&createdUsername)
	return err
}
