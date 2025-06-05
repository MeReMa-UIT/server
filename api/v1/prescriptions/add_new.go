package prescriptions_api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/merema-uit/server/models"
	errs "github.com/merema-uit/server/models/errors"
	prescription_services "github.com/merema-uit/server/services/prescriptions"
	"github.com/merema-uit/server/utils"
)

// Add new prescription godoc
// @Summary Add New Prescription (doctor)
// @Description Add a new prescription for a patient record
// @Tags prescriptions
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param request body models.NewPrescriptionRequest true "Add New Prescription Request"
// @Success 201
// @Failure 400
// @Failure 401
// @Failure 403
// @Failure 500
// @Router /prescriptions [post]
func AddNewPrescriptionHandler(c *gin.Context) {
	var req models.NewPrescriptionRequest
	authHeader := c.GetHeader("Authorization")
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	err := prescription_services.AddNewPrescription(c, authHeader, req)
	if err != nil {
		switch err {
		case errs.ErrWrongDosageCalulation, errs.ErrInvalidDosage:
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		case errs.ErrExpiredToken, errs.ErrInvalidToken, errs.ErrMalformedToken:
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
		case errs.ErrPermissionDenied:
			c.JSON(http.StatusForbidden, gin.H{"error": "Permission denied"})
		default:
			utils.Logger.Error("Failed to add new prescription", "error", err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
		}
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Prescription added successfully"})
}
