package repo

import (
	"context"
	"time"

	"github.com/jackc/pgx/v5"
	"github.com/merema-uit/server/models"
)

func GetRecordListByTime(ctx context.Context, timestamp1, timestamp2 time.Time) ([]models.RecordInfoForStatistic, error) {
	const query = `
		SELECT patient_id, doctor_id, primary_diagnosis, created_at
		FROM records 
		WHERE created_at >= $1 AND created_at <= $2
		ORDER BY created_at
	`

	rows, _ := dbpool.Query(ctx, query, timestamp1, timestamp2)
	list, err := pgx.CollectRows(rows, pgx.RowToStructByName[models.RecordInfoForStatistic])

	if err != nil {
		return nil, err
	}

	return list, nil
}

func GetDoctorList(ctx context.Context) ([]int64, error) {
	const query = `
		SELECT staff_id
		FROM staffs JOIN accounts ON staffs.acc_id = accounts.acc_id
		WHERE role = 'doctor'
	`

	rows, _ := dbpool.Query(ctx, query)
	list, err := pgx.CollectRows(rows, pgx.RowTo[int64])

	if err != nil {
		return nil, err
	}

	return list, nil
}
