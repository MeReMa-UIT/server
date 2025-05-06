package repo

import (
	"context"

	"github.com/merema-uit/server/models"
)

func StorePatientInfo(ctx context.Context, req models.PatientRegisterRequest, accID int) error {
	patientTableLock.Lock()
	defer patientTableLock.Unlock()

	const query = `
		INSERT INTO patients (acc_id, full_name, date_of_birth, gender, ethnicity, nationality, address, health_insurance_expired_date, health_insurance_number, emergency_contact_info)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10)
		RETURNING patient_id
	`

	var patientID int
	err := dbpool.QueryRow(ctx, query,
		accID,
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

	return err
}
