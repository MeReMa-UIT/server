package catalog_api

import "github.com/gin-gonic/gin"

func Routes(r *gin.RouterGroup) {
	r.GET("/medications", GetMedicationListHandler)
	r.GET("/medications/:medication_id", GetMedicationInfoHandler)
	r.GET("/diagnoses", GetDiagnosisListHandler)
	r.GET("/diagnoses/:icd_code", GetDiagnosisInfoHandler)
	r.GET("/record-types", GetMedicalRecordTypeListHandler)
	r.GET("/record-types/:type_id/template", GetMedicalRecordTemplateHandler)
}
