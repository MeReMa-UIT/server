package repo

import (
	"context"

	"github.com/jackc/pgx/v5"
	"github.com/merema-uit/server/models"
	errs "github.com/merema-uit/server/models/errors"
)

func GetPatientIDListByAccID(ctx context.Context, accID string) ([]int, error) {
	const query = `
		SELECT patient_id
		FROM patients
		WHERE acc_id = $1::BIGINT
	`

	var patientIDList []int
	rows, _ := dbpool.Query(ctx, query, accID)
	patientIDList, err := pgx.AppendRows(patientIDList, rows, pgx.RowTo[int])
	if err != nil {
		return nil, err
	}
	return patientIDList, nil
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

func GetPatientList(ctx context.Context, accID *string) ([]models.PatientBriefInfo, error) {
	const query = `
		SELECT patient_id, full_name, date_of_birth, gender
		FROM patients
		WHERE (acc_id = $1::BIGINT OR $1::BIGINT IS NULL)
	`

	rows, _ := dbpool.Query(ctx, query, accID)
	list, err := pgx.CollectRows(rows, pgx.RowToStructByName[models.PatientBriefInfo])
	if err != nil {
		return nil, err
	}
	return list, nil
}

func GetPatientInfo(ctx context.Context, patientID string, accID string) (models.PatientInfo, error) {
	const query = `
		SELECT patient_id, full_name, date_of_birth, gender, ethnicity, nationality, address, health_insurance_expired_date, health_insurance_number, emergency_contact_info
		FROM patients
		WHERE patient_id = $1 AND (acc_id = $2::BIGINT OR $2::BIGINT = 0)
	`

	rows, _ := dbpool.Query(ctx, query, patientID, accID)
	info, err := pgx.CollectExactlyOneRow(rows, pgx.RowToStructByName[models.PatientInfo])
	if err != nil {
		if err == pgx.ErrNoRows {
			return models.PatientInfo{}, errs.ErrPatientNotExist
		}
		return models.PatientInfo{}, err
	}
	return info, nil
}

func UpdatePatientInfo(ctx context.Context, patientID string, req models.PatientInfoUpdateRequest) error {
	const query = `
		UPDATE patients
		SET
				full_name = COALESCE(NULLIF($1, ''), full_name),
				date_of_birth = COALESCE(NULLIF($2, '0001-01-01 00:00:00'::timestamp), date_of_birth),
				gender = COALESCE(NULLIF($3, ''), gender),
				ethnicity = COALESCE(NULLIF($4, ''), ethnicity),
				nationality = COALESCE(NULLIF($5, ''), nationality),
				address = COALESCE(NULLIF($6, ''), address),
				health_insurance_expired_date = COALESCE(NULLIF($7, '0001-01-01 00:00:00'::timestamp), health_insurance_expired_date),
				health_insurance_number = COALESCE(NULLIF($8, ''), health_insurance_number),
				emergency_contact_info = COALESCE(NULLIF($9, ''), emergency_contact_info)
		WHERE patient_id = $10::BIGINT
		RETURNING patient_id
	`

	tx, err := dbpool.BeginTx(ctx, pgx.TxOptions{
		IsoLevel: pgx.Serializable,
	})
	if err != nil {
		return err
	}
	defer tx.Rollback(ctx)

	var updatedPatientID int
	err = tx.QueryRow(ctx, query,
		req.FullName,
		req.DateOfBirth,
		req.Gender,
		req.Ethnicity,
		req.Nationality,
		req.Address,
		req.HealthInsuranceExpiredDate,
		req.HealthInsuranceNumber,
		req.EmergencyContactInfo,
		patientID,
	).Scan(&updatedPatientID)

	if err != nil {
		if err == pgx.ErrNoRows {
			return errs.ErrPatientNotExist
		}
		return err
	}

	if err := tx.Commit(ctx); err != nil {
		return err
	}

	return nil
}
