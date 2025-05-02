package main

import (
	"github.com/gin-gonic/gin"
	"github.com/merema-uit/server/api"
)

func main() {
	r := gin.Default()
	r.POST("/api/login", api.Login)
	r.POST("/api/register", api.Register)
	r.Run("0.0.0.0:8080")
}
