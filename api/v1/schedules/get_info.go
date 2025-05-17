package schedules_api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/merema-uit/server/models"
	errs "github.com/merema-uit/server/models/errors"
	schedule_services "github.com/merema-uit/server/services/schedules"
	"github.com/merema-uit/server/utils"
)

// GetScheduleListHandler godoc
// @Summary Get Schedule List (patient, receptionist)
// @Description Get Schedule List
// @Tags schedules
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param type[] query []int false "Type of examination (1: Regular, 2: Service). Ex: ?type[]=1&type[]=2"
// @Param status[] query []int false "Status of the schedule (1: Waiting, 2: Completed, 3: Cancelled. Ex: ?status[]=1&status[]=2"
// @Success 200 {array} models.ScheduleInfo
// @Failure 400
// @Failure 401
// @Failure 403
// @Failure 500
// @Router /schedules [get]
func GetScheduleListHandler(c *gin.Context) {
	authHeader := c.GetHeader("Authorization")

	var req models.GetScheduleListRequest

	if err := c.ShouldBindQuery(&req); err != nil {
		utils.Logger.Error("Error binding query parameters", "error", err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	schedules, err := schedule_services.GetScheduleList(c, authHeader, req)
	if err != nil {
		switch err {
		case errs.ErrInvalidToken, errs.ErrExpiredToken, errs.ErrMalformedToken:
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
		case errs.ErrPermissionDenied:
			c.JSON(http.StatusForbidden, gin.H{"error": "Permission denied"})
		default:
			utils.Logger.Error("Error getting schedules list", "error", err.Error())
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
		}
		return
	}
	c.IndentedJSON(http.StatusOK, schedules)
}
