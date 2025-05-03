package accounts_api

import (
	"context"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/merema-uit/server/models"
	"github.com/merema-uit/server/services/register"
)

// Register godoc
// @Summary Register new user
// @Description Create a new user account
// @Tags accounts
// @Accept json
// @Produce json
// @Param user body models.AccountRegisterRequest true "User registration data"
// @Success 201 {object} map[string]string
// @Failure 404 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /accounts/register [post]
func RegisterHandler(ctx *gin.Context) {
	var req models.AccountRegisterRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	err := register.RegisterAccount(context.Background(), req)

	if err != nil {
		log.Println("Register error:", err)
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{"message": "User registered successfully"})
}
