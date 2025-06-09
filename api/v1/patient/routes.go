package patient_api

import "github.com/gin-gonic/gin"

func Routes(r *gin.RouterGroup) {
	r.GET("", GetPatientListHandler)
	r.GET("/:patient_id", GetPatientInfoHandler)
	r.PUT("/:patient_id", UpdatePatientInfoHandler)
}
