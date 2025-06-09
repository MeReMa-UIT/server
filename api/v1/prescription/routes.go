package prescription_api

import "github.com/gin-gonic/gin"

func Routes(r *gin.RouterGroup) {
	r.GET("", GetPrescriptionListHandler)
	r.POST("", AddNewPrescriptionHandler)
	r.GET("/patients/:patient_id", GetPrescriptionListByPatientIDHandler)
	r.GET("/records/:record_id", GetPrescriptionInfoByRecordIDHandler)
	r.GET("/:prescription_id", GetPrescriptionDetailsHandler)
	r.POST("/:prescription_id", AddPrescriptionDetailsHandler)
	r.PUT("/:prescription_id", UpdatePrescriptionHandler)
	r.PUT("/:prescription_id/confirm", ConfirmReceivingPrescriptionHandler)
	r.PUT("/:prescription_id/:med_id", UpdatePrescriptionDetailHandler)
	r.DELETE("/:prescription_id/:med_id", DeletePrescriptionDetailHandler)
}
