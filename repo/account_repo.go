package repo

import (
	"context"
	"errors"
	"fmt"
	"strconv"
	"strings"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/merema-uit/server/models"
	errs "github.com/merema-uit/server/models/errors"
)

type Credentials struct {
	AccID        string `json:"acc_id" db:"acc_id"`
	PasswordHash string `json:"password_hash" db:"password_hash"`
	Role         string `json:"role" db:"role"`
}

func GetAccountCredentials(ctx context.Context, accIdentifier string) (Credentials, error) {
	const query = `
		SELECT acc_id, password_hash, role
		FROM accounts
		WHERE citizen_id = $1 OR phone = $1 OR email = $1 OR acc_id = $2
	`

	accID, err := strconv.Atoi(accIdentifier)
	println("accID:", accID)

	rows, _ := dbpool.Query(ctx, query, accIdentifier, accID)
	creds, err := pgx.CollectExactlyOneRow(rows, pgx.RowToStructByName[Credentials])

	if err != nil {
		if err == pgx.ErrNoRows {
			return Credentials{}, errs.ErrAccountNotExist
		}
	}

	return creds, nil
}

func GetAccountList(ctx context.Context) ([]models.AccountInfo, error) {
	const query = `
		SELECT acc_id, citizen_id, phone, email, role, created_at
		FROM accounts
	`
	rows, _ := dbpool.Query(ctx, query)
	list, err := pgx.CollectRows(rows, pgx.RowToStructByName[models.AccountInfo])
	if err != nil {
		return nil, err
	}
	return list, nil
}

func GetAccountInfo(ctx context.Context, accID string) (models.AccountInfo, error) {
	const query = `
		SELECT acc_id, citizen_id, phone, email, role, created_at
		FROM accounts
		WHERE acc_id = $1
	`

	rows, _ := dbpool.Query(ctx, query, accID)
	accountInfo, err := pgx.CollectExactlyOneRow(rows, pgx.RowToStructByName[models.AccountInfo])

	if err != nil {
		if err == pgx.ErrNoRows {
			return models.AccountInfo{}, errs.ErrAccountNotExist
		}
		return models.AccountInfo{}, err
	}

	return accountInfo, nil
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

func StoreAccountInfo(ctx context.Context, req models.AccountRegistrationRequest, password_hash string) (int, error) {
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
	err = tx.QueryRow(ctx, query, req.CitizenID, password_hash, req.Phone, req.Email, req.Role).Scan(&createdAccID)

	if err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) {
			switch pgErr.Code {
			case "23505":
				if strings.Contains(pgErr.ConstraintName, "accounts_citizen_id_key") {
					return -1, errs.ErrCitizenIDExists
				}
				if strings.Contains(pgErr.ConstraintName, "accounts_email_key") || strings.Contains(pgErr.ConstraintName, "accounts_phone_key") {
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

func UpdatePassword(ctx context.Context, accID, newPassword string) error {
	const query = `
		UPDATE accounts
		SET password_hash = $1
		WHERE acc_id = $2
		RETURNING acc_id
	`
	tx, err := dbpool.BeginTx(ctx, pgx.TxOptions{
		IsoLevel: pgx.Serializable,
	})
	if err != nil {
		return err
	}
	defer tx.Rollback(ctx)

	var updatedAccID int
	err = tx.QueryRow(ctx, query, newPassword, accID).Scan(&updatedAccID)
	if err != nil {
		if err == pgx.ErrNoRows {
			return errs.ErrAccountNotExist
		}
		return err
	}
	if err := tx.Commit(ctx); err != nil {
		return err
	}
	return nil
}

func UpdateAccountInfo(ctx context.Context, accID string, req models.UpdateAccountInfoRequest) error {
	query := fmt.Sprintf(`
		UPDATE accounts
		SET %s = $1
		WHERE acc_id = $2
		RETURNING acc_id
	`, req.Field)

	tx, err := dbpool.BeginTx(ctx, pgx.TxOptions{
		IsoLevel: pgx.Serializable,
	})
	if err != nil {
		return err
	}
	defer tx.Rollback(ctx)

	var updatedAccID int
	err = tx.QueryRow(ctx, query, req.NewValue, accID).Scan(&updatedAccID)

	println(updatedAccID)

	if err != nil {
		if err == pgx.ErrNoRows {
			return errs.ErrAccountNotExist
		}
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) {
			switch pgErr.Code {
			case "23505":
				if strings.Contains(pgErr.ConstraintName, "accounts_citizen_id_key") {
					return errs.ErrCitizenIDExists
				}
				if strings.Contains(pgErr.ConstraintName, "accounts_email_key") || strings.Contains(pgErr.ConstraintName, "accounts_phone_key") {
					return errs.ErrEmailOrPhoneAlreadyUsed
				}
			}
		}
		return err
	}

	if err := tx.Commit(ctx); err != nil {
		return err
	}
	return nil
}
