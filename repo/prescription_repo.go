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

func GetPrescriptionIDListWithAccID(ctx context.Context, accID string) ([]int, error) {
	const query = `
		SELECT prescription_id
		FROM prescriptions p JOIN records r ON p.record_id = r.record_id
		JOIN patients pa ON r.patient_id = pa.patient_id
		WHERE pa.acc_id = $1::BIGINT
	`
	var prescriptionIDList []int
	rows, _ := dbpool.Query(ctx, query, accID)
	prescriptionIDList, err := pgx.AppendRows(prescriptionIDList, rows, pgx.RowTo[int])
	if err != nil {
		return nil, err
	}
	return prescriptionIDList, nil
}

func StorePrescription(ctx context.Context, req models.NewPrescriptionRequest) error {
	const query = `
		INSERT INTO prescriptions (record_id, is_insurance_covered, prescription_note)
		VALUES ($1, $2, $3)
		RETURNING prescription_id
	`
	tx, err := dbpool.BeginTx(ctx, pgx.TxOptions{
		IsoLevel: pgx.Serializable,
	})
	if err != nil {
		return err
	}
	defer tx.Rollback(ctx)

	var prescriptionID int
	err = tx.QueryRow(ctx, query, req.RecordID, req.IsInsuranceCovered, req.PrescriptionNote).Scan(&prescriptionID)
	if err != nil {
		return err
	}

	_, err = tx.CopyFrom(
		context.Background(),
		pgx.Identifier{"prescription_details"},
		[]string{"prescription_id", "med_id", "morning_dosage", "afternoon_dosage", "evening_dosage", "total_dosage", "duration_days", "dosage_unit", "instructions"},
		pgx.CopyFromSlice(len(req.Details), func(i int) ([]any, error) {
			return []any{
				prescriptionID,
				req.Details[i].MedicationID,
				req.Details[i].MorningDosage,
				req.Details[i].AfternoonDosage,
				req.Details[i].EveningDosage,
				req.Details[i].TotalDosage,
				req.Details[i].DurationDays,
				req.Details[i].DosageUnit,
				req.Details[i].Instructions,
			}, nil
		}),
	)

	if err != nil {
		return err
	}

	if err := tx.Commit(ctx); err != nil {
		return err
	}

	return nil
}

func GetPrescriptionListWithRecordID(ctx context.Context, recordID string) ([]models.PrescriptionInfo, error) {
	const query = `
		SELECT prescription_id, record_id, is_insurance_covered, prescription_note, created_at, received_at
		FROM prescriptions
		WHERE record_id = $1
	`
	rows, _ := dbpool.Query(ctx, query, recordID)
	list, err := pgx.CollectRows(rows, pgx.RowToStructByName[models.PrescriptionInfo])
	if err != nil {
		return nil, err
	}
	return list, nil
}

func GetPrescriptionListWithPatientID(ctx context.Context, patientID string) ([]models.PrescriptionInfo, error) {
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

func GetPrescriptionDetails(ctx context.Context, prescriptionID string) ([]models.PrescriptionDetailInfo, error) {
	const query = `
		SELECT detail_id, med_id, morning_dosage, afternoon_dosage, evening_dosage, duration_days, total_dosage, dosage_unit, instructions
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

	// for _, detail := range req.Details {
	// 	const query2 = `
	// 		UPDATE prescription_details
	// 		SET
	// 			med_id = $1,
	// 			morning_dosage = $2,
	// 			afternoon_dosage = $3,
	// 			evening_dosage = $4,
	// 			duration_days = $5,
	// 			total_dosage = $6,
	// 			dosage_unit = $7,
	// 			instructions = $8
	// 		WHERE detail_id = $9 AND prescription_id = $10
	// 	`

	// 	_, err = tx.Exec(ctx, query2,
	// 		detail.MedicationID,
	// 		detail.MorningDosage,
	// 		detail.AfternoonDosage,
	// 		detail.EveningDosage,
	// 		detail.DurationDays,
	// 		detail.TotalDosage,
	// 		detail.DosageUnit,
	// 		detail.Instructions,
	// 		detail.DetailID,
	// 		prescriptionID,
	// 	)

	// 	if err != nil {
	// 		return err
	// 	}
	// }

	// A patch to skip updating, overriding instead

	const query2 = `
		DELETE FROM prescription_details
		WHERE prescription_id = $1	
	`

	_, err = tx.Exec(ctx, query2, prescriptionID)

	if err != nil {
		return err
	}

	_, err = tx.CopyFrom(
		context.Background(),
		pgx.Identifier{"prescription_details"},
		[]string{"prescription_id", "med_id", "morning_dosage", "afternoon_dosage", "evening_dosage", "total_dosage", "duration_days", "dosage_unit", "instructions"},
		pgx.CopyFromSlice(len(req.Details), func(i int) ([]any, error) {
			return []any{
				prescriptionID,
				req.Details[i].MedicationID,
				req.Details[i].MorningDosage,
				req.Details[i].AfternoonDosage,
				req.Details[i].EveningDosage,
				req.Details[i].TotalDosage,
				req.Details[i].DurationDays,
				req.Details[i].DosageUnit,
				req.Details[i].Instructions,
			}, nil
		}),
	)

	if err != nil {
		return err
	}
	// end patch

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

func AddPrescriptionDetail(ctx context.Context, prescriptionID string, details []models.PrescriptionDetail) error {

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

func DeletePrescriptionDetail(ctx context.Context, prescriptionID, detailID string) error {
	const query = `
		DELETE FROM prescription_details
		WHERE prescription_id = $1 AND detail_id = $2
	`

	tx, err := dbpool.BeginTx(ctx, pgx.TxOptions{
		IsoLevel: pgx.Serializable,
	})
	if err != nil {
		return err
	}
	defer tx.Rollback(ctx)

	result, err := tx.Exec(ctx, query, prescriptionID, detailID)
	if err != nil {
		return err
	}

	if result.RowsAffected() == 0 {
		return errs.ErrPrescriptionNotFound
	}

	if err := tx.Commit(ctx); err != nil {
		return err
	}

	return nil
}
