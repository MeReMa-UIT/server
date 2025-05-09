package repo

import (
	"context"
	"errors"
	"strings"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/merema-uit/server/models"
	errs "github.com/merema-uit/server/models/errors"
)

type Credentials struct {
	CitizenID    string `json:"citizen_id" db:"citizen_id"`
	PasswordHash string `json:"password_hash" db:"password_hash"`
	Role         string `json:"role" db:"role"`
}

func GetAccountCredentials(ctx context.Context, accIdentifier string) (Credentials, error) {
	const query = `
		SELECT citizen_id, password_hash, role
		FROM accounts
		WHERE citizen_id = $1 OR phone = $1 OR email = $1
		LIMIT 1
	`

	rows, _ := dbpool.Query(ctx, query, accIdentifier)
	creds, err := pgx.CollectExactlyOneRow(rows, pgx.RowToStructByName[Credentials])

	if err != nil {
		if err == pgx.ErrNoRows {
			return Credentials{}, errs.ErrAccountNotExist
		}
	}

	return creds, nil
}

func CheckEmailAndCitizenID(ctx context.Context, req models.AccountRecoverRequest) error {
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
		return errs.ErrAccountNotExist
	}

	return nil
}

func GetAccIDByCitizenID(ctx context.Context, citizenID string) (int, error) {
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
			return -1, errs.ErrAccountNotExist
		}
		return -1, err
	}
	return accID, nil
}

func StoreAccountInfo(ctx context.Context, req models.AccountRegistrationRequest, citizenID, password_hash string) (int, error) {
	const query = `
		INSERT INTO accounts (citizen_id, password_hash, phone, email, role)
		VALUES ($1, $2, $3, $4, $5)
		RETURNING acc_id
	`
	tx, err := dbpool.BeginTx(ctx, pgx.TxOptions{
		IsoLevel: pgx.Serializable,
	})
	if err != nil {
		return -1, err
	}
	defer tx.Rollback(ctx)

	var createdAccID int
	err = tx.QueryRow(ctx, query, citizenID, password_hash, req.Phone, req.Email, req.Role).Scan(&createdAccID)

	if err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) {
			switch pgErr.Code {
			case "23505":
				if strings.Contains(pgErr.ConstraintName, "unique_citizen_id") {
					return -1, errs.ErrCitizenIDExists
				}
				if strings.Contains(pgErr.ConstraintName, "unique_phone") || strings.Contains(pgErr.ConstraintName, "unique_email") {
					return -1, errs.ErrEmailOrPhoneAlreadyUsed
				}
			}
		}
		return -1, err
	}
	if err := tx.Commit(ctx); err != nil {
		return -1, err
	}
	return createdAccID, nil
}

func UpdatePassword(ctx context.Context, citizenID, newPassword string) error {
	const query = `
		UPDATE accounts
		SET password_hash = $1
		WHERE citizen_id = $2
	`
	tx, err := dbpool.BeginTx(ctx, pgx.TxOptions{
		IsoLevel: pgx.Serializable,
	})
	if err != nil {
		return err
	}
	defer tx.Rollback(ctx)

	_, err = tx.Exec(ctx, query, newPassword, citizenID)
	if err != nil {
		if err == pgx.ErrNoRows {
			return errs.ErrAccountNotExist
		}
		return err
	}
	return tx.Commit(ctx)
}
