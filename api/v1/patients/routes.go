package patients_api

import "github.com/gin-gonic/gin"

func Routes(r *gin.RouterGroup) {
	r.GET("", GetPatientList)
	r.GET("/:patient_id", GetPatientInfoHandler)
}
