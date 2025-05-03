package api

import (
	"github.com/gin-gonic/gin"
	accounts_api "github.com/merema-uit/server/api/accounts"
)

func RegisterRoutesV1(r *gin.Engine) {
	v1 := r.Group("/api/v1")
	{
		accounts_api.Routes(v1.Group("/accounts"))
	}
}
