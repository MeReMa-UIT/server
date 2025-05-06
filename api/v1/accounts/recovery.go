package accounts_api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/merema-uit/server/models"
	"github.com/merema-uit/server/models/errors"
	"github.com/merema-uit/server/services/recovery"
	"github.com/merema-uit/server/utils"
)

// Account Recovery request godoc
// @Summary Get important info to send recovery email
// @Description Send recovery email to user
// @Tags accounts
// @Accept json
// @Produce json
// @Param credentials body models.AccountRecoverRequest true "Recovery credentials"
// @Success 200
// @Failure 400
// @Failure 401
// @Failure 500
// @Router /accounts/recovery [post]
func RecoveryHandler(ctx *gin.Context) {
	var req models.AccountRecoverRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request format"})
		return
	}

	err := recovery.SendRecoveryEmail(ctx, req)
	if err != nil {
		switch err {
		case errors.ErrAccountNotExist:
			ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Citizen ID or Email is incorrect"})
		default:
			utils.Logger.Error("Can't send recovery email", "error", err.Error())
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
		}
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Recovery email sent"})
}

// Recovery Confirm godoc
// @Summary Confirm recovery OTP
// @Description Confirm recovery OTP
// @Tags accounts
// @Accept json
// @Produce json
// @Param credentials body models.AccountRecoverConfirmRequest true "Recovery OTP"
// @Success 200 {object} models.AccountRecoverConfirmResponse
// @Failure 400
// @Failure 401
// @Failure 500
// @Router /accounts/recovery/confirm [post]
func RecoveryConfirmHandler(ctx *gin.Context) {
	var req models.AccountRecoverConfirmRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request format"})
		return
	}

	token, err := recovery.VerifyRecoveryOTP(ctx, req)
	if err != nil {
		switch err {
		case errors.ErrInvalidToken, errors.ErrWrongOTP:
			ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid OTP"})
		default:
			utils.Logger.Error("Can't confirm the recovery request", "error", err.Error())
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
		}
		return
	}

	ctx.JSON(http.StatusOK, models.AccountRecoverConfirmResponse{Token: token})
}

// Reset Password godoc
// @Summary Reset password
// @Description Reset password after OTP verification
// @Tags accounts
// @Accept json
// @Produce json
// @Param credentials body models.PasswordResetRequest true "Password reset request"
// @Security BearerAuth
// @Success 200
// @Failure 400
// @Failure 401
// @Failure 403
// @Failure 500
// @Router /accounts/recovery/reset [put]
func ResetPasswordHandler(ctx *gin.Context) {
	var req models.PasswordResetRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request format"})
		return
	}

	authHeader := ctx.GetHeader("Authorization")

	err := recovery.ResetPassword(ctx, req, authHeader)
	if err != nil {
		switch err {
		case errors.ErrInvalidToken, errors.ErrMalformedToken, errors.ErrExpiredToken:
			ctx.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		case errors.ErrExpiredOTP:
			ctx.JSON(http.StatusForbidden, gin.H{"error": "Expired OTP"})
		default:
			utils.Logger.Error("Can't change to the new password", "error", err.Error())
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
		}
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Password changed successfully"})
}
