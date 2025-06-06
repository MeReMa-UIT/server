package prescriptions_api

import "github.com/gin-gonic/gin"

func Routes(r *gin.RouterGroup) {
	r.POST("", AddNewPrescriptionHandler)
	r.GET("/patients/:patient_id", GetPrescriptionListWithPatientIDHandler)
	r.GET("/records/:record_id", GetPrescriptionListWithRecordIDHandler)
	r.GET("/:prescription_id", GetPrescriptionDetailsHandler)
	r.PUT("/:prescription_id", UpdatePrescriptionHandler)
	r.PUT("/:prescription_id/confirm", ConfirmReceivingPrescriptionHandler)
	r.POST("/:prescription_id/details", AddPrescriptionDetailHandler)
	// r.PUT("/:prescription_id/:detail_id", DeletePrescriptionDetailHandler)
	r.DELETE("/:prescription_id/:detail_id", DeletePrescriptionDetailHandler)
}
