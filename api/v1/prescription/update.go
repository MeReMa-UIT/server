package prescription_api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/merema-uit/server/models"
	errs "github.com/merema-uit/server/models/errors"
	prescription_services "github.com/merema-uit/server/services/prescription"
	"github.com/merema-uit/server/utils"
)

// Update prescription godoc
// @Summary Update prescription (doctor)
// @Description Update prescription
// @Tags prescriptions
// @Accept json
// @Produce json
// @Param prescription_id path string true "Prescription ID"
// @Param body body models.PrescriptionUpdateRequest true "Prescription Update Request"
// @Security BearerAuth
// @Success 200
// @Failure 400
// @Failure 401
// @Failure 403
// @Failure 404
// @Failure 500
// @Router /prescriptions/{prescription_id} [put]
func UpdatePrescriptionHandler(c *gin.Context) {
	prescriptionID := c.Param("prescription_id")
	authHeader := c.GetHeader("Authorization")
	var req models.PrescriptionUpdateRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	err := prescription_services.UpdatePrescription(c, authHeader, prescriptionID, req)
	if err != nil {
		switch err {
		case errs.ErrWrongDosageCalulation, errs.ErrInvalidDosage:
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		case errs.ErrExpiredToken, errs.ErrInvalidToken, errs.ErrMalformedToken:
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
		case errs.ErrPermissionDenied, errs.ErrReceivedPrescription:
			c.JSON(http.StatusForbidden, gin.H{"error": err.Error()})
		case errs.ErrPrescriptionNotFound:
			c.JSON(http.StatusNotFound, gin.H{"error": "Prescription not found"})
		default:
			utils.Logger.Error("Failed to update prescription", "error", err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
		}
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Prescription updated successfully"})
}

// Confirm receiving prescription godoc
// @Summary Confirm receiving prescription (doctor)
// @Description Confirm that the prescription has been received
// @Tags prescriptions
// @Accept json
// @Produce json
// @Param prescription_id path string true "Prescription ID"
// @Security BearerAuth
// @Success 200
// @Failure 401
// @Failure 403
// @Failure 404
// @Failure 500
// @Router /prescriptions/{prescription_id}/confirm [put]
func ConfirmReceivingPrescriptionHandler(c *gin.Context) {
	authHeader := c.GetHeader("Authorization")
	prescriptionID := c.Param("prescription_id")
	err := prescription_services.ConfirmReceivingPrescription(c, authHeader, prescriptionID)
	if err != nil {
		switch err {
		case errs.ErrExpiredToken, errs.ErrInvalidToken, errs.ErrMalformedToken:
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
		case errs.ErrPermissionDenied, errs.ErrReceivedPrescription:
			c.JSON(http.StatusForbidden, gin.H{"error": err.Error()})
		case errs.ErrPrescriptionNotFound:
			c.JSON(http.StatusNotFound, gin.H{"error": "Prescription not found"})
		default:
			utils.Logger.Error("Failed to confirm receiving prescription", "error", err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
		}
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Prescription receiving confirmed successfully"})
}

// Update prescription detail godoc
// @Summary Update prescription detail (doctor)
// @Description Update prescription detail
// @Tags prescriptions
// @Accept json
// @Produce json
// @Param prescription_id path string true "Prescription ID"
// @Param med_id path string true "Medication ID"
// @Param body body models.PrescriptionDetail true "New Prescription Detail"
// @Security BearerAuth
// @Success 200
// @Failure 400
// @Failure 401
// @Failure 403
// @Failure 404
// @Failure 500
// @Router /prescriptions/{prescription_id}/{med_id} [put]
func UpdatePrescriptionDetailHandler(c *gin.Context) {
	authHeader := c.GetHeader("Authorization")
	prescriptionID := c.Param("prescription_id")
	medID := c.Param("med_id")
	var req models.PrescriptionDetailInfo
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	err := prescription_services.UpdatePrescriptionDetail(c, authHeader, prescriptionID, medID, req)
	if err != nil {
		switch err {
		case errs.ErrWrongDosageCalulation, errs.ErrInvalidDosage:
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		case errs.ErrExpiredToken, errs.ErrInvalidToken, errs.ErrMalformedToken:
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
		case errs.ErrPermissionDenied:
			c.JSON(http.StatusForbidden, gin.H{"error": err.Error()})
		case errs.ErrPrescriptionNotFound, errs.ErrPrescriptionDetailNotFound:
			c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		default:
			utils.Logger.Error("Failed to update prescription detail", "error", err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
		}
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Prescription detail updated successfully"})
}
