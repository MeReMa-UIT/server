package accounts_api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	errs "github.com/merema-uit/server/models/errors"
	"github.com/merema-uit/server/services/retrieval"
	"github.com/merema-uit/server/utils"
)

// Get account info godoc
// @Summary Get account info
// @Description API for user to get account info
// @Tags accounts
// @Accept json
// @Produce json
// @Security BearerAuth
// @Success 200 {object} map[string]interface{} "account_info: models.AccountInfo, additional_info: []models.PatientBriefInfo or models.StaffInfo"
// @Failure 400
// @Failure 401
// @Failure 404
// @Failure 500
// @Router /accounts/profile [get]
func GetAccountInfoHandler(c *gin.Context) {
	authHeader := c.GetHeader("Authorization")
	accountInfo, additionalInfo, err := retrieval.GetAccountInfo(c.Request.Context(), authHeader)
	if err != nil {
		switch err {
		case errs.ErrAccountNotExist:
			c.JSON(http.StatusNotFound, gin.H{"error": "Account not found"})
		case errs.ErrExpiredToken, errs.ErrMalformedToken, errs.ErrInvalidToken:
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
		default:
			utils.Logger.Error("Error getting account info", "error", err.Error())
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
		}
		return
	}
	c.IndentedJSON(http.StatusOK, gin.H{"account_info": accountInfo, "additional_info": additionalInfo})
}

// Get account list godoc
// @Summary Get account list (admin)
// @Description API for admin to get account list
// @Tags accounts
// @Accept json
// @Produce json
// @Security BearerAuth
// @Success 200 {array} models.AccountInfo
// @Failure 400
// @Failure 401
// @Failure 403
// @Failure 500
// @Router /accounts [get]
func GetAccountListHandler(c *gin.Context) {
	authHeader := c.GetHeader("Authorization")
	accList, err := retrieval.GetAccountList(c.Request.Context(), authHeader)
	if err != nil {
		switch err {
		case errs.ErrExpiredToken, errs.ErrMalformedToken, errs.ErrInvalidToken:
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
		case errs.ErrPermissionDenied:
			c.JSON(http.StatusForbidden, gin.H{"error": "Permission denied"})
		default:
			utils.Logger.Error("Error getting accounts list", "error", err.Error())
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
		}
		return
	}
	c.IndentedJSON(http.StatusOK, accList)
}
