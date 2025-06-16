package models

import "time"

type RecordInfoForStatistic struct {
	PatientID        int64     `json:"patient_id"`
	DoctorID         int64     `json:"doctor_id"`
	PrimaryDiagnosis string    `json:"primary_diagnosis"`
	CreatedAt        time.Time `json:"created_at"`
}

type RecordStatisticRequest struct {
	TimeUnit  string    `json:"time_unit"` // "day", "week", "month", "year"
	Timestamp time.Time `json:"timestamp"`
}

type AmountOfRecordsByTime struct {
	TimestampStart time.Time `json:"timestamp_start"`
	Amount         int       `json:"amount"`
}

type AmountOfRecordsByDoctor struct {
	DoctorID     int64                   `json:"doctor_id"`
	AmountByTime []AmountOfRecordsByTime `json:"amount_by_time"`
}

type AmountOfRecordsByDiagnosis struct {
	DiagnosisID  string                  `json:"diagnosis_id"`
	AmountByTime []AmountOfRecordsByTime `json:"amount_by_time"`
}
