package api_v1

import (
	"github.com/gin-gonic/gin"
	accounts_api "github.com/merema-uit/server/api/v1/accounts"
	catalogs_api "github.com/merema-uit/server/api/v1/catalog"
	comms_api "github.com/merema-uit/server/api/v1/communications"
	patients_api "github.com/merema-uit/server/api/v1/patients"
	prescriptions_api "github.com/merema-uit/server/api/v1/prescriptions"
	records_api "github.com/merema-uit/server/api/v1/records"
	schedules_api "github.com/merema-uit/server/api/v1/schedules"
	staffs_api "github.com/merema-uit/server/api/v1/staffs"
)

func RegisterRoutesV1(r *gin.Engine) {
	v1 := r.Group("/api/v1")
	{
		accounts_api.Routes(v1.Group("/accounts"))
		patients_api.Routes(v1.Group("/patients"))
		staffs_api.Routes(v1.Group("/staffs"))
		catalogs_api.Routes(v1.Group("/catalog"))
		records_api.Routes(v1.Group("/records"))
		schedules_api.Routes(v1.Group("/schedules"))
		prescriptions_api.Routes(v1.Group("/prescriptions"))
		comms_api.Routes(v1.Group("/comms"))
	}
}
