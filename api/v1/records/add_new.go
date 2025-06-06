package records_api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	_ "github.com/jackc/pgtype"
	"github.com/merema-uit/server/models"
	errs "github.com/merema-uit/server/models/errors"
	record_services "github.com/merema-uit/server/services/records"
	"github.com/merema-uit/server/utils"
)

// Add new record godoc
// @Summary Add a new record (doctor)
// @Description Add a new record for a patient
// @Tags records
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param NewMedicalRecordRequest body models.NewMedicalRecordRequest true "New Medical Record Request"
// @Success 201 {object} models.NewMedicalRecordResponse
// @Failure 400
// @Failure 401
// @Failure 403
// @Failure 500
// @Router /records [post]
func AddNewRecordHandler(c *gin.Context) {
	authHeader := c.GetHeader("Authorization")

	var req models.NewMedicalRecordRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}
	res, err := record_services.AddNewRecord(c, authHeader, &req)
	if err != nil {
		switch err {
		case errs.ErrInvalidMedicalRecordStructure, errs.ErrPrimaryDiagnosisMissing:
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		case errs.ErrExpiredToken, errs.ErrInvalidToken, errs.ErrMalformedToken:
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
		case errs.ErrPermissionDenied:
			c.JSON(http.StatusForbidden, gin.H{"error": "Permission denied"})
		default:
			utils.Logger.Error("Error adding new record:", "error", err.Error())
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
		}
		return
	}

	c.JSON(http.StatusCreated, res)
}

// Add record attachments godoc
// @Summary Add attachments to a record (doctor)
// @Description Add attachments to an existing record
// @Tags records
// @Accept multipart/form-data
// @Produce json
// @Security BearerAuth
// @Param record_id path string true "Record ID"
// @Param attachments formData file true "Attachments, support 5 types of prefix: xray_, ct_, ultrasound_, test_, other_. Ex: xray_1.jpg"
// @Success 201
// @Failure 400
// @Failure 401
// @Failure 403
// @Failure 500
// @Router /records/{record_id}/attachments [post]
func AddRecordAttachmentsHandler(c *gin.Context) {
	authHeader := c.GetHeader("Authorization")
	form, err := c.MultipartForm()

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid form data"})
		return
	}
	attachments := form.File["attachments"]
	recordID := c.Param("record_id")

	err = record_services.AddRecordAttachments(c, authHeader, recordID, attachments)
	if err != nil {
		switch err {
		case errs.ErrInvalidAttachmentPrefix:
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid attachment prefix"})
		case errs.ErrExpiredToken, errs.ErrInvalidToken, errs.ErrMalformedToken:
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
		case errs.ErrPermissionDenied:
			c.JSON(http.StatusForbidden, gin.H{"error": "Permission denied"})
		default:
			utils.Logger.Error("Error adding attachments:", "error", err.Error())
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
		}
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Attachments added successfully"})
}
