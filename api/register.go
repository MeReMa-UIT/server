package api

import (
	"context"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/merema-uit/server/models"
	"github.com/merema-uit/server/services/register"
)

func Register(ctx *gin.Context) {
	var req models.AccountRegisterRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	err := register.RegisterPatientAccount(context.Background(), req)

	if err != nil {
		log.Println("Register error:", err)
		// if err == models.ErrUsernameExists {
		// 	ctx.JSON(http.StatusConflict, gin.H{"error": "Username already exists"})
		// 	return
		// }
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "User registered successfully"})
}
