package schedule_api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/merema-uit/server/models"
	errs "github.com/merema-uit/server/models/errors"
	schedule_services "github.com/merema-uit/server/services/schedule"
	"github.com/merema-uit/server/utils"
)

// Update Schedule status godoc
// @Summary Update schedule status (receptionist)
// @Description Update schedule status from waiting to completed or cancelled
// @Tags schedules
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param request body models.UpdateScheduleStatusRequest true "Update schedule status request"
// @Success 200
// @Failure 400
// @Failure 401
// @Failure 403
// @Failure 500
// @Router /schedules/update-status [put]
func UpdateScheduleStatusHandler(c *gin.Context) {
	authHeader := c.GetHeader("Authorization")
	var req models.UpdateScheduleStatusRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}
	err := schedule_services.UpdateScheduleStatus(c.Request.Context(), authHeader, req)
	if err != nil {
		switch err {
		case errs.ErrInvalidScheduleStatus:
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		case errs.ErrInvalidToken, errs.ErrExpiredToken, errs.ErrMalformedToken:
			c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		case errs.ErrPermissionDenied:
			c.JSON(http.StatusForbidden, gin.H{"error": "Permission denied"})
		default:
			utils.Logger.Error("Error updating schedule status", "error", err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
		}
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Schedule status updated successfully"})
}
