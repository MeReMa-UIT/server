package repo

import (
	"context"

	"github.com/jackc/pgx/v5"
	"github.com/merema-uit/server/models"
	errs "github.com/merema-uit/server/models/errors"
)

func CheckPrescriptionExists(ctx context.Context, prescriptionID string) error {
	const query = `
		SELECT EXISTS (
			SELECT 1
			FROM prescriptions
			WHERE prescription_id = $1
		)
	`
	var exists bool
	err := dbpool.QueryRow(ctx, query, prescriptionID).Scan(&exists)
	if err != nil {
		return err
	}
	if !exists {
		return errs.ErrPrescriptionNotFound
	}
	return nil
}

func GetPrescriptionIDListByAccID(ctx context.Context, accID string) ([]int, error) {
	const query = `
		SELECT prescription_id
		FROM prescriptions
		WHERE record_id = ANY($1)
	`

	recordIDList, err := GetRecordIDListByAccID(ctx, accID)

	if err != nil {
		return nil, err
	}

	rows, _ := dbpool.Query(ctx, query, recordIDList)
	prescriptionIDList, err := pgx.CollectRows(rows, pgx.RowTo[int])
	if err != nil {
		return nil, err
	}
	return prescriptionIDList, nil
}

func GetPrescriptionList(ctx context.Context, recordIDList []int) ([]models.PrescriptionInfo, error) {
	const query = `
		SELECT prescription_id, record_id, is_insurance_covered, prescription_note, created_at, received_at
		FROM prescriptions 
		WHERE record_id = ANY($1) OR $1::BIGINT[] IS NULL
	`
	rows, _ := dbpool.Query(ctx, query, recordIDList)
	list, err := pgx.CollectRows(rows, pgx.RowToStructByName[models.PrescriptionInfo])
	if err != nil {
		return nil, err
	}
	return list, nil
}

func GetPrescriptionListByPatientID(ctx context.Context, patientID string) ([]models.PrescriptionInfo, error) {
	const query = `
		SELECT prescription_id, p.record_id, is_insurance_covered, prescription_note, p.created_at, received_at
		FROM prescriptions p JOIN records r ON p.record_id = r.record_id
		WHERE r.patient_id = $1
	`
	rows, _ := dbpool.Query(ctx, query, patientID)
	list, err := pgx.CollectRows(rows, pgx.RowToStructByName[models.PrescriptionInfo])
	if err != nil {
		return nil, err
	}
	return list, nil
}

func GetPrescriptionInfoByRecordID(ctx context.Context, recordID int) (models.PrescriptionInfo, error) {
	const query = `
		SELECT prescription_id, record_id, is_insurance_covered, prescription_note, created_at, received_at
		FROM prescriptions 
		WHERE record_id = $1
	`
	rows, _ := dbpool.Query(ctx, query, recordID)
	info, err := pgx.CollectExactlyOneRow(rows, pgx.RowToStructByName[models.PrescriptionInfo])
	if err != nil {
		return models.PrescriptionInfo{}, err
	}
	return info, nil
}

func GetPrescriptionDetails(ctx context.Context, prescriptionID string) ([]models.PrescriptionDetailInfo, error) {
	const query = `
		SELECT med_id, morning_dosage, afternoon_dosage, evening_dosage, duration_days, total_dosage, dosage_unit, instructions
		FROM prescription_details
		WHERE prescription_id = $1
	`
	rows, _ := dbpool.Query(ctx, query, prescriptionID)
	details, err := pgx.CollectRows(rows, pgx.RowToStructByName[models.PrescriptionDetailInfo])
	if err != nil {
		return nil, err
	}
	return details, nil
}

func StorePrescription(ctx context.Context, req models.NewPrescriptionRequest) (models.NewPrescriptionResponse, error) {
	const query = `
		INSERT INTO prescriptions (record_id, is_insurance_covered, prescription_note)
		VALUES ($1, $2, $3)
		RETURNING prescription_id
	`
	tx, err := dbpool.BeginTx(ctx, pgx.TxOptions{
		IsoLevel: pgx.Serializable,
	})
	if err != nil {
		return models.NewPrescriptionResponse{}, err
	}
	defer tx.Rollback(ctx)

	var createdPrescriptionID int
	err = tx.QueryRow(ctx, query, req.RecordID, req.IsInsuranceCovered, req.PrescriptionNote).Scan(&createdPrescriptionID)
	if err != nil {
		return models.NewPrescriptionResponse{}, err
	}

	if err := tx.Commit(ctx); err != nil {
		return models.NewPrescriptionResponse{}, err
	}

	return models.NewPrescriptionResponse{PrescriptionID: createdPrescriptionID}, nil
}

