package models

import "time"

type PatientInfo struct {
	PatientID                  int       `json:"patient_id" db:"patient_id"`
	FullName                   string    `json:"full_name" db:"full_name"`
	DateOfBirth                time.Time `json:"date_of_birth" db:"date_of_birth"`
	Gender                     string    `json:"gender" db:"gender"`
	Ethnicity                  string    `json:"ethnicity" db:"ethnicity"`
	Nationality                string    `json:"nationality" db:"nationality"`
	Address                    string    `json:"address" db:"address"`
	HealthInsuranceExpiredDate time.Time `json:"health_insurance_expired_date" db:"health_insurance_expired_date"`
	HealthInsuranceNumber      string    `json:"health_insurance_number" db:"health_insurance_number"`
	EmergencyContactInfo       string    `json:"emergency_contact_info" db:"emergency_contact_info"`
}

type PatientBriefInfo struct {
	PatientID   int       `json:"patient_id" db:"patient_id"`
	FullName    string    `json:"full_name" db:"full_name"`
	DateOfBirth time.Time `json:"date_of_birth" db:"date_of_birth"`
	Gender      string    `json:"gender" db:"gender"`
}
