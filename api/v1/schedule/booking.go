package schedule_api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/merema-uit/server/models"
	errs "github.com/merema-uit/server/models/errors"
	schedule_services "github.com/merema-uit/server/services/schedule"
	"github.com/merema-uit/server/utils"
)

// Book Exmination Schedule godoc
// @Summary Book Examination Schedule (patient)
// @Description Book Examination Schedule
// @Tags schedules
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param schedule body models.ScheduleBookingRequest true "Schedule Booking Request"
// @Success 201 {object} models.ScheduleBookingResponse
// @Failure 400
// @Failure 401
// @Failure 403
// @Failure 404
// @Failure 500
// @Router /schedules/book [post]
func BookScheduleHandler(c *gin.Context) {
	authHeader := c.GetHeader("Authorization")
	var req models.ScheduleBookingRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	response, err := schedule_services.BookSchedule(c, authHeader, req)
	if err != nil {
		switch err {
		case errs.ErrInvalidExaminationType:
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		case errs.ErrInvalidToken, errs.ErrExpiredToken, errs.ErrMalformedToken:
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
		case errs.ErrPermissionDenied:
			c.JSON(http.StatusForbidden, gin.H{"error": "Permission denied"})
		case errs.ErrPatientNotExist:
			c.JSON(http.StatusNotFound, gin.H{"error": "Patient not found"})
		default:
			utils.Logger.Error("Failed to book schedule", "error", err.Error())
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
		}
		return
	}

	c.JSON(http.StatusCreated, response)
}
