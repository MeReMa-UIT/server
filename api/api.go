package api

import (
	"github.com/gin-gonic/gin"
	api_v1 "github.com/merema-uit/server/api/v1"
)

func RegisterRoutes(r *gin.Engine) {
	api_v1.RegisterRoutesV1(r)
}
