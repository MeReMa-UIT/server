package comms_api

import "github.com/gin-gonic/gin"

func Routes(r *gin.RouterGroup) {
	r.GET("/contacts", GetContactListHandler)
}
