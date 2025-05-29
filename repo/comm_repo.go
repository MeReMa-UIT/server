package repo

import (
	"context"

	"github.com/jackc/pgx/v5"
	"github.com/merema-uit/server/models"
)

func GetContactListDoctor(ctx context.Context, accID string) ([]models.ContactInfo, error) {
	const query = `
		SELECT DISTINCT p.acc_id, p.full_name
		FROM records r JOIN patients p ON r.patient_id = p.patient_id JOIN staffs s ON r.doctor_id = s.staff_id
		WHERE s.acc_id = $1
	`

	rows, _ := dbpool.Query(ctx, query, accID)
	contacts, err := pgx.CollectRows(rows, pgx.RowToStructByName[models.ContactInfo])
	if err != nil {
		return nil, err
	}
	return contacts, nil
}

func GetContactListPatient(ctx context.Context, accID string) ([]models.ContactInfo, error) {
	const query = `
		SELECT DISTINCT s.acc_id, s.full_name
		FROM records r JOIN patients p ON r.patient_id = p.patient_id JOIN staffs s ON r.doctor_id = s.staff_id
		WHERE p.acc_id = $1
	`

	rows, _ := dbpool.Query(ctx, query, accID)
	contacts, err := pgx.CollectRows(rows, pgx.RowToStructByName[models.ContactInfo])
	if err != nil {
		return nil, err
	}
	return contacts, nil
}
