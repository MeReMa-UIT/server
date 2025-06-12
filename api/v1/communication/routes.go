package comm_api

import (
	"github.com/gin-gonic/gin"
	comm_services "github.com/merema-uit/server/services/communication"
)

func Routes(r *gin.RouterGroup) {
	chatHub := comm_services.NewHub()
	go chatHub.Run()

	var chatService comm_services.Service = chatHub
	wsHandler := NewWebSocketHandler(chatService)

	r.GET("/chat", wsHandler.ServeWSHandler)
}
