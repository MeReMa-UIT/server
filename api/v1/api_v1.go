package api_v1

import (
	"github.com/gin-gonic/gin"
	accounts_api "github.com/merema-uit/server/api/v1/accounts"
)

func RegisterRoutesV1(r *gin.Engine) {
	v1 := r.Group("/api/v1")
	{
		accounts_api.Routes(v1.Group("/accounts"))
	}
}
