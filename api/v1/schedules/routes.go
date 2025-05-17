package schedules_api

import "github.com/gin-gonic/gin"

func Routes(r *gin.RouterGroup) {
	r.GET("", GetScheduleListHandler)
	r.POST("/book", BookScheduleHandler)
}
