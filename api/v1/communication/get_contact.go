package comm_api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	errs "github.com/merema-uit/server/models/errors"
	comm_services "github.com/merema-uit/server/services/communication"
	"github.com/merema-uit/server/utils"
)

// Get contact list godoc
// @Summary      Get contact list (doctor, patient)
// @Description  Get contact list for the authenticated user
// @Tags         communications
// @Accept       json
// @Produce      json
// @Security     BearerAuth
// @Success      200  {array}  models.ContactInfo
// @Failure      401
// @Failure      403
// @Failure      500
// @Router       /comms/contacts [get]
func GetContactListHandler(c *gin.Context) {
	authHeader := c.GetHeader("Authorization")
	contactList, err := comm_services.GetContactList(c.Request.Context(), authHeader)
	if err != nil {
		switch err {
		case errs.ErrExpiredToken, errs.ErrInvalidToken, errs.ErrMalformedToken:
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
		case errs.ErrPermissionDenied:
			c.JSON(http.StatusForbidden, gin.H{"error": "Permission denied"})
		default:
			utils.Logger.Error("Failed to retrieve contact list", "error", err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve contact list"})
		}
		return
	}
	c.JSON(http.StatusOK, contactList)
}
