package api_v1

import (
	"github.com/gin-gonic/gin"
	accounts_api "github.com/merema-uit/server/api/v1/accounts"
	patients_api "github.com/merema-uit/server/api/v1/patients"
	staffs_api "github.com/merema-uit/server/api/v1/staffs"
)

func RegisterRoutesV1(r *gin.Engine) {
	v1 := r.Group("/api/v1")
	{
		accounts_api.Routes(v1.Group("/accounts"))
		patients_api.Routes(v1.Group("/patients"))
		staffs_api.Routes(v1.Group("/staffs"))
	}
}
