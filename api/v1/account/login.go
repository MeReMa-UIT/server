package account_api

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/merema-uit/server/models"
	errs "github.com/merema-uit/server/models/errors"
	auth_services "github.com/merema-uit/server/services/auth"
	"github.com/merema-uit/server/utils"
)

// Login godoc
// @Summary Login
// @Description API for user to login
// @Tags accounts
// @Accept json
// @Produce json
// @Param credentials body models.LoginRequest true "Login credentials"
// @Success 200 {object} models.LoginResponse
// @Failure 400
// @Failure 401
// @Failure 404
// @Failure 500
// @Router /accounts/login [post]
func LoginHandler(ctx *gin.Context) {
	var req models.LoginRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request format"})
		return
	}

	token, err := auth_services.NewSession(context.Background(), req)

	if err != nil {
		switch err {
		case errs.ErrPasswordIncorrect, errs.ErrAccountNotExist:
			ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Identifier or Password is incorrect"})
		default:
			utils.Logger.Error("Can't create new session", "error", err.Error())
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
		}
		return
	}

	ctx.JSON(http.StatusOK, models.LoginResponse{Token: token})
}
