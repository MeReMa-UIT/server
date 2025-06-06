package records_api

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	errs "github.com/merema-uit/server/models/errors"
	record_services "github.com/merema-uit/server/services/records"
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

// Get record info godoc
// @Summary Get record info (doctor, patient)
// @Description Get detailed information about a specific medical record
// @Tags records
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param recordID path string true "Record ID"
// @Success 200 {object} models.MedicalRecordInfo
// @Failure 400
// @Failure 401
// @Failure 403
// @Failure 500
// @Router /records/{recordID} [get]
func GetRecordInfoHandler(c *gin.Context) {
	authHeader := c.GetHeader("Authorization")
	recordID := c.Param("recordID")

	res, err := record_services.GetRecordInfo(c, authHeader, recordID)
	if err != nil {
		switch err {
		case errs.ErrExpiredToken, errs.ErrInvalidToken, errs.ErrMalformedToken:
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
		case errs.ErrPermissionDenied:
			c.JSON(http.StatusForbidden, gin.H{"error": "Permission denied"})
		default:
			utils.Logger.Error("Error getting record info:", "error", err.Error())
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
		}
		return
	}

	c.JSON(http.StatusOK, res)
}

// Get record attachments godoc
// @Summary Get record attachments (doctor, patient)
// @Description Get all attachments of a specific medical record
// @Tags records
// @Accept json
// @Produce application/zip
// @Security BearerAuth
// @Param recordID path string true "Record ID"
// @Success 200 {file} file "ZIP file containing all record's attachments"
// @Failure 400
// @Failure 401
// @Failure 403
// @Failure 500
// @Router /records/{recordID}/attachments [get]
func GetRecordAttachmentsHandler(c *gin.Context) {
	authHeader := c.GetHeader("Authorization")
	recordID := c.Param("recordID")

	attachments, err := record_services.GetRecordAttachments(c, authHeader, recordID)
	if err != nil {
		switch err {
		case errs.ErrExpiredToken, errs.ErrInvalidToken, errs.ErrMalformedToken:
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
		case errs.ErrPermissionDenied:
			c.JSON(http.StatusForbidden, gin.H{"error": "Permission denied"})
		default:
			utils.Logger.Error("Error getting record attachments:", "error", err.Error())
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
		}
		return
	}
	c.Header("Content-Disposition", fmt.Sprintf(`attachment; filename="%s.zip"`, recordID))
	c.Data(http.StatusOK, "application/zip", attachments.Bytes())
}
