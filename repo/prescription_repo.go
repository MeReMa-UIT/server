package repo

import (
	"context"

	"github.com/jackc/pgx/v5"
	"github.com/merema-uit/server/models"
)

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
