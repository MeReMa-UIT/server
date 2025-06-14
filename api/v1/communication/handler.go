package comm_api

import (
	"log"
	"net/http"

	// "strings" // Not strictly needed if only using query param for token

	"github.com/gin-gonic/gin" // Example JWT library; use your preferred one
	"github.com/gorilla/websocket"
	auth_services "github.com/merema-uit/server/services/auth"
	comm_services "github.com/merema-uit/server/services/communication"
)

// upgrader is used to upgrade HTTP GET requests to WebSocket connections.
var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	// CheckOrigin is crucial for security in production.
	// It prevents CSRF attacks by ensuring the request originates from an allowed domain.
	CheckOrigin: func(r *http.Request) bool {
		// TODO: In a production environment, you MUST implement proper origin checking.
		// Example:
		// origin := r.Header.Get("Origin")
		// allowedOrigins := []string{"https://yourfrontend.com", "http://localhost:3000"}
		// for _, allowed := range allowedOrigins {
		// 	if origin == allowed {
		// 		return true
		// 	}
		// }
		// log.Printf("WebSocket connection from disallowed origin: %s", origin)
		// return false
		log.Printf("WebSocket CheckOrigin: Allowing connection from origin %s (DEV MODE - implement strict check for PROD)", r.Header.Get("Origin"))
		return true // For development, allow all origins.
	},
}

// WebSocketHandler holds dependencies for WebSocket related HTTP handlers.
type WebSocketHandler struct {
	ChatService comm_services.Service // The core chat service (Hub)
}

// Chat service godoc
// @Summary WebSocket connection for chat (doctor, patient)
// @Description Establish a WebSocket connection for real-time chat between doctor and patient.
// @Tags communications
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param rqeuest body models.InboundMessage true "Comm request (type = {'sendMessage', 'loadHistory', 'markSeenMessage'})"
// @Success 101 {object} models.OutboundMessage "type = {'yourID', 'newMessage', 'messageHistory', 'conversationList', 'error'}"
// @Router /comms/chat [get]
func NewWebSocketHandler(chatService comm_services.Service) *WebSocketHandler {
	if chatService == nil {
		log.Fatal("WebSocketHandler: ChatService cannot be nil")
	}
	return &WebSocketHandler{
		ChatService: chatService,
	}
}

func (h *WebSocketHandler) ServeWSHandler(c *gin.Context) {
	token := c.GetHeader("Authorization")
	log.Printf("API: Received WebSocket connection attempt with token.")

	// 2. Validate token and extract acc_id
	claims, err := auth_services.ParseToken(token)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
		c.Abort()
		return
	}

	// 3. Upgrade HTTP connection to WebSocket connection
	// The upgrader.Upgrade method will write an HTTP error response to the client if
	// the upgrade fails (e.g., if CheckOrigin returns false or other handshake errors).
	wsConn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		// upgrader.Upgrade handles writing the HTTP error response, so no c.JSON here.
		log.Printf("API: Failed to upgrade WebSocket connection for acc_id %s: %v", claims.ID, err)
		// Gin might complain if we don't write anything, but Upgrade already did.
		// If you need to ensure Gin is satisfied, you could c.AbortWithError(status, err)
		// but it's usually fine as Upgrade handles the response.
		return
	}
	log.Printf("API: WebSocket connection successfully upgraded for acc_id: %s from remote: %s", claims.ID, wsConn.RemoteAddr())

	// 4. Pass the authenticated connection and client ID to the chat service
	// The chat service (Hub) will now manage this client.
	h.ChatService.HandleNewConnection(claims.ID, wsConn)

	// Note: Once upgraded, the HTTP handler's responsibility for this connection ends.
	// The WebSocket connection is now managed by the Client's ReadPump and WritePump goroutines
	// started by HandleNewConnection. Gin's context is no longer directly in control of the
	// persistent WebSocket communication.
}
