package models

import (
	"time"

	"github.com/jackc/pgtype"
)

type NewMedicalRecordRequest struct {
	PatientID    int         `json:"patient_id" `
	TypeID       string      `json:"type_id" `
	RecordDetail pgtype.JSON `json:"record_detail" swaggertype:"object"`
}

type NewMedicalRecordResponse struct {
	RecordID int `json:"record_id"`
}

type ExtractedRecordInfo struct {
	PrimaryDiagnosis   string `json:"primary_diagnosis"`
	SecondaryDiagnosis string `json:"secondary_diagnosis"`
}

type MedicalRecordInfo struct {
	RecordID     int         `json:"record_id" db:"record_id"`
	PatientID    int         `json:"patient_id" db:"patient_id"`
	DoctorID     int         `json:"doctor_id" db:"doctor_id"`
	TypeID       string      `json:"type_id" db:"type_id"`
	RecordDetail pgtype.JSON `json:"record_detail" db:"record_detail" swaggertype:"object"`
	CreatedAt    time.Time   `json:"created_at" db:"created_at"`
	ExpiredAt    time.Time   `json:"expired_at" db:"expired_at"`
}

type MedicalRecordBriefInfo struct {
	RecordID           int     `json:"record_id" db:"record_id"`
	PatientID          int     `json:"patient_id" db:"patient_id"`
	DoctorID           int     `json:"doctor_id" db:"doctor_id"`
	TypeID             string  `json:"type_id" db:"type_id"`
	PrimaryDiagnosis   string  `json:"primary_diagnosis" db:"primary_diagnosis"`
	SecondaryDiagnosis *string `json:"secondary_diagnosis" db:"secondary_diagnosis"`
}

type UpdateMedicalRecordRequest struct {
	NewRecordDetail pgtype.JSON `json:"new_record_detail" swaggertype:"object"`
}

type DeleteRecordAttachmentRequest struct {
	AttachmentFileName string `json:"attachment_file_name"`
}
