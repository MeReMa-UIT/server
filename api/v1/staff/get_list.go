package staff_api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	errs "github.com/merema-uit/server/models/errors"
	retrieval_services "github.com/merema-uit/server/services/retrieval"
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
func GetStaffListHandler(c *gin.Context) {
	authHeader := c.GetHeader("Authorization")
	list, err := retrieval_services.GetStaffList(c, authHeader)
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
