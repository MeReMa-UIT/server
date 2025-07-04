package account_api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/merema-uit/server/models"
	errs "github.com/merema-uit/server/models/errors"
	info_update_services "github.com/merema-uit/server/services/info_update"
	"github.com/merema-uit/server/utils"
)

// Update Account Info godoc
// @Summary Update Account Info
// @Description Update account information such as citizen ID, phone, email, or password
// @Tags accounts
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param AccountInfoUpdateRequest body models.AccountInfoUpdateRequest true "Account information to update"
// @Success 200
// @Failure 400
// @Failure 401
// @Failure 403
// @Failure 404
// @Failure 500
// @Router /accounts/profile [put]
func UpdateAccountInfoHandler(c *gin.Context) {
	authHeader := c.GetHeader("Authorization")
	var req models.AccountInfoUpdateRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}
	err := info_update_services.UpdateAccountInfo(c, authHeader, req)
	if err != nil {
		switch err {
		case errs.ErrInvalidField, errs.ErrEmailOrPhoneAlreadyUsed:
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		case errs.ErrPasswordIncorrect:
			c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		case errs.ErrInvalidToken, errs.ErrExpiredToken, errs.ErrMalformedToken:
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
		case errs.ErrPermissionDenied:
			c.JSON(http.StatusForbidden, gin.H{"error": "Permission denied"})
		case errs.ErrAccountNotExist:
			c.JSON(http.StatusNotFound, gin.H{"error": "Account not found"})
		default:
			utils.Logger.Error("Failed to update account info", "error", err.Error())
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
		}
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": req.Field + " updated successfully"})
}
