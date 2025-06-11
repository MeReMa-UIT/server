package comm_services

import "github.com/gorilla/websocket" // Or your chosen WebSocket library

// Service defines the interface for the core chat operations.
// The Hub struct will implement this interface.
type Service interface {
	Run() // Starts the service's main event loop (e.g., Hub's run loop).

	// HandleNewConnection is called by the API layer when a new WebSocket connection
	// has been successfully established and authenticated.
	// clientID is the authenticated user's account ID (e.g., from JWT).
	// conn is the raw WebSocket connection.
	HandleNewConnection(clientID string, conn *websocket.Conn)
}

// Ensure Hub implements Service (compile-time check).
var _ Service = (*Hub)(nil)

// HandleNewConnection is the Hub's implementation of the Service interface method.
// It creates a new Client, starts its read/write pumps, and registers it with the Hub.
func (h *Hub) HandleNewConnection(clientID string, conn *websocket.Conn) {
	// Create a new Client instance for this connection.
	// clientID is the authenticated acc_id (string).
	client := NewClient(clientID, h, conn)

	// Register the client with the hub. This will send it to the hub's registerCh.
	h.Register(client) // This is the Hub's own Register method, not a recursive call to this one.

	// Start the goroutines for reading from and writing to the WebSocket connection.
	// These run for the lifetime of the client's connection.
	go client.WritePump()
	go client.ReadPump()
}
