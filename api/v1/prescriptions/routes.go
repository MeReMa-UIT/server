package prescriptions_api

import "github.com/gin-gonic/gin"

func Routes(r *gin.RouterGroup) {
	r.POST("/new", AddNewPrescriptionHandler)
	r.GET("/patients/:patient_id", GetPrescriptionListWithPatientIDHandler)
	r.GET("/records/:record_id", GetPrescriptionListWithRecordIDHandler)
	r.GET("/:prescription_id", GetPrescriptionDetailsHandler)
}
