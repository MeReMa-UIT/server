package comm_services

import (
	"context"
	"log"
	"time"

	"github.com/gorilla/websocket" // Assuming this is still the chosen WebSocket library
)

const (
	writeWait      = 10 * time.Second
	pongWait       = 60 * time.Second
	pingPeriod     = (pongWait * 9) / 10
	maxMessageSize = 1024
)

// ClientMessageEnvelope wraps a raw message with its sender Client pointer.
// This is used internally within the hub to pass messages from Client.ReadPump to Hub.processMsgCh.
type ClientMessageEnvelope struct {
	RawMessage []byte
	Sender     *Client // Pointer to the chat.Client that sent this message
}

// Client is a middleman between the websocket connection and the hub.
type Client struct {
	ID   string          // User's account ID (string, from JWT)
	Hub  *Hub            // Reference to the Hub that manages this client
	Conn *websocket.Conn // The WebSocket connection itself
	Send chan []byte     // Buffered channel of outbound messages for this client
	Ctx  context.Context
}

// NewClient creates a new Client instance.
func NewClient(id string, hub *Hub, conn *websocket.Conn) *Client {
	return &Client{
		ID:   id,
		Hub:  hub,
		Conn: conn,
		Send: make(chan []byte, 256),
		Ctx:  context.Background(), // Initialize with a background context
	}
}

// ReadPump pumps messages from the WebSocket connection to the hub.
func (c *Client) ReadPump() {
	defer func() {
		c.Hub.Unregister(c)
		c.Conn.Close()
		log.Printf("Client %s: ReadPump closed, connection terminated.", c.ID)
	}()
	c.Conn.SetReadLimit(maxMessageSize)
	if err := c.Conn.SetReadDeadline(time.Now().Add(pongWait)); err != nil {
		log.Printf("Client %s: Error setting read deadline in ReadPump: %v", c.ID, err)
		return
	}
	c.Conn.SetPongHandler(func(string) error {
		return c.Conn.SetReadDeadline(time.Now().Add(pongWait))
	})

	for {
		_, message, err := c.Conn.ReadMessage()
		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure, websocket.CloseNoStatusReceived) {
				log.Printf("Client %s: ReadPump WebSocket unexpected close error: %v", c.ID, err)
			} else {
				log.Printf("Client %s: ReadPump WebSocket read error (or normal close): %v", c.ID, err)
			}
			break
		}
		envelope := &ClientMessageEnvelope{RawMessage: message, Sender: c}
		c.Hub.processMsgCh <- envelope // Send to Hub's central processing channel
	}
}

// WritePump pumps messages from the hub to the WebSocket connection.
func (c *Client) WritePump() {
	ticker := time.NewTicker(pingPeriod)
	defer func() {
		ticker.Stop()
		c.Conn.Close()
		log.Printf("Client %s: WritePump closed.", c.ID)
	}()
	for {
		select {
		case message, ok := <-c.Send:
			if err := c.Conn.SetWriteDeadline(time.Now().Add(writeWait)); err != nil {
				log.Printf("Client %s: Error setting write deadline in WritePump: %v", c.ID, err)
				return
			}
			if !ok {
				log.Printf("Client %s: Hub closed send channel, sending close message.", c.ID)
				_ = c.Conn.WriteMessage(websocket.CloseMessage, []byte{})
				return
			}
			if err := c.Conn.WriteMessage(websocket.TextMessage, message); err != nil {
				log.Printf("Client %s: Error writing message in WritePump: %v", c.ID, err)
				return
			}
		case <-ticker.C:
			if err := c.Conn.SetWriteDeadline(time.Now().Add(writeWait)); err != nil {
				log.Printf("Client %s: Error setting write deadline for ping: %v", c.ID, err)
				return
			}
			if err := c.Conn.WriteMessage(websocket.PingMessage, nil); err != nil {
				log.Printf("Client %s: Error sending ping in WritePump: %v", c.ID, err)
				return
			}
		}
	}
}
