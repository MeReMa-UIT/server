package models

import "time"

type PrescriptionDetail struct {
	MorningDosage   float32 `json:"morning_dosage" db:"morning_dosage"`
	AfternoonDosage float32 `json:"afternoon_dosage" db:"afternoon_dosage"`
	EveningDosage   float32 `json:"evening_dosage" db:"evening_dosage"`
	DurationDays    int     `json:"duration_days" db:"duration_days"`
	TotalDosage     float32 `json:"total_dosage" db:"total_dosage"`
	DosageUnit      string  `json:"dosage_unit" db:"dosage_unit"`
	Instructions    string  `json:"instructions" db:"instructions"`
}

type PrescriptionDetailInfo struct {
	MedicationID int `json:"med_id" db:"med_id"`
	PrescriptionDetail
}

type NewPrescriptionRequest struct {
	RecordID           int    `json:"record_id" db:"record_id"`
	IsInsuranceCovered bool   `json:"is_insurance_covered" db:"is_insurance_covered"`
	PrescriptionNote   string `json:"prescription_note" db:"prescription_note"`
}

type NewPrescriptionResponse struct {
	PrescriptionID int `json:"prescription_id" db:"prescription_id"`
}

type PrescriptionInfo struct {
	PrescriptionID     int        `json:"prescription_id" db:"prescription_id"`
	RecordID           int        `json:"record_id" db:"record_id"`
	IsInsuranceCovered bool       `json:"is_insurance_covered" db:"is_insurance_covered"`
	PrescriptionNote   string     `json:"prescription_note" db:"prescription_note"`
	CreatedAt          time.Time  `json:"created_at" db:"created_at"`
	ReceivedAt         *time.Time `json:"received_at" db:"received_at"`
}

type PrescriptionUpdateRequest struct {
	IsInsuranceCovered bool   `json:"is_insurance_covered" db:"is_insurance_covered"`
	PrescriptionNote   string `json:"prescription_note" db:"prescription_note"`
}
