package models

import "github.com/jackc/pgtype"

type NewMedicalRecordRequest struct {
	PatientID    int          `json:"patient_id" `
	TypeID       string       `json:"type_id" `
	RecordDetail pgtype.JSONB `json:"record_detail" swaggertype:"object"`
}

type NewMedicalRecordResponse struct {
	RecordID int `json:"record_id"`
}

type ExtractedRecordInfo struct {
	PrimaryDiagnosis   string `json:"primary_diagnosis"`
	SecondaryDiagnosis string `json:"secondary_diagnosis"`
}
