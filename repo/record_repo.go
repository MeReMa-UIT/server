package repo

import (
	"context"

	"github.com/jackc/pgx/v5"
)

func GetRecordIDListByAccID(ctx context.Context, accID string) ([]int, error) {
	const query = `
		SELECT record_id
		FROM records r JOIN patients p ON r.patient_id = p.patient_id
		WHERE acc_id = $1::BIGINT
	`

	var recordIDList []int
	rows, _ := dbpool.Query(ctx, query, accID)
	recordIDList, err := pgx.AppendRows(recordIDList, rows, pgx.RowTo[int])
	if err != nil {
		return nil, err
	}
	return recordIDList, nil
}

func StoreMedicalRecord(ctx context.Context) error {
	const query = `
		INSERT INTO records (patient_id, doctor_id)
		VALUES ()
		ON CONFLICT () DO NOTHING
	`

	tx, err := dbpool.BeginTx(ctx, pgx.TxOptions{
		IsoLevel: pgx.ReadCommitted,
	})
	if err != nil {
		return err
	}
	defer tx.Rollback(ctx)

	_, err = tx.Exec(ctx, query)
	if err != nil {
		return err
	}
	return nil
}
