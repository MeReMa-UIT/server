package accounts_api

import (
	"context"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/merema-uit/server/models"
	"github.com/merema-uit/server/services/auth"
)

// Login godoc
// @Summary Login user
// @Description Return JWT token to auth user
// @Tags accounts
// @Accept json
// @Produce json
// @Param credentials body models.LoginRequest true "Login credentials"
// @Success 200 {object} models.LoginResponse
// @Failure 401 {object} map[string]string
// @Failure 404 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /accounts/login [post]
func LoginHandler(ctx *gin.Context) {
	var req models.LoginRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	token, err := auth.NewSession(context.Background(), req)

	if err != nil {
		log.Println("Login error:", err)
		if err == models.ErrPasswordIncorrect || err == models.ErrAccountNotExist {
			ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Wrong identifier (Email/Phone/Citizen ID) or Password"})
			return
		}
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
		return
	}

	ctx.JSON(http.StatusOK, models.LoginResponse{Token: token})
}
