package main

import (
	"context"

	"github.com/gin-gonic/gin"
	"github.com/merema-uit/server/api"
	"github.com/merema-uit/server/repo"
)

func main() {
	r := gin.Default()
	repo.ConnectToDB(context.Background(), repo.DATABASE_URL)
	defer repo.CloseDB()
	r.POST("/api/login", api.Login)
	r.POST("/api/register", api.Register)
	r.Run("0.0.0.0:8080")
}
