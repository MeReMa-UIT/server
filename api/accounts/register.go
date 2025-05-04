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
// @Success 201 {object} map[string]string
// @Failure 404 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /accounts/register/patients [post]
// @Security BearerAuth
func RegisterPatientHandler(ctx *gin.Context) {
	var req models.PatientRegisterRequest
	authHeader := ctx.GetHeader("Authorization")
	if authHeader == "" {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Invalid authorization header"})
	}
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	err := register.RegisterPatient(context.Background(), req, authHeader)

	if err != nil {
		log.Println("Register error:", err)
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{"message": "User registered successfully"})
}
