package accounts_api

import (
	"github.com/gin-gonic/gin"
)

func Routes(r *gin.RouterGroup) {
	r.POST("/login", LoginHandler)
	r.POST("/recovery", RecoveryHandler)
	r.POST("/recovery/verify", RecoveryConfirmHandler)
	r.PUT("/recovery/reset", ResetPasswordHandler)
	r.POST("/register", InitRegistrationHandler)
	r.POST("/register/create", RegisterAccountHandler)
	r.POST("/register/patients", RegisterPatientHandler)
	r.POST("/register/staffs", RegisterStaffHandler)
	r.GET("/profile", GetAccountInfoHandler)
}
