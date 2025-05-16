package schedules_api

import "github.com/gin-gonic/gin"

func Routes(r *gin.RouterGroup) {
	r.POST("/book", BookScheduleHandler)
}
