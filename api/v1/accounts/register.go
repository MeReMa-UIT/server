package accounts_api

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/merema-uit/server/models"
	"github.com/merema-uit/server/services/register"
	"github.com/merema-uit/server/utils"
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
// @Failure 400
// @Failure 401
// @Failure 403
// @Failure 500
// @Router /accounts/register/patients [post]
func RegisterPatientHandler(ctx *gin.Context) {
	var req models.PatientRegisterRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request format"})
		return
	}

	authHeader := ctx.GetHeader("Authorization")

	err := register.RegisterPatient(context.Background(), req, authHeader)

	if err != nil {
		switch err {
		default:
			utils.Logger.Error("Can't register new patient", "error", err.Error())
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
		}
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{"message": "User registered successfully"})
}
