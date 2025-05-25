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

func StoreStaffInfo(ctx context.Context, req models.StaffRegistrationRequest) error {
	const query = `
		INSERT INTO staffs (acc_id, full_name, date_of_birth, gender, department)
		VALUES ($1, $2, $3, $4, $5)
		RETURNING staff_id
	`
	tx, err := dbpool.BeginTx(ctx, pgx.TxOptions{
		IsoLevel: pgx.Serializable,
	})
	if err != nil {
		return err
	}
	defer tx.Rollback(ctx)

	var staffID int
	err = tx.QueryRow(ctx, query,
		req.AccID,
		req.FullName,
		req.DateOfBirth,
		req.Gender,
		req.Department,
	).Scan(&staffID)

	if err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) {
			switch pgErr.Code {
			case "23505":
				if strings.Contains(pgErr.ConstraintName, "accounts_acc_id_key") {
					return errs.ErrAccountAlreadyLinked
				}
			}
		}
		return err
	}

	return tx.Commit(ctx)
}

func GetStaffList(ctx context.Context) ([]models.StaffInfo, error) {
	const query = `
		SELECT staff_id, full_name, date_of_birth, gender, department
		FROM staffs
	`
	rows, _ := dbpool.Query(ctx, query)
	list, err := pgx.CollectRows(rows, pgx.RowToStructByName[models.StaffInfo])
	if err != nil {
		return nil, err
	}
	return list, nil
}

func GetStaffInfo(ctx context.Context, staffID string, accID string) (models.StaffInfo, error) {
	const query = `
		SELECT staff_id, full_name, date_of_birth, gender, department 
		FROM staffs
		WHERE staff_id = $1 OR acc_id = $2
	`

	rows, _ := dbpool.Query(ctx, query, staffID, accID)
	staffInfo, err := pgx.CollectExactlyOneRow(rows, pgx.RowToStructByName[models.StaffInfo])

	if err != nil {
		return models.StaffInfo{}, err
	}

	return staffInfo, nil
}
