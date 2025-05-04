package repo

import (
	"context"

	"github.com/jackc/pgx/v5"
	"github.com/merema-uit/server/models"
)

type Credentials struct {
	CitizenID    string `json:"citizen_id" db:"citizen_id"`
	PasswordHash string `json:"password_hash" db:"password_hash"`
	Role         string `json:"role" db:"role"`
}

func GetCredentialsByCitizenID(ctx context.Context, citizenID string) (Credentials, error) {
	mutexLock.RLock()
	defer mutexLock.RUnlock()

	const query = `
		SELECT password_hash, role
		FROM accounts
		WHERE citizen_id = $1
	`

	var creds Credentials
	err := dbpool.QueryRow(ctx, query, citizenID).Scan(&creds.PasswordHash, &creds.Role)
	creds.CitizenID = citizenID

	if err != nil {
		if err == pgx.ErrNoRows {
			return Credentials{}, models.ErrCitizenIDNotExists
		}
	}

	return creds, err
}

func CheckCitizenIDExists(ctx context.Context, citizenID string) error {
	mutexLock.RLock()
	defer mutexLock.RUnlock()

	var citizenIDExists bool
	err := dbpool.QueryRow(ctx,
		"SELECT EXISTS(SELECT 1 FROM accounts WHERE citizen_id = $1)",
		citizenID,
	).Scan(&citizenIDExists)

	if citizenIDExists {
		return models.ErrCitizenIDExists
	}
	return err
}

func GetEmailByCitizenID(ctx context.Context, citizenID string) (string, error) {
	mutexLock.RLock()
	defer mutexLock.RUnlock()

	const query = `
		SELECT email
		FROM accounts
		WHERE citizen_id = $1
	`

	var email string
	err := dbpool.QueryRow(ctx, query, citizenID).Scan(&email)
	if err != nil {
		if err == pgx.ErrNoRows {
			return "", models.ErrCitizenIDNotExists
		}
		return "", err
	}

	return email, nil
}

func StoreAccountInfo(ctx context.Context, req models.AccountRegisterRequest) error {
	mutexLock.Lock()
	defer mutexLock.Unlock()

	const query = `
		INSERT INTO accounts (citizen_id, password_hash, phone, email, role)
		VALUES ($1, $2, $3, $4, $5)
		RETURNING citizen_id
	`
	var createdUsername string
	err := dbpool.QueryRow(ctx, query, req.CitizenID, req.Password, req.Phone, req.Email, req.Role).Scan(&createdUsername)
	return err
}

func UpdatePassword(ctx context.Context, req models.PasswordResetRequest) error {
	mutexLock.Lock()
	defer mutexLock.Unlock()

	const query = `
		UPDATE accounts
		SET password_hash = $1
		WHERE citizen_id = $2
	`

	_, err := dbpool.Exec(ctx, query, req.NewPassword, req.CitizenID)
	if err != nil {
		if err == pgx.ErrNoRows {
			return models.ErrCitizenIDNotExists
		}
		return err
	}

	return nil
}
