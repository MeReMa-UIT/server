package repo

import (
	"context"

	"github.com/jackc/pgx/v5"
	"github.com/merema-uit/server/models"
)

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

func GetPrescriptionDetails(ctx context.Context, prescriptionID string) ([]models.PrescriptionDetail, error) {
	const query = `
		SELECT med_id, morning_dosage, afternoon_dosage, evening_dosage, duration_days, total_dosage, dosage_unit, instructions
		FROM prescription_details
		WHERE prescription_id = $1
	`
	rows, _ := dbpool.Query(ctx, query, prescriptionID)
	details, err := pgx.CollectRows(rows, pgx.RowToStructByName[models.PrescriptionDetail])
	if err != nil {
		return nil, err
	}
	return details, nil
}
