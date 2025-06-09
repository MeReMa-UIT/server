package patient_api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	errs "github.com/merema-uit/server/models/errors"
	retrieval_services "github.com/merema-uit/server/services/retrieval"
	"github.com/merema-uit/server/utils"
)

// Get Patient list godoc
// @Summary Get patient list (receptionist, doctor)
// @Description Get patient list
// @Tags patients
// @Accept json
// @Produce json
// @Security BearerAuth
// @Success 200 {array} models.PatientBriefInfo
// @Failure 401
// @Failure 403
// @Failure 500
// @Router /patients [get]
func GetPatientListHandler(c *gin.Context) {
	authHeader := c.GetHeader("Authorization")
	list, err := retrieval_services.GetPatientList(c, authHeader)
	if err != nil {
		switch err {
		case errs.ErrInvalidToken, errs.ErrExpiredToken, errs.ErrMalformedToken:
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
		case errs.ErrPermissionDenied:
			c.JSON(http.StatusForbidden, gin.H{"error": "Permission denied"})
		default:
			utils.Logger.Error("Error getting patients list", "error", err.Error())
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
		}
		return
	}
	c.IndentedJSON(http.StatusOK, list)
}
