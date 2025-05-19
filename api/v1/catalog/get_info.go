package catalog_api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	errs "github.com/merema-uit/server/models/errors"
	"github.com/merema-uit/server/services/retrieval"
	"github.com/merema-uit/server/utils"
)

// Get Medication List godoc
// @Summary      Get Medication List (doctor)
// @Description  Get a list of medications
// @Tags         Catalog
// @Accept       json
// @Produce      json
// @Security     BearerAuth
// @Success      200  {array}  models.MedicationInfo
// @Failure      401
// @Failure      403
// @Failure      500
// @Router       /catalog/medications [get]
func GetMedicationListHandler(c *gin.Context) {
	authHeader := c.GetHeader("Authorization")
	list, err := retrieval.GetMedicationList(c, authHeader)
	if err != nil {
		switch err {
		case errs.ErrPermissionDenied:
			c.JSON(http.StatusForbidden, gin.H{"error": "Permission denied"})
		case errs.ErrExpiredToken, errs.ErrMalformedToken, errs.ErrInvalidToken:
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
		default:
			utils.Logger.Error("Error retrieving medication list", "error", err.Error())
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
		}
		return
	}
	c.IndentedJSON(http.StatusOK, list)
}

// Get Diagnosis List godoc
// @Summary      Get Diagnosis List (doctor)
// @Description  Get a list of diagnoses
// @Tags         Catalog
// @Accept       json
// @Produce      json
// @Security     BearerAuth
// @Success      200  {array}  models.DiagnosisInfo
// @Failure      401
// @Failure      403
// @Failure      500
// @Router       /catalog/diagnoses [get]
func GetDiagnosisListHandler(c *gin.Context) {
	authHeader := c.GetHeader("Authorization")
	list, err := retrieval.GetDiagnosisList(c, authHeader)
	if err != nil {
		switch err {
		case errs.ErrPermissionDenied:
			c.JSON(http.StatusForbidden, gin.H{"error": "Permission denied"})
		case errs.ErrExpiredToken, errs.ErrMalformedToken, errs.ErrInvalidToken:
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
		default:
			utils.Logger.Error("Error retrieving diagnosis list", "error", err.Error())
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
		}
		return
	}
	c.IndentedJSON(http.StatusOK, list)
}
