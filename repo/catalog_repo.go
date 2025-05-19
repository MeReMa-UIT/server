package repo

import (
	"context"

	"github.com/jackc/pgx/v5"
	"github.com/merema-uit/server/models"
)

func GetMedicationList(ctx context.Context) ([]models.MedicationInfo, error) {
	const query = `
		SELECT * FROM medications
	`
	rows, _ := dbpool.Query(ctx, query)
	list, err := pgx.CollectRows(rows, pgx.RowToStructByName[models.MedicationInfo])
	if err != nil {
		return nil, err
	}
	return list, nil
}

func GetDiagnosisList(ctx context.Context) ([]models.DiagnosisInfo, error) {
	const query = `
		SELECT * FROM diagnoses
	`
	rows, _ := dbpool.Query(ctx, query)
	list, err := pgx.CollectRows(rows, pgx.RowToStructByName[models.DiagnosisInfo])
	if err != nil {
		return nil, err
	}
	return list, nil
}
