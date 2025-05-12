package accounts_api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	errs "github.com/merema-uit/server/models/errors"
	getinfo "github.com/merema-uit/server/services/get_info"
)

// @Get account info godoc
// @Summary Get account info
// @Description API for user to get account info
// @Tags accounts
// @Accept json
// @Produce json
// @Security BearerAuth
// @Success 200 {object} models.AccountInfo
// @Failure 400
// @Failure 401
// @Failure 404
// @Failure 500
// @Router /accounts/get_info [get]
func GetAccountInfoHandler(c *gin.Context) {
	authHeader := c.GetHeader("Authorization")
	accountInfo, err := getinfo.GetAccountInfo(c.Request.Context(), authHeader)
	if err != nil {
		switch err {
		case errs.ErrAccountNotExist:
			c.JSON(http.StatusNotFound, gin.H{"error": "Account not found"})
		case errs.ErrExpiredToken, errs.ErrMalformedToken, errs.ErrInvalidToken:
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
		default:
			c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		}
		return
	}
	c.JSON(http.StatusOK, accountInfo)
}
