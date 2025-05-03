package accounts_api

import (
	"context"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/merema-uit/server/models"
	"github.com/merema-uit/server/services/auth"
)

func LoginHandler(ctx *gin.Context) {
	var req models.LoginRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	token, err := auth.NewSession(context.Background(), req)

	if err != nil {
		log.Println("Login error:", err)
		if err == models.ErrPasswordIncorrect || err == models.ErrCitizenIDExists {
			ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Wrong citizen ID or password"})
			return
		}
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
		return
	}

	ctx.JSON(http.StatusOK, models.LoginResponse{Token: token})
}
