package accounts_api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	errs "github.com/merema-uit/server/models/errors"
	getinfo "github.com/merema-uit/server/services/get_info"
)

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
