package models

import "time"

type RecordInfoForStatistic struct {
	PatientID        int64     `json:"patient_id"`
	DoctorID         int64     `json:"doctor_id"`
	PrimaryDiagnosis string    `json:"primary_diagnosis"`
	CreatedAt        time.Time `json:"created_at"`
}

type RecordStatisticRequest struct {
	TimeUnit string `json:"time_unit"` // "day", "week", "month", "year"
	/*
		This use for determine the period of time we want to compile, must be time in RFC3339 format.
		If time unit is "day", the d-m-y you provide in timestamp will determine that exactly date (e.g, the d-m-y itself).
		If time unit is "week", the d-m-y you provide in timestamp will determine the week which that date belong to (e.g, determine the nearest Monday and nearest Sunday which that date sit in between).
		If time unit is "month", the d-m-y you provide in timestamp will determine the month which that date belong to (e.g, the month m in year y).
		If time unit is "week", the d-m-y you provide in timestamp will determine the year which that date belong to (e.g, the year y).
	*/
	Timestamp time.Time `json:"timestamp"`
}

type AmountOfRecordsByTime struct {
	TimestampStart time.Time `json:"timestamp_start"`
	Amount         int       `json:"amount"`
}

type AmountOfRecordsByDoctor struct {
	DoctorID     int64                   `json:"doctor_id"`
	AmountByTime []AmountOfRecordsByTime `json:"amount_by_time"`
	TotalAmount  int                     `json:"total_amount"`
}

type AmountOfRecordsByDiagnosis struct {
	DiagnosisID  string                  `json:"diagnosis_id"`
	AmountByTime []AmountOfRecordsByTime `json:"amount_by_time"`
	TotalAmount  int                     `json:"total_amount"`
}
