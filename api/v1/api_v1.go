package api_v1

import (
	"github.com/gin-gonic/gin"
	account_api "github.com/merema-uit/server/api/v1/account"
	catalog_api "github.com/merema-uit/server/api/v1/catalog"
	comm_api "github.com/merema-uit/server/api/v1/communication"
	patient_api "github.com/merema-uit/server/api/v1/patient"
	prescription_api "github.com/merema-uit/server/api/v1/prescription"
	record_api "github.com/merema-uit/server/api/v1/record"
	schedule_api "github.com/merema-uit/server/api/v1/schedule"
	staff_api "github.com/merema-uit/server/api/v1/staff"
)

func RegisterRoutesV1(r *gin.Engine) {
	v1 := r.Group("/api/v1")
	{
		account_api.Routes(v1.Group("/accounts"))
		patient_api.Routes(v1.Group("/patients"))
		staff_api.Routes(v1.Group("/staffs"))
		catalog_api.Routes(v1.Group("/catalog"))
		record_api.Routes(v1.Group("/records"))
		prescription_api.Routes(v1.Group("/prescriptions"))
		schedule_api.Routes(v1.Group("/schedules"))
		comm_api.Routes(v1.Group("/comms"))
	}
}
