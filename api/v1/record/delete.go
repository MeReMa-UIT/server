package record_api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/merema-uit/server/models"
	errs "github.com/merema-uit/server/models/errors"
	record_services "github.com/merema-uit/server/services/record"
	"github.com/merema-uit/server/utils"
)

// Delete record attachment godoc
// @Summary      Delete record attachment (doctor)
// @Description  Delete a record attachment by record ID and file name
// @Tags         records
// @Accept       json
// @Produce      json
// @Security     BearerAuth
// @Param        record_id path string true "Record ID"
// @Param        body body models.DeleteRecordAttachmentRequest true "Attachment file name to delete"
// @Success      200
// @Failure      400
// @Failure      401
// @Failure      403
// @Failure      404
// @Failure      500
// @Router       /records/{record_id}/attachments [delete]
func DeleteRecordAttachmentHandler(c *gin.Context) {
	authHeader := c.GetHeader("Authorization")
	recordID := c.Param("record_id")
	var req models.DeleteRecordAttachmentRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	err := record_services.DeleteRecordAttachment(c, authHeader, recordID, req)
	if err != nil {
		switch err {
		case errs.ErrInvalidAttachmentPrefix:
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid attachment prefix"})
		case errs.ErrInvalidToken, errs.ErrExpiredToken, errs.ErrMalformedToken:
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
		case errs.ErrPermissionDenied:
			c.JSON(http.StatusForbidden, gin.H{"error": "Permission denied"})
		case errs.ErrAttachmentNotFound:
			c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		default:
			utils.Logger.Error("Failed to delete attachment", "error", err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete attachment", "details": err.Error()})
		}
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Attachment deleted successfully"})
}
