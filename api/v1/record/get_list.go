package record_api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	errs "github.com/merema-uit/server/models/errors"
	record_services "github.com/merema-uit/server/services/record"
	"github.com/merema-uit/server/utils"
)

// Get record list godoc
// @Summary Get record list (doctor, patient)
// @Description Get a list of medical records for a patient or all records for a doctor
// @Tags records
// @Accept json
// @Produce json
// @Security BearerAuth
// @Success 200 {array} models.MedicalRecordBriefInfo
// @Failure 400
// @Failure 401
// @Failure 403
// @Failure 500
// @Router /records [get]
func GetRecordListHandler(c *gin.Context) {
	authHeader := c.GetHeader("Authorization")

	res, err := record_services.GetRecordList(c, authHeader)
	if err != nil {
		switch err {
		case errs.ErrExpiredToken, errs.ErrInvalidToken, errs.ErrMalformedToken:
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
		case errs.ErrPermissionDenied:
			c.JSON(http.StatusForbidden, gin.H{"error": "Permission denied"})
		default:
			utils.Logger.Error("Error getting record list:", "error", err.Error())
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
		}
		return
	}

	c.JSON(http.StatusOK, res)
}
