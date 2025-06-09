package staff_api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/merema-uit/server/models"
	errs "github.com/merema-uit/server/models/errors"
	info_update_services "github.com/merema-uit/server/services/info_update"
	"github.com/merema-uit/server/utils"
)

// Update staff info godoc
// @Summary Update staff info (admin)
// @Description Update staff information by staff ID
// @Tags staffs
// @Accept json
// @Produce json
// @Param staff_id path string true "Staff ID"
// @Param StaffInfoUpdateRequest body models.StaffInfoUpdateRequest true "Staff info update request"
// @Security BearerAuth
// @Success 200
// @Failure 400
// @Failure 401
// @Failure 403
// @Failure 404
// @Failure 500
// @Router /staffs/{staff_id} [put]
func UpdateStaffInfoHandler(c *gin.Context) {
	staffID := c.Param("staff_id")
	authHeader := c.GetHeader("Authorization")

	var updatedInfo models.StaffInfoUpdateRequest
	if err := c.ShouldBindJSON(&updatedInfo); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	err := info_update_services.UpdateStaffInfo(c, authHeader, staffID, updatedInfo)

	if err != nil {
		switch err {
		case errs.ErrInvalidToken, errs.ErrExpiredToken, errs.ErrMalformedToken:
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
		case errs.ErrPermissionDenied:
			c.JSON(http.StatusForbidden, gin.H{"error": "Permission denied"})
		case errs.ErrStaffNotExist:
			c.JSON(http.StatusNotFound, gin.H{"error": "Staff not found"})
		default:
			utils.Logger.Error("Failed to update staff info", "error", err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
		}
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Staff info updated successfully"})
}
