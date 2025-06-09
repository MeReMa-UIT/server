package prescription_api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	errs "github.com/merema-uit/server/models/errors"
	prescription_services "github.com/merema-uit/server/services/prescription"
	"github.com/merema-uit/server/utils"
)

// Get prescription list with medical record ID godoc
// @Summary Get prescription list with medical record ID (doctor, patient)
// @Description Get prescription list
// @Tags prescriptions
// @Accept json
// @Produce json
// @Security BearerAuth
// @Success 200 {array} models.PrescriptionInfo
// @Failure 400
// @Failure 401
// @Failure 403
// @Failure 500
// @Router /prescriptions [get]
func GetPrescriptionListHandler(c *gin.Context) {
	authHeader := c.GetHeader("Authorization")

	list, err := prescription_services.GetPrescriptionList(c, authHeader)
	if err != nil {
		switch err {
		case errs.ErrPermissionDenied:
			c.JSON(http.StatusForbidden, gin.H{"error": "Permission denied"})
		case errs.ErrExpiredToken, errs.ErrInvalidToken, errs.ErrMalformedToken:
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
		default:
			utils.Logger.Error("Failed to get prescription list", "error", err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
		}
		return
	}
	c.JSON(http.StatusOK, list)
}

// Get prescription list with patient ID godoc
// @Summary Get prescription list with patient ID (doctor, patient)
// @Description Get prescription list
// @Tags prescriptions
// @Accept json
// @Produce json
// @Param patient_id path string true "Patient ID"
// @Security BearerAuth
// @Success 200 {array} models.PrescriptionInfo
// @Failure 400
// @Failure 401
// @Failure 403
// @Failure 500
// @Router /prescriptions/patients/{patient_id} [get]
func GetPrescriptionListByPatientIDHandler(c *gin.Context) {
	authHeader := c.GetHeader("Authorization")
	patientID := c.Param("patient_id")

	list, err := prescription_services.GetPrescriptionListByPatientID(c, authHeader, patientID)
	if err != nil {
		switch err {
		case errs.ErrPermissionDenied:
			c.JSON(http.StatusForbidden, gin.H{"error": "Permission denied"})
		case errs.ErrExpiredToken, errs.ErrInvalidToken, errs.ErrMalformedToken:
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
		default:
			utils.Logger.Error("Failed to get prescription list", "error", err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
		}
		return
	}
	c.JSON(http.StatusOK, list)
}

// Get prescription info (exluding details) by medical record ID godoc
// @Summary Get prescription info by medical record ID (doctor, patient)
// @Description Get prescription info by medical record ID
// @Tags prescriptions
// @Accept json
// @Produce json
// @Param record_id path string true "Medical Record ID"
// @Security BearerAuth
// @Success 200 {object} models.PrescriptionInfo
// @Failure 400
// @Failure 401
// @Failure 403
// @Failure 500
// @Router /prescriptions/records/{record_id} [get]
func GetPrescriptionInfoByRecordIDHandler(c *gin.Context) {
	authHeader := c.GetHeader("Authorization")
	recordID := c.Param("record_id")

	prescription, err := prescription_services.GetPrescriptionInfoByRecordID(c, authHeader, recordID)
	if err != nil {
		switch err {
		case errs.ErrPermissionDenied:
			c.JSON(http.StatusForbidden, gin.H{"error": "Permission denied"})
		case errs.ErrExpiredToken, errs.ErrInvalidToken, errs.ErrMalformedToken:
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
		default:
			utils.Logger.Error("Failed to get prescription ID by record ID", "error", err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
		}
		return
	}
	c.JSON(http.StatusOK, prescription)
}
