package repo

import (
	"context"

	"github.com/jackc/pgx/v5"
	"github.com/merema-uit/server/models"
)

func StoreStaffInfo(ctx context.Context, req models.StaffRegistrationRequest, accID int) error {
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
		accID,
		req.FullName,
		req.DateOfBirth,
		req.Gender,
		req.Department,
	).Scan(&staffID)

	if err != nil {
		return err
	}

	return tx.Commit(ctx)
}
