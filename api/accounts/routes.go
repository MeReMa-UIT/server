package accounts_api

import (
	"github.com/gin-gonic/gin"
)

func Routes(r *gin.RouterGroup) {
	r.POST("/login", LoginHandler)
	r.POST("/register", RegisterHandler)
	r.POST("/recovery", RecoveryHandler)
	r.POST("/recovery/verify", RecoveryConfirmHandler)
	r.PUT("/recovery/reset", ResetPasswordHandler)
}