func StorePrescriptionDetails(ctx context.Context, prescriptionID string, details []models.PrescriptionDetailInfo) error {
	tx, err := dbpool.BeginTx(ctx, pgx.TxOptions{
		IsoLevel: pgx.Serializable,
	})
	if err != nil {
		return err
	}
	defer tx.Rollback(ctx)

	_, err = tx.CopyFrom(
		context.Background(),
		pgx.Identifier{"prescription_details"},
		[]string{"prescription_id", "med_id", "morning_dosage", "afternoon_dosage", "evening_dosage", "total_dosage", "duration_days", "dosage_unit", "instructions"},
		pgx.CopyFromSlice(len(details), func(i int) ([]any, error) {
			return []any{
				prescriptionID,
				details[i].MedicationID,
				details[i].MorningDosage,
				details[i].AfternoonDosage,
				details[i].EveningDosage,
				details[i].TotalDosage,
				details[i].DurationDays,
				details[i].DosageUnit,
				details[i].Instructions,
			}, nil
		}),
	)

	if err != nil {
		return err
	}

	if err := tx.Commit(ctx); err != nil {
		return err
	}

	return err
}

func UpdatePrescription(ctx context.Context, prescriptionID string, req models.PrescriptionUpdateRequest) error {
	const query = `
		UPDATE prescriptions
		SET 
			is_insurance_covered = $1, 
			prescription_note = $2
		WHERE prescription_id = $3 AND received_at IS NULL
		RETURNING prescription_id
	`
	if err := CheckPrescriptionExists(ctx, prescriptionID); err != nil {
		return err
	}

	tx, err := dbpool.BeginTx(ctx, pgx.TxOptions{
		IsoLevel: pgx.Serializable,
	})
	if err != nil {
		return err
	}
	defer tx.Rollback(ctx)

	var updatedPrescriptionID int
	err = tx.QueryRow(ctx, query, req.IsInsuranceCovered, req.PrescriptionNote, prescriptionID).Scan(&updatedPrescriptionID)
	if err != nil {
		if err == pgx.ErrNoRows {
			return errs.ErrReceivedPrescription
		}
		return err
	}

	if err := tx.Commit(ctx); err != nil {
		return err
	}

	return nil
}

func ComfirmReceivingPrescription(ctx context.Context, prescriptionID string) error {
	const query = `
		UPDATE prescriptions
		SET received_at = NOW()
		WHERE prescription_id = $1 AND received_at IS NULL
		RETURNING prescription_id
	`

	if err := CheckPrescriptionExists(ctx, prescriptionID); err != nil {
		return err
	}

	tx, err := dbpool.BeginTx(ctx, pgx.TxOptions{
		IsoLevel: pgx.Serializable,
	})
	if err != nil {
		return err
	}
	defer tx.Rollback(ctx)

	result, err := tx.Exec(ctx, query, prescriptionID)
	if err != nil {
		return err
	}

	if result.RowsAffected() == 0 {
		return errs.ErrReceivedPrescription
	}

	if err := tx.Commit(ctx); err != nil {
		return err
	}

	return nil
}

func UpdatePrescriptionDetail(ctx context.Context, prescriptionID, medID string, detail models.PrescriptionDetailInfo) error {
	const query = `
		UPDATE prescription_details
		SET 
			morning_dosage = $1, 
			afternoon_dosage = $2, 
			evening_dosage = $3, 
			total_dosage = $4, 
			duration_days = $5, 
			dosage_unit = $6, 
			instructions = $7
		WHERE prescription_id = $8 AND med_id = $9
	`

	if err := CheckPrescriptionExists(ctx, prescriptionID); err != nil {
		return err
	}

	tx, err := dbpool.BeginTx(ctx, pgx.TxOptions{
		IsoLevel: pgx.Serializable,
	})
	if err != nil {
		return err
	}
	defer tx.Rollback(ctx)

	result, err := tx.Exec(ctx, query,
		detail.MorningDosage,
		detail.AfternoonDosage,
		detail.EveningDosage,
		detail.TotalDosage,
		detail.DurationDays,
		detail.DosageUnit,
		detail.Instructions,
		prescriptionID,
		medID,
	)
	if err != nil {
		return err
	}

	if result.RowsAffected() == 0 {
		return errs.ErrPrescriptionDetailNotFound
	}

	if err := tx.Commit(ctx); err != nil {
		return err
	}

	return nil
}

func DeletePrescriptionDetail(ctx context.Context, prescriptionID, medID string) error {
	const query = `
		DELETE FROM prescription_details
		WHERE prescription_id = $1 AND med_id = $2
	`

	tx, err := dbpool.BeginTx(ctx, pgx.TxOptions{
		IsoLevel: pgx.Serializable,
	})
	if err != nil {
		return err
	}
	defer tx.Rollback(ctx)

	result, err := tx.Exec(ctx, query, prescriptionID, medID)
	if err != nil {
		return err
	}

	if result.RowsAffected() == 0 {
		return errs.ErrPrescriptionDetailNotFound
	}

	if err := tx.Commit(ctx); err != nil {
		return err
	}

	return nil
}
