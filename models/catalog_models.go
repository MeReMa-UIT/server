package models

type MedicationInfo struct {
	MedicationID   int    `json:"med_id" db:"med_id"`
	Name           string `json:"name" db:"name"`
	GenericName    string `json:"generic_name" db:"generic_name"`
	MedicationType string `json:"med_type" db:"med_type"`
	Strength       string `json:"strength" db:"strength"`
	Manufacturer   string `json:"manufacturer" db:"manufacturer"`
}

type DiagnosisInfo struct {
	ICDCode     string  `json:"icd_code" db:"icd_code"`
	Name        string  `json:"name" db:"name"`
	Description *string `json:"description" db:"description"`
}
