package staff_api

import "github.com/gin-gonic/gin"

func Routes(r *gin.RouterGroup) {
	r.GET("", GetStaffListHandler)
	r.GET("/:staff_id", GetStaffInfoHandler)
	r.PUT("/:staff_id", UpdateStaffInfoHandler)
}
