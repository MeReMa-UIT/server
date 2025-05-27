package catalogs_api

import "github.com/gin-gonic/gin"

func Routes(r *gin.RouterGroup) {
	r.GET("/medications", GetMedicationListHandler)
	r.GET("/diagnoses", GetDiagnosisListHandler)
}
