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
