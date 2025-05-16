package repo

import (
	"context"

	"github.com/jackc/pgx/v5"
	"github.com/merema-uit/server/models"
	errs "github.com/merema-uit/server/models/errors"
)

func GetPatientID(ctx context.Context, accID string) (int, error) {
	const query = `
		SELECT patient_id
		FROM patients
		WHERE acc_id = $1
	`
	var patientID int
	err := dbpool.QueryRow(ctx, query, accID).Scan(&patientID)
	if err != nil {
		if err == pgx.ErrNoRows {
			return -1, errs.ErrPatientNotExist
		}
		return -1, err
	}
	return patientID, nil
}

func StorePatientInfo(ctx context.Context, req models.PatientRegistrationRequest) error {
	const query = `
		INSERT INTO patients (acc_id, full_name, date_of_birth, gender, ethnicity, nationality, address,
		 											health_insurance_expired_date, health_insurance_number, emergency_contact_info)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10)
		RETURNING patient_id
	`

	tx, err := dbpool.BeginTx(ctx, pgx.TxOptions{
		IsoLevel: pgx.Serializable,
	})
	if err != nil {
		return err
	}
	defer tx.Rollback(ctx)

	var patientID int
	err = tx.QueryRow(ctx, query,
		req.AccID,
		req.FullName,
		req.DateOfBirth,
		req.Gender,
		req.Ethnicity,
		req.Nationality,
		req.Address,
		req.HealthInsuranceExpiredDate,
		req.HealthInsuranceNumber,
		req.EmergencyContactInfo,
	).Scan(&patientID)

	if err != nil {
		return err
	}

	return tx.Commit(ctx)
}

func GetPatientList(ctx context.Context) ([]models.PatientBriefInfo, error) {
	const query = `
		SELECT patient_id, full_name, date_of_birth, gender
		FROM patients
	`

	rows, _ := dbpool.Query(ctx, query)
	list, err := pgx.CollectRows(rows, pgx.RowToStructByName[models.PatientBriefInfo])
	if err != nil {
		return nil, err
	}
	return list, nil
}

// func GetPatientInfo(ctx context.Context, patientID string) (models.PatientInfo, error) {
// 	const query = `
// 		SELECT patient_id, full_name, date_of_birth, gender, ethnicity, nationality, address, health_insurance_expired_date, health_insurance_number, emergency_contact_info
// 		FROM patients
// 		WHERE patient_id = $1
// 	`

// 	rows, _ := dbpool.Query(ctx, query, patientID)
// 	info, err := pgx.CollectExactlyOneRow(rows, pgx.RowToStructByName[models.PatientInfo])
// 	if err != nil {
// 		return models.PatientInfo{}, err
// 	}
// 	return info, nil
// }
