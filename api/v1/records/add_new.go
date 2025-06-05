package records_api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	_ "github.com/jackc/pgtype"
	"github.com/merema-uit/server/models"
	errs "github.com/merema-uit/server/models/errors"
	record_services "github.com/merema-uit/server/services/records"
)

// Add new record godoc
// @Summary Add a new record (doctor)
// @Description Add a new record for a patient
// @Tags records
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param NewMedicalRecordRequest body models.NewMedicalRecordRequest true "New Medical Record Request"
// @Success 201
// @Failure 400
// @Failure 401
// @Failure 500
// @Router /records [post]
func AddNewRecordHandler(c *gin.Context) {
	authHeader := c.GetHeader("Authorization")
	var req models.NewMedicalRecordRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}
	err := record_services.AddNewRecord(c, authHeader, req)
	if err != nil {
		switch err {
		case errs.ErrExpiredToken, errs.ErrInvalidToken, errs.ErrMalformedToken:
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
		case errs.ErrPermissionDenied:
			c.JSON(http.StatusForbidden, gin.H{"error": "Permission denied"})
		default:
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
		}
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Record added successfully"})
}
