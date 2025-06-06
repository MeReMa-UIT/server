package prescriptions_api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	errs "github.com/merema-uit/server/models/errors"
	prescription_services "github.com/merema-uit/server/services/prescriptions"
	"github.com/merema-uit/server/utils"
)

// Delete prescription detail godoc
// @Summary Delete prescription detail (doctor)
// @Description Delete a detail from the prescription
// @Tags prescriptions
// @Accept json
// @Produce json
// @Param prescription_id path string true "Prescription ID"
// @Param detail_id path string true "Detail ID"
// @Security BearerAuth
// @Success 200
// @Failure 401
// @Failure 403
// @Failure 404
// @Failure 500
// @Router /prescriptions/{prescription_id}/{detail_id} [delete]
func DeletePrescriptionDetailHandler(c *gin.Context) {
	authHeader := c.GetHeader("Authorization")
	prescriptionID := c.Param("prescription_id")
	detailID := c.Param("detail_id")

	err := prescription_services.DeletePrescriptionDetail(c, authHeader, prescriptionID, detailID)
	if err != nil {
		switch err {
		case errs.ErrExpiredToken, errs.ErrInvalidToken, errs.ErrMalformedToken:
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
		case errs.ErrPermissionDenied:
			c.JSON(http.StatusForbidden, gin.H{"error": err.Error()})
		case errs.ErrPrescriptionNotFound:
			c.JSON(http.StatusNotFound, gin.H{"error": "Prescription not found"})
		default:
			utils.Logger.Error("Failed to delete prescription detail", "error", err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
		}
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Prescription detail deleted successfully"})
}

// func DeletePrescriptionHandler(c *gin.Context) {
// 	authHeader := c.GetHeader("Authorization")
// 	prescriptionID := c.Param("prescription_id")

// 	err := prescription_services.DeletePrescription(c, authHeader, prescriptionID)
// 	if err != nil {
// 		switch err {
// 		case errs.ErrExpiredToken, errs.ErrInvalidToken, errs.ErrMalformedToken:
// 			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
// 		case errs.ErrPermissionDenied:
// 			c.JSON(http.StatusForbidden, gin.H{"error": err.Error()})
// 		case errs.ErrPrescriptionNotFound:
// 			c.JSON(http.StatusNotFound, gin.H{"error": "Prescription not found"})
// 		default:
// 			utils.Logger.Error("Failed to delete prescription", "error", err)
// 			c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
// 		}
// 		return
// 	}
// 	c.JSON(http.StatusOK, gin.H{"message": "Prescription deleted successfully"})
// }
