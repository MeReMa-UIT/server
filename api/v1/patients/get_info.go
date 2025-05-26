package patients_api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	errs "github.com/merema-uit/server/models/errors"
	"github.com/merema-uit/server/services/retrieval"
	"github.com/merema-uit/server/utils"
)

// Get Patient list godoc
// @Summary Get patient list (receptionist, doctor)
// @Description Get patient list
// @Tags patients
// @Accept json
// @Produce json
// @Security BearerAuth
// @Success 200 {array} models.PatientBriefInfo
// @Failure 401
// @Failure 403
// @Failure 500
// @Router /patients [get]
func GetPatientList(c *gin.Context) {
	authHeader := c.GetHeader("Authorization")
	list, err := retrieval.GetPatientList(c, authHeader)
	if err != nil {
		switch err {
		case errs.ErrInvalidToken, errs.ErrExpiredToken, errs.ErrMalformedToken:
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
		case errs.ErrPermissionDenied:
			c.JSON(http.StatusForbidden, gin.H{"error": "Permission denied"})
		default:
			utils.Logger.Error("Error getting patients list", "error", err.Error())
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
		}
		return
	}
	c.IndentedJSON(http.StatusOK, list)
}

// Get Patient info godoc
// @Summary Get patient info (receptionist, doctor, patient)
// @Description Get patient info
// @Tags patients
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param patient_id path string true "Patient ID"
// @Success 200 {object} models.PatientInfo
// @Failure 401
// @Failure 403
// @Failure 404
// @Failure 500
// @Router /patients/{patient_id} [get]
func GetPatientInfoHandler(c *gin.Context) {
	authHeader := c.GetHeader("Authorization")
	patientID := c.Param("patient_id")
	print("Patient ID: ", patientID)
	info, err := retrieval.GetPatientInfo(c, authHeader, patientID)
	if err != nil {
		switch err {
		case errs.ErrInvalidToken, errs.ErrExpiredToken, errs.ErrMalformedToken:
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
		case errs.ErrPermissionDenied:
			c.JSON(http.StatusForbidden, gin.H{"error": "Permission denied"})
		case errs.ErrPatientNotExist:
			c.JSON(http.StatusNotFound, gin.H{"error": "Patient not found"})
		default:
			utils.Logger.Error("Error getting patient info", "error", err.Error())
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
		}
		return
	}
	c.IndentedJSON(http.StatusOK, info)
}
