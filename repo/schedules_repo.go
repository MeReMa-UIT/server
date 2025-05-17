package repo

import (
	"context"
	"time"

	"github.com/jackc/pgx/v5"
	"github.com/merema-uit/server/models"
)

func GetQueueNumber(ctx context.Context, date time.Time) (int, error) {
	const query = `
		WITH updated AS (
				INSERT INTO queue_number (date, number)
				VALUES ($1, 2)
				ON CONFLICT (date) DO UPDATE
				SET number = queue_number.number + 1
				RETURNING number, (xmax = 0) AS is_insert
		)
		SELECT 
				CASE 
						WHEN is_insert THEN 1
						ELSE number - 1
				END AS number
		FROM updated;
	`

	tx, err := dbpool.BeginTx(ctx, pgx.TxOptions{
		IsoLevel: pgx.Serializable,
	})
	if err != nil {
		return -1, err
	}
	defer tx.Rollback(ctx)

	var queueNumber int
	err = tx.QueryRow(ctx, query, date).Scan(&queueNumber)
	if err != nil {
		return -1, err
	}
	if err = tx.Commit(ctx); err != nil {
		return -1, err
	}
	return queueNumber, nil
}

func CreateSchedule(ctx context.Context, req models.ScheduleBookingRequest, queue_number int, patientID string) (models.ScheduleBookingResponse, error) {
	const query = `
		INSERT INTO schedules (patient_id, examination_date, queue_number, type, expected_reception_time)
		VALUES ($1, $2, $3, $4, $5)
		RETURNING examination_date, type, queue_number, expected_reception_time, status
	`

	tx, err := dbpool.BeginTx(ctx, pgx.TxOptions{
		IsoLevel: pgx.Serializable,
	})
	if err != nil {
		return models.ScheduleBookingResponse{}, err
	}
	defer tx.Rollback(ctx)

	expectedTime := req.ExaminationDate.Add(time.Hour * 7).Add(time.Minute * 5 * time.Duration(queue_number)) // add 7 hours and 5 minutes for each queue number

	rows, _ := tx.Query(ctx, query, patientID, req.ExaminationDate, queue_number, req.Type, expectedTime)
	createdSchedule, err := pgx.CollectExactlyOneRow(rows, pgx.RowToStructByName[models.ScheduleBookingResponse])
	if err != nil {
		return models.ScheduleBookingResponse{}, err
	}

	if err = tx.Commit(ctx); err != nil {
		return models.ScheduleBookingResponse{}, err
	}

	return createdSchedule, nil
}

func GetScheduleList(ctx context.Context, patientID *int, req models.GetScheduleListRequest) ([]models.ScheduleInfo, error) {
	const query = `
		SELECT schedule_id, examination_date, type, queue_number, expected_reception_time, status
		FROM schedules
		WHERE ($1::BIGINT IS NULL OR patient_id = $1::BIGINT) 
			AND type = ANY($2) 
			AND status = ANY($3)
		ORDER BY examination_date, queue_number 
	`

	rows, _ := dbpool.Query(ctx, query, patientID, req.Type, req.Status)
	schedules, err := pgx.CollectRows(rows, pgx.RowToStructByName[models.ScheduleInfo])

	if err != nil {
		return nil, err
	}

	return schedules, nil
}
