package prescription_api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	errs "github.com/merema-uit/server/models/errors"
	prescription_services "github.com/merema-uit/server/services/prescription"
	"github.com/merema-uit/server/utils"
)

// Get prescription details godoc
// @Summary Get prescription details (doctor, patient)
// @Description Get prescription details by ID
// @Tags prescriptions
// @Accept json
// @Produce json
// @Param prescription_id path string true "Prescription ID"
// @Security BearerAuth
// @Success 200 {array} models.PrescriptionDetailInfo
// @Failure 400
// @Failure 401
// @Failure 403
// @Failure 500
// @Router /prescriptions/{prescription_id} [get]
func GetPrescriptionDetailsHandler(c *gin.Context) {
	authHeader := c.GetHeader("Authorization")
	prescriptionID := c.Param("prescription_id")

	details, err := prescription_services.GetPrescriptionDetails(c, authHeader, prescriptionID)
	if err != nil {
		switch err {
		case errs.ErrPermissionDenied:
			c.JSON(http.StatusForbidden, gin.H{"error": "Permission denied"})
		case errs.ErrExpiredToken, errs.ErrInvalidToken, errs.ErrMalformedToken:
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
		default:
			utils.Logger.Error("Failed to get prescription details", "error", err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
		}
		return
	}
	c.JSON(http.StatusOK, details)
}
