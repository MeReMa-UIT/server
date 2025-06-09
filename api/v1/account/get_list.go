package account_api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	errs "github.com/merema-uit/server/models/errors"
	retrieval_services "github.com/merema-uit/server/services/retrieval"
	"github.com/merema-uit/server/utils"
)

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
	accList, err := retrieval_services.GetAccountList(c.Request.Context(), authHeader)
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
