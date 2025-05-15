package api_v1

import (
	"github.com/gin-gonic/gin"
	accounts_api "github.com/merema-uit/server/api/v1/accounts"
	catalog_api "github.com/merema-uit/server/api/v1/catalog"
	patients_api "github.com/merema-uit/server/api/v1/patients"
	records_api "github.com/merema-uit/server/api/v1/records"
	staffs_api "github.com/merema-uit/server/api/v1/staffs"
)

func RegisterRoutesV1(r *gin.Engine) {
	v1 := r.Group("/api/v1")
	{
		accounts_api.Routes(v1.Group("/accounts"))
		patients_api.Routes(v1.Group("/patients"))
		staffs_api.Routes(v1.Group("/staffs"))
		catalog_api.Routes(v1.Group("/catalog"))
		records_api.Routes(v1.Group("/records"))
	}
}
