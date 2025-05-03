package accounts_api

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/merema-uit/server/models"
	"github.com/merema-uit/server/services/recovery"
)

// Recovery godoc
// @Summary Get important info to send recovery email
// @Description Send recovery email to user
// @Tags accounts
// @Accept json
// @Produce json
// @Param credentials body models.AccountRecoverRequest true "Recovery credentials"
// @Success 200 {object} map[string]string
// @Failure 401 {object} map[string]string
// @Failure 404 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /accounts/recovery [post]
func RecoveryHandler(ctx *gin.Context) {
	var req models.AccountRecoverRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	err := recovery.SendRecoveryEmail(ctx, req)
	if err != nil {
		if err == models.ErrEmailDoesNotMatchCitizenID || err == models.ErrCitizenIDNotExists {
			ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Wrong citizen ID or email"})
			return
		}
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Recovery email sent"})
}

// RecoveryConfirm godoc
// @Summary Confirm recovery OTP
// @Description Confirm recovery OTP
// @Tags accounts
// @Accept json
// @Produce json
// @Param credentials body models.AccountRecoverConfirmRequest true "Recovery OTP"
// @Success 200 {object} map[string]string
// @Failure 401 {object} map[string]string
// @Failure 404 {object} map[string]string
// @Router /accounts/recovery/confirm [post]
func RecoveryConfirmHandler(ctx *gin.Context) {
	var req models.AccountRecoverConfirmRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	err := recovery.VerifyRecoveryOTP(ctx, req)
	if err != nil {
		log.Println("Verifcation error:", err)
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid OTP"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "OTP verified"})
}

// ResetPassword godoc
// @Summary Reset password
// @Description Reset password after OTP verification
// @Tags accounts
// @Accept json
// @Produce json
// @Param credentials body models.PasswordResetRequest true "Password reset request"
// @Success 200 {object} map[string]string
// @Failure 401 {object} map[string]string
// @Failure 404 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /accounts/recovery/reset [put]
func ResetPasswordHandler(ctx *gin.Context) {
	var req models.PasswordResetRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	err := recovery.ResetPassword(ctx, req)
	if err != nil {
		log.Println("Reset password error:", err)
		if err == models.ErrExpiredOTP || err == models.ErrUnverifiedOTP {
			ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Unverified or expired OTP"})
			return
		}
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Password changed successfully"})
}
