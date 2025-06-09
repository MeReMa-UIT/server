package catalog_api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	errs "github.com/merema-uit/server/models/errors"
	retrieval_services "github.com/merema-uit/server/services/retrieval"
	"github.com/merema-uit/server/utils"
)

// Get Medication Info godoc
// @Summary      Get Medication Info (doctor, patient)
// @Description  Get info of a medication
// @Tags         catalogs
// @Accept       json
// @Produce      json
// @Security     BearerAuth
// @Param        medication_id  path  string  true  "Medication ID"
// @Success      200  {object}  models.MedicationInfo
// @Failure      401
// @Failure      403
// @Failure      500
// @Router       /catalog/medications/{medication_id} [get]
func GetMedicationInfoHandler(c *gin.Context) {
	authHeader := c.GetHeader("Authorization")
	medicationID := c.Param("medication_id")
	medicationInfo, err := retrieval_services.GetMedicationInfo(c, authHeader, medicationID)
	if err != nil {
		switch err {
		case errs.ErrPermissionDenied:
			c.JSON(http.StatusForbidden, gin.H{"error": "Permission denied"})
		case errs.ErrExpiredToken, errs.ErrMalformedToken, errs.ErrInvalidToken:
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
		default:
			utils.Logger.Error("Error retrieving medication info", "error", err.Error())
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
		}
		return
	}
	c.IndentedJSON(http.StatusOK, medicationInfo)

}

// Get Diagnosis Info godoc
// @Summary      Get Diagnosis Info (doctor, patient)
// @Description  Get info of a diagnosis
// @Tags         catalogs
// @Accept       json
// @Produce      json
// @Security     BearerAuth
// @Param        icd_code  path  string  true  "ICD Code"
// @Success      200  {object}  models.DiagnosisInfo
// @Failure      401
// @Failure      403
// @Failure      500
// @Router       /catalog/diagnoses/{icd_code} [get]
func GetDiagnosisInfoHandler(c *gin.Context) {
	authHeader := c.GetHeader("Authorization")
	icdCode := c.Param("icd_code")
	diagnosisInfo, err := retrieval_services.GetDiagnosisInfo(c, authHeader, icdCode)
	if err != nil {
		switch err {
		case errs.ErrPermissionDenied:
			c.JSON(http.StatusForbidden, gin.H{"error": "Permission denied"})
		case errs.ErrExpiredToken, errs.ErrMalformedToken, errs.ErrInvalidToken:
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
		default:
			utils.Logger.Error("Error retrieving diagnosis info", "error", err.Error())
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
		}
		return
	}
	c.IndentedJSON(http.StatusOK, diagnosisInfo)
}

// Get Medical Record Template godoc
// @Summary      Get Medical Record Template (doctor)
// @Description  Get a medical record template by type ID
// @Tags         catalogs
// @Accept       json
// @Produce      json
// @Security     BearerAuth
// @Param        type_id  path  string  true  "Medical Record Type ID  (01BV1, 02BV1, ...)"
// @Success      200  {object} interface{}
// @Failure      401
// @Failure      403
// @Failure      500
// @Router       /catalog/record-types/{type_id}/template [get]
func GetMedicalRecordTemplateHandler(c *gin.Context) {
	authHeader := c.GetHeader("Authorization")
	typeID := c.Param("type_id")
	template, err := retrieval_services.GetMedicalRecordTemplate(c, authHeader, typeID)
	if err != nil {
		switch err {
		case errs.ErrPermissionDenied:
			c.JSON(http.StatusForbidden, gin.H{"error": "Permission denied"})
		case errs.ErrExpiredToken, errs.ErrMalformedToken, errs.ErrInvalidToken:
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
		default:
			utils.Logger.Error("Error retrieving record template", "error", err.Error())
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
		}
		return
	}
	c.IndentedJSON(http.StatusOK, template)
}
