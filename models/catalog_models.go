package models

type MedicationInfo struct {
	MedicationID          int    `json:"med_id" db:"med_id"`
	Name                  string `json:"name" db:"name"`
	GenericName           string `json:"generic_name" db:"generic_name"`
	MedicationType        string `json:"med_type" db:"med_type"`
	Strength              string `json:"strength" db:"strength"`
	RouteOfAdministration string `json:"route_of_administration" db:"route_of_administration"`
	Manufacturer          string `json:"manufacturer" db:"manufacturer"`
}

type DiagnosisInfo struct {
	ICDCode     string  `json:"icd_code" db:"icd_code"`
	Name        string  `json:"name" db:"name"`
	Description *string `json:"description" db:"description"`
}

type MedicalRecordType struct {
	TypeID   string `json:"type_id" db:"type_id"`
	TypeName string `json:"type_name" db:"type_name"`
}

type MedicalRecordTypeInfo struct {
	MedicalRecordType
	Description  *string `json:"description" db:"description"`
	TemplatePath string  `json:"template_path" db:"template_path"`
	SchemaPath   string  `json:"schema_path" db:"schema_path"`
}
