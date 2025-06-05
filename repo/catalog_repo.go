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

func GetMedicationInfo(ctx context.Context, medicationID string) (models.MedicationInfo, error) {
	const query = `
		SELECT * FROM medications
		WHERE med_id = $1
	`

	rows, _ := dbpool.Query(ctx, query, medicationID)
	info, err := pgx.CollectExactlyOneRow(rows, pgx.RowToStructByName[models.MedicationInfo])
	if err != nil {
		return models.MedicationInfo{}, err
	}
	return info, nil
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

func GetDiagnosisInfo(ctx context.Context, icdCode string) (models.DiagnosisInfo, error) {
	const query = `
		SELECT * FROM diagnoses
		WHERE icd_code = $1
	`

	rows, _ := dbpool.Query(ctx, query, icdCode)
	info, err := pgx.CollectExactlyOneRow(rows, pgx.RowToStructByName[models.DiagnosisInfo])
	if err != nil {
		return models.DiagnosisInfo{}, err
	}
	return info, nil
}

func GetMedicalRecordTypeList(ctx context.Context) ([]models.MedicalRecordType, error) {
	const query = `
		SELECT type_id, type_name
		FROM record_types
	`
	rows, _ := dbpool.Query(ctx, query)
	list, err := pgx.CollectRows(rows, pgx.RowToStructByName[models.MedicalRecordType])
	if err != nil {
		return nil, err
	}
	return list, nil
}

func GetMedicalRecordType(ctx context.Context, typeID string) (models.MedicalRecordTypeInfo, error) {
	const query = `
		SELECT type_id, type_name, description, template_path, schema_path
		FROM record_types
		WHERE type_id = $1
	`

	rows, _ := dbpool.Query(ctx, query, typeID)
	info, err := pgx.CollectExactlyOneRow(rows, pgx.RowToStructByName[models.MedicalRecordTypeInfo])
	if err != nil {
		return models.MedicalRecordTypeInfo{}, err
	}
	return info, nil
}
