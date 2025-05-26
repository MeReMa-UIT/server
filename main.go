package main

import (
	"context"
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/merema-uit/server/api"
	"github.com/merema-uit/server/repo"

	"github.com/merema-uit/server/docs"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title MeReMa Server API
// @version 1.0
// @description API for medical records management app
// @host localhost:8080
// @BasePath /api/v1
// @license.name MIT
// @license.url https://opensource.org/licenses/MIT
// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
// @description Provide the JWT token as a header with format "Bearer \<token\>"
func main() {
	repo.ConnectToDB(context.Background(), os.Getenv("DB_URL"))
	defer repo.CloseDBConnect()

	docs.SwaggerInfo.BasePath = "/api/v1"
	r := gin.Default()
	api.RegisterRoutes(r)
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	fmt.Println("Swagger UI available at http://localhost:8080/swagger/index.html")

	r.Run(":8080")
}
