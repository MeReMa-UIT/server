package record_api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/merema-uit/server/models"
	record_services "github.com/merema-uit/server/services/record"
	"github.com/merema-uit/server/utils"
)

// Update medical record godoc
// @Summary Update a medical record (doctor)
// @Description Update a medical record by record ID
// @Tags records
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param record_id path string true "Record ID"
// @Param body body models.UpdateMedicalRecordRequest true "Update medical record request"
// @Success 200
// @Failure 400
// @Failure 401
// @Failure 403
// @Failure 500
// @Router /records/{record_id} [put]
func UpdateMedicalRecordHandler(c *gin.Context) {
	authHeader := c.GetHeader("Authorization")
	recordID := c.Param("record_id")

	var req models.UpdateMedicalRecordRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	err := record_services.UpdateMedicalRecord(c, authHeader, recordID, req)

	if err != nil {
		utils.Logger.Error("Failed to update medical record:", "error", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Medical record updated successfully"})
}
