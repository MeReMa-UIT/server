package accounts_api

import (
	"context"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/merema-uit/server/models"
	"github.com/merema-uit/server/services/register"
)

// Register patient godoc
// @Summary Register new patient
// @Description Create a new patient account
// @Tags accounts
// @Accept json
// @Produce json
// @Param user body models.PatientRegisterRequest true "User registration data"
// @Security BearerAuth
// @Success 201
// @Failure 404
// @Failure 500
// @Router /accounts/register/patients [post]
func RegisterPatientHandler(ctx *gin.Context) {
	var req models.PatientRegisterRequest
	authHeader := ctx.GetHeader("Authorization")
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}
	if authHeader == "" {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Invalid authorization header"})
	}

	err := register.RegisterPatient(context.Background(), req, authHeader)

	if err != nil {
		log.Println("Register error:", err)
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{"message": "User registered successfully"})
}
