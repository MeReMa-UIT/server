package patient_api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/merema-uit/server/models"
	errs "github.com/merema-uit/server/models/errors"
	info_update_services "github.com/merema-uit/server/services/info_update"
	"github.com/merema-uit/server/utils"
)

// Update patient info godoc
// @Summary Update patient info (doctor, receptionist)
// @Description Update patient info
// @Tags patients
// @Accept json
// @Produce json
// @Param patient_id path string true "Patient ID"
// @Param patient_info body models.PatientInfoUpdateRequest true "Patient Info Update Request"
// @Security BearerAuth
// @Success 200
// @Failure 400
// @Failure 401
// @Failure 403
// @Failure 404
// @Failure 500
// @Router /patients/{patient_id} [put]
func UpdatePatientInfoHandler(c *gin.Context) {
	patientID := c.Param("patient_id")
	authHeader := c.GetHeader("Authorization")

	var updatedInfo models.PatientInfoUpdateRequest
	if err := c.ShouldBindJSON(&updatedInfo); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	err := info_update_services.UpdatePatientInfo(c, authHeader, patientID, updatedInfo)

	if err != nil {
		switch err {
		case errs.ErrInvalidToken, errs.ErrExpiredToken, errs.ErrMalformedToken:
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
		case errs.ErrPermissionDenied:
			c.JSON(http.StatusForbidden, gin.H{"error": "Permission denied"})
		case errs.ErrPatientNotExist:
			c.JSON(http.StatusNotFound, gin.H{"error": "Patient not found"})
		default:
			utils.Logger.Error("Failed to update patient info", "error", err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
		}
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Patient info updated successfully"})
}
