package staffs_api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	errs "github.com/merema-uit/server/models/errors"
	"github.com/merema-uit/server/services/retrieval"
	"github.com/merema-uit/server/utils"
)

// Get Staff list godoc
// @Summary Get staff list (admin)
// @Description Get staff list
// @Tags staffs
// @Accept json
// @Produce json
// @Security BearerAuth
// @Success 200 {array} models.StaffInfo
// @Failure 401
// @Failure 403
// @Failure 500
// @Router /staffs [get]
func GetStaffList(c *gin.Context) {
	authHeader := c.GetHeader("Authorization")
	list, err := retrieval.GetStaffList(c, authHeader)
	if err != nil {
		switch err {
		case errs.ErrInvalidToken, errs.ErrExpiredToken, errs.ErrMalformedToken:
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
		case errs.ErrPermissionDenied:
			c.JSON(http.StatusForbidden, gin.H{"error": "Permission denied"})
		default:
			utils.Logger.Error("Error getting staffs list", "error", err.Error())
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
		}
	}
	c.IndentedJSON(http.StatusOK, list)
}

// Get Staff info godoc
// @Summary Get staff info (primary for admin; doctor, receptionist will only get their own info for whichever staff_id they set)
// @Description Get staff info
// @Tags staffs
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param staff_id path string true "Staff ID"
// @Success 200 {object} models.StaffInfo
// @Failure 401
// @Failure 403
// @Failure 404
// @Failure 500
// @Router /staffs/{staff_id} [get]
func GetStaffInfoHandler(c *gin.Context) {
	authHeader := c.GetHeader("Authorization")
	staffID := c.Param("staff_id")
	info, err := retrieval.GetStaffInfo(c, authHeader, staffID)
	if err != nil {
		switch err {
		case errs.ErrInvalidToken, errs.ErrExpiredToken, errs.ErrMalformedToken:
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
		case errs.ErrPermissionDenied:
			c.JSON(http.StatusForbidden, gin.H{"error": "Permission denied"})
		case errs.ErrStaffNotExist:
			c.JSON(http.StatusNotFound, gin.H{"error": "Staff not found"})
		default:
			utils.Logger.Error("Error getting staff info", "error", err.Error())
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
		}
		return
	}
	c.IndentedJSON(http.StatusOK, info)
}
