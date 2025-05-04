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

func GetAccountCredentials(ctx context.Context, accIdentifier string) (Credentials, error) {
	accLock.RLock()
	defer accLock.RUnlock()

	const query = `
		SELECT citizen_id, password_hash, role
		FROM accounts
		WHERE citizen_id = $1 OR phone = $1 OR email = $1
	`

	var creds Credentials
	err := dbpool.QueryRow(ctx, query, accIdentifier).Scan(&creds.CitizenID, &creds.PasswordHash, &creds.Role)

	if err != nil {
		if err == pgx.ErrNoRows {
			return Credentials{}, models.ErrAccountNotExist
		}
	}

	return creds, err
}

func GetEmailByCitizenID(ctx context.Context, citizenID string) (string, error) {
	accLock.RLock()
	defer accLock.RUnlock()

	const query = `
		SELECT email
		FROM accounts
		WHERE citizen_id = $1
	`

	var email string
	err := dbpool.QueryRow(ctx, query, citizenID).Scan(&email)
	if err != nil {
		if err == pgx.ErrNoRows {
			return "", models.ErrAccountNotExist
		}
		return "", err
	}

	return email, nil
}

func GetAccIDByCitizenID(ctx context.Context, citizenID string) (int, error) {
	accLock.RLock()
	defer accLock.RUnlock()
	const query = `
		SELECT acc_id 
		FROM accounts
		WHERE citizen_id = $1
	`
	var accID int
	err := dbpool.QueryRow(ctx, query, citizenID).Scan(&accID)
	if err != nil {
		if err == pgx.ErrNoRows {
			return -1, models.ErrAccountNotExist
		}
		return -1, err
	}
	return accID, nil
}

func StoreAccountInfo(ctx context.Context, req models.AccountRegisterRequest) (int, error) {
	accLock.Lock()
	defer accLock.Unlock()

	var emailOrPhoneExists bool
	err := dbpool.QueryRow(ctx, "SELECT EXISTS (SELECT 1 FROM accounts WHERE email = $1 OR phone = $2)", req.Email, req.Phone).Scan(&emailOrPhoneExists)

	if err != nil {
		return -1, err
	}

	if emailOrPhoneExists {
		return -1, models.ErrEmailOrPhoneAlreadyUsed
	}

	const query = `
		INSERT INTO accounts (citizen_id, password_hash, phone, email, role)
		VALUES ($1, $2, $3, $4, $5)
		RETURNING acc_id
	`
	var createdAccID int
	err = dbpool.QueryRow(ctx, query, req.CitizenID, req.Password, req.Phone, req.Email, req.Role).Scan(&createdAccID)

	if err != nil {
		return -1, err
	}

	return createdAccID, nil
}

func UpdatePassword(ctx context.Context, req models.PasswordResetRequest) error {
	accLock.Lock()
	defer accLock.Unlock()

	const query = `
		UPDATE accounts
		SET password_hash = $1
		WHERE citizen_id = $2
	`

	_, err := dbpool.Exec(ctx, query, req.NewPassword, req.CitizenID)
	if err != nil {
		if err == pgx.ErrNoRows {
			return models.ErrAccountNotExist
		}
		return err
	}

	return nil
}
