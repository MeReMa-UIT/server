package repo

import (
	"context"

	"github.com/merema-uit/server/models"
)

func StoreStaffInfo(ctx context.Context, req models.StaffRegistrationRequest, accID int) error {
	staffTableLock.Lock()
	defer staffTableLock.Unlock()

	const query = `
		INSERT INTO staffs (acc_id, full_name, date_of_birth, gender, department)
		VALUES ($1, $2, $3, $4, $5)
		RETURNING staff_id
	`

	var staffID int
	err := dbpool.QueryRow(ctx, query,
		accID,
		req.FullName,
		req.DateOfBirth,
		req.Gender,
		req.Department,
	).Scan(&staffID)

	return err
}
