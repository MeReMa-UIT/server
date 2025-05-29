package comms_api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/merema-uit/server/models"
	errs "github.com/merema-uit/server/models/errors"
	comm_services "github.com/merema-uit/server/services/communication"
	"github.com/merema-uit/server/utils"
)

// Send message godoc
// @Summary Send message (patient, doctor)
// @Description Send message to a doctor or patient
// @Tags communications
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param message body models.SendingMessage true "Message to send"
// @Success 200
// @Failure 400
// @Failure 401
// @Failure 403
// @Failure 500
// @Router /comms/messages [post]
func SendMessageHandler(c *gin.Context) {
	authHeader := c.GetHeader("Authorization")
	var message models.SendingMessage
	if err := c.ShouldBindJSON(&message); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}
	err := comm_services.SendMessage(c, authHeader, message)
	if err != nil {
		switch err {
		case errs.ErrInvalidToken, errs.ErrExpiredToken, errs.ErrMalformedToken:
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
		case errs.ErrPermissionDenied:
			c.JSON(http.StatusForbidden, gin.H{"error": "Permission denied"})
		case errs.ErrInvalidRecipient:
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid recipient"})
		default:
			utils.Logger.Error("Failed to send message", "error", err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
		}
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Message was delievered"})
}

// Load conversation godoc
// @Summary Load conversation (patient, doctor)
// @Description Load conversation with a doctor or patient
// @Tags communications
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param contact_id path string true "Contact ID"
// @Success 200 {array} models.Message
// @Failure 400
// @Failure 401
// @Failure 403
// @Failure 500
// @Router /comms/messages/{contact_id} [get]
func LoadConversationHandler(c *gin.Context) {
	authHeader := c.GetHeader("Authorization")
	contactID := c.Param("contact_id")
	if contactID == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}
	messages, err := comm_services.LoadConversation(c, authHeader, contactID)
	if err != nil {
		switch err {
		case errs.ErrInvalidToken, errs.ErrExpiredToken, errs.ErrMalformedToken:
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
		case errs.ErrPermissionDenied:
			c.JSON(http.StatusForbidden, gin.H{"error": "Permission denied"})
		default:
			utils.Logger.Error("Failed to load conversation", "error", err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
		}
		return
	}
	c.JSON(http.StatusOK, messages)
}
