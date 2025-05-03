package main

import (
	"context"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/merema-uit/server/api"
	"github.com/merema-uit/server/repo"
)

func main() {
	r := gin.Default()
	repo.ConnectToDB(context.Background(), os.Getenv("DB_URL"))
	defer repo.CloseDB()
	api.RegisterRoutesV1(r)
	r.Run(":8080")
}
