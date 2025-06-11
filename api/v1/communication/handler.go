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

// NewWebSocketHandler creates a new WebSocketHandler.
func NewWebSocketHandler(chatService comm_services.Service) *WebSocketHandler {
	if chatService == nil {
		log.Fatal("WebSocketHandler: ChatService cannot be nil")
	}
	return &WebSocketHandler{
		ChatService: chatService,
	}
}

// ServeWS is the Gin handler function for WebSocket connection requests.
// It expects a 'token' query parameter containing the JWT.
// Example URL: ws://localhost:8080/ws?token=YOUR_JWT_HERE
func (h *WebSocketHandler) ServeWS(c *gin.Context) {
	// 1. Extract JWT from query parameter
	tokenString := c.Query("token")
	if tokenString == "" {
		log.Println("API: WebSocket connection attempt without 'token' query parameter.")
		// For WebSockets, returning JSON before upgrade might not be seen by all clients.
		// It's often better to let the upgrader handle the error response if CheckOrigin fails,
		// or simply close the underlying connection if a token is mandatory for the upgrade attempt itself.
		// However, Gin expects a response to be written.
		c.JSON(http.StatusBadRequest, gin.H{"error": "Missing authentication token in query parameter"})
		c.Abort() // Abort further processing by Gin
		return
	}
	log.Printf("API: Received WebSocket connection attempt with token.")

	// 2. Validate token and extract acc_id
	claims, err := auth_services.ParseToken(tokenString)
	if err != nil {
		log.Printf("API: WebSocket token validation failed for token '%s': %v", tokenString, err)
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid or expired token"})
		c.Abort()
		return
	}
	log.Printf("API: WebSocket token validated. AccID: %s", claims.ID)

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
