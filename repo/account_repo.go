package repo

import (
	"context"

	"github.com/jackc/pgx/v5"
	"github.com/merema-uit/server/models"
	"github.com/merema-uit/server/models/errors"
)

type Credentials struct {
	CitizenID    string `json:"citizen_id" db:"citizen_id"`
	PasswordHash string `json:"password_hash" db:"password_hash"`
	Role         string `json:"role" db:"role"`
}

func GetAccountCredentials(ctx context.Context, accIdentifier string) (Credentials, error) {
	accountTableLock.RLock()
	defer accountTableLock.RUnlock()

	const query = `
		SELECT citizen_id, password_hash, role
		FROM accounts
		WHERE citizen_id = $1 OR phone = $1 OR email = $1
		LIMIT 1
	`

	var creds Credentials
	err := dbpool.QueryRow(ctx, query, accIdentifier).Scan(&creds.CitizenID, &creds.PasswordHash, &creds.Role)

	if err != nil {
		if err == pgx.ErrNoRows {
			return Credentials{}, errors.ErrAccountNotExist
		}
	}

	return creds, err
}

func CheckEmailAndCitizenID(ctx context.Context, req models.AccountRecoverRequest) error {
	accountTableLock.RLock()
	defer accountTableLock.RUnlock()

	const query = `
		SELECT EXISTS(
			SELECT 1
			FROM accounts
			WHERE citizen_id = $1 AND email = $2
		)
	`

	var exist bool
	err := dbpool.QueryRow(ctx, query, req.CitizenID, req.Email).Scan(&exist)
	if err != nil {
		return err
	}
	if exist == false {
		return errors.ErrAccountNotExist
	}

	return nil
}

func GetAccIDByCitizenID(ctx context.Context, citizenID string) (int, error) {
	accountTableLock.RLock()
	defer accountTableLock.RUnlock()
	const query = `
		SELECT acc_id 
		FROM accounts
		WHERE citizen_id = $1
		LIMIT 1
	`
	var accID int
	err := dbpool.QueryRow(ctx, query, citizenID).Scan(&accID)
	if err != nil {
		if err == pgx.ErrNoRows {
			return -1, errors.ErrAccountNotExist
		}
		return -1, err
	}
	return accID, nil
}

func StoreAccountInfo(ctx context.Context, req models.AccountRegistrationRequest, citizenID, password_hash string) (int, error) {
	accountTableLock.Lock()
	defer accountTableLock.Unlock()

	var emailOrPhoneExists bool
	err := dbpool.QueryRow(ctx, "SELECT EXISTS (SELECT 1 FROM accounts WHERE email = $1 OR phone = $2)", req.Email, req.Phone).Scan(&emailOrPhoneExists)

	if err != nil {
		return -1, err
	}

	if emailOrPhoneExists {
		return -1, errors.ErrEmailOrPhoneAlreadyUsed
	}

	const query = `
		INSERT INTO accounts (citizen_id, password_hash, phone, email, role)
		VALUES ($1, $2, $3, $4, $5)
		RETURNING acc_id
	`
	var createdAccID int
	err = dbpool.QueryRow(ctx, query, citizenID, password_hash, req.Phone, req.Email, req.Role).Scan(&createdAccID)

	if err != nil {
		return -1, err
	}

	return createdAccID, nil
}

func UpdatePassword(ctx context.Context, citizenID, newPassword string) error {
	accountTableLock.Lock()
	defer accountTableLock.Unlock()

	const query = `
		UPDATE accounts
		SET password_hash = $1
		WHERE citizen_id = $2
	`

	_, err := dbpool.Exec(ctx, query, newPassword, citizenID)
	if err != nil {
		if err == pgx.ErrNoRows {
			return errors.ErrAccountNotExist
		}
		return err
	}

	return nil
}
