package comm_api

import (
	"github.com/gin-gonic/gin"
)

func Routes(r *gin.RouterGroup) {
	r.GET("/contacts", GetContactListHandler)
	r.POST("/messages", SendMessageHandler)
	r.GET("/messages/:contact_id", LoadConversationHandler)
}
