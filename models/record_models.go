package models

import "github.com/jackc/pgtype"

type NewMedicalRecordRequest struct {
	PatientID int          `json:"patient_id" `
	Details   pgtype.JSONB `json:"details" swaggertype:"object"`
}
