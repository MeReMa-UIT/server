package statistic_api

import "github.com/gin-gonic/gin"

func Routes(r *gin.RouterGroup) {
	r.POST("/records", CompileRecordStatisticByTimeHandler)
	r.POST("/records/doctor", CompileRecordStatisticByDoctorHandler)
	r.POST("/records/diagnosis", CompileRecordStatisticByDiagnosisHandler)
}
