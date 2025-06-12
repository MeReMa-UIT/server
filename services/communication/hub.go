package comm_services

import (
	"encoding/json"
	"log"
	"strconv"
	"sync"

	"github.com/merema-uit/server/models"
	"github.com/merema-uit/server/repo"
)

type Hub struct {
	clients      map[string]*Client // Map client acc_id (string) to Client struct
	clientsMux   sync.RWMutex
	registerCh   chan *Client
	unregisterCh chan *Client
	processMsgCh chan *ClientMessageEnvelope // Receives messages from all clients via ClientMessageEnvelope
}

func NewHub() *Hub {
	return &Hub{
		clients:      make(map[string]*Client),
		registerCh:   make(chan *Client),
		unregisterCh: make(chan *Client),
		processMsgCh: make(chan *ClientMessageEnvelope),
	}
}

func (h *Hub) Run() {
	log.Println("Chat Hub started and listening for events.")
	for {
		select {
		case client := <-h.registerCh:
			h.handleRegister(client)
		case client := <-h.unregisterCh:
			h.handleUnregister(client)
		case envelope := <-h.processMsgCh:
			h.handleProcessClientMessage(envelope)
		}
	}
}

func (h *Hub) handleRegister(client *Client) {
	h.clientsMux.Lock()
	h.clients[client.ID] = client
	h.clientsMux.Unlock()
	// utils.Logger.Info("Hub: Client registered: %s. Total clients: %d", client.ID, len(h.clients))

	idMsg := models.OutboundMessage{Type: "yourID", ID: client.ID}
	h.sendToClient(client, idMsg)
	h.sendUserConversations(client)
	// h.broadcastOnlineUserList()
}

func (h *Hub) handleUnregister(client *Client) {
	h.clientsMux.Lock()
	if _, ok := h.clients[client.ID]; ok {
		delete(h.clients, client.ID)
		close(client.Send) // Important: Close the send channel to signal WritePump to exit
		// utils.Logger.Info("Hub: Client unregistered: %s. Total clients: %d", client.ID, len(h.clients))
	}
	h.clientsMux.Unlock()
	// h.broadcastOnlineUserList() // Notify remaining clients
}

func (h *Hub) sendUserConversations(client *Client) {
	userConvos, err := repo.GetConversationList(client.Ctx, client.ID)

	if err != nil {
		log.Printf("Hub: Error fetching conversation list for client %s: %v", client.ID, err)
		h.sendErrorToClient(client, "Failed to load conversations.")
		return
	}
	outMsg := models.OutboundMessage{Type: "conversationList", Conversations: userConvos}
	h.sendToClient(client, outMsg)
}

// func (h *Hub) broadcastOnlineUserList() {
// 	h.clientsMux.RLock()
// 	onlineUserIDs := make([]string, 0, len(h.clients))
// 	for id := range h.clients {
// 		onlineUserIDs = append(onlineUserIDs, id)
// 	}
// 	h.clientsMux.RUnlock()

// 	outMsg := models.OutboundMessage{Type: "userList", UserList: onlineUserIDs}
// 	h.broadcastToAll(outMsg)
// }

func (h *Hub) handleProcessClientMessage(envelope *ClientMessageEnvelope) {
	var inboundMsg models.InboundMessage
	if err := json.Unmarshal(envelope.RawMessage, &inboundMsg); err != nil {
		log.Printf("Hub: Error unmarshalling message from %s: %v", envelope.Sender.ID, err)
		h.sendErrorToClient(envelope.Sender, "Invalid message format.")
		return
	}

	log.Printf("Hub: Received type '%s' from %s: %+v", inboundMsg.Type, envelope.Sender.ID, inboundMsg)

	switch inboundMsg.Type {
	case "sendMessage":
		h.handleSendMessage(envelope.Sender, inboundMsg)
	case "loadHistory":
		h.handleLoadHistory(envelope.Sender, inboundMsg)
	case "markSeenMessage":
		h.handleMarkSeenMessage(envelope.Sender, inboundMsg)
	default:
		h.sendErrorToClient(envelope.Sender, "Unknown message type: "+inboundMsg.Type)
	}
}

func (h *Hub) handleSendMessage(sender *Client, msg models.InboundMessage) {
	if msg.ConversationID == 0 || msg.Text == "" {
		h.sendErrorToClient(sender, "Missing conversationId or text for sendMessage.")
		return
	}

	dbMsg := &models.NewMessage{
		ConversationID: msg.ConversationID,
		Content:        msg.Text,
	}

	senderAccID, _ := strconv.ParseInt(sender.ID, 10, 64)
	if msg.PartnerAccID == senderAccID {
		log.Printf("Hub: Sender %s tried to send a message to themselves in conversation %d.", sender.ID, msg.ConversationID)
		return
	}

	storedMessage, err := repo.StoreMessage(sender.Ctx, *dbMsg, sender.ID)

	if err != nil {
		log.Printf("Hub: Error storing message for client %s: %v", sender.ID, err)
		h.sendErrorToClient(sender, "Failed to send message.")
		return
	}

	err = repo.UpdateConversationLastMessage(sender.Ctx, storedMessage.ConversationID, storedMessage.SentAt)

	if err != nil {
		log.Printf("Hub: Error updating last message time for conversation %d: %v", msg.ConversationID, err)
		h.sendErrorToClient(sender, "Failed to update conversation last message time.")
		return
	}

	outboundToSender := models.OutboundMessage{Type: "newMessage", Message: &storedMessage, ConversationID: msg.ConversationID}
	h.sendToClient(sender, outboundToSender)

	h.clientsMux.RLock()
	recipientClient, isRecipientOnline := h.clients[strconv.FormatInt(msg.PartnerAccID, 10)]
	h.clientsMux.RUnlock()

	if isRecipientOnline {
		outboundToRecipient := models.OutboundMessage{Type: "newMessage", Message: &storedMessage, ConversationID: msg.ConversationID}
		h.sendToClient(recipientClient, outboundToRecipient)
	} else {
		log.Printf("Hub: Recipient %d for message in convo %d is offline.", msg.PartnerAccID, msg.ConversationID)
		// Here you might implement push notifications or unread counts later
	}
}

func (h *Hub) handleLoadHistory(sender *Client, msg models.InboundMessage) {
	if msg.ConversationID == 0 {
		h.sendErrorToClient(sender, "Missing conversation ID for load history.")
		return
	}
	limit := msg.Limit
	if limit <= 0 || limit > 50 {
		limit = 20 // Default/max limit
	}
	offset := msg.Offset
	if offset < 0 {
		offset = 0
	}

	convoMessages, err := repo.GetConversationMessage(sender.Ctx, strconv.FormatInt(msg.ConversationID, 10))

	if err != nil {
		log.Printf("Hub: Error fetching messages for conversation %d: %v", msg.ConversationID, err)
		h.sendErrorToClient(sender, "Failed to load message history.")
		return
	}

	// Verify sender is part of this conversation (simplified)
	// In a real scenario, you'd check convo.DoctorAccID or convo.PatientAccID against sender.ID

	messagesSlice := []models.Message{}
	totalMessages := len(convoMessages)

	// Paginate from the end of the list (most recent)
	start := totalMessages - offset - limit
	end := totalMessages - offset

	if start < 0 {
		start = 0
	}
	if end < 0 { // Should not happen if offset is reasonable
		end = 0
	}
	if end > totalMessages {
		end = totalMessages
	}
	if start >= end { // No messages in range or invalid range
		// send empty list
	} else {
		for i := start; i < end; i++ {
			messagesSlice = append(messagesSlice, convoMessages[i]) // Dereference and copy
		}
	}

	outMsg := models.OutboundMessage{Type: "messageHistory", Messages: messagesSlice, ConversationID: msg.ConversationID}
	h.sendToClient(sender, outMsg)
}

// // rehandling
// func (h *Hub) handleStartConversation(sender *Client, msg models.InboundMessage) {
// 	senderAccID, _ := strconv.ParseInt(sender.ID, 10, 64)
// 	recipientAccID := msg.RecipientAccID

// 	if senderAccID == recipientAccID {
// 		h.sendErrorToClient(sender, "Cannot start a conversation with yourself.")
// 		return
// 	}

// 	// Find existing or create new conversation (simplified for doctor/patient)
// 	// For simplicity, assume sender is Patient, recipient is Doctor for pair uniqueness.
// 	// A real app needs robust role handling.
// 	accID1, accID2 := senderAccID, recipientAccID
// 	if accID1 > accID2 {
// 		accID1, accID2 = accID2, accID1
// 	}

// 	var existingConvo *models.Conversation

// 	// Notify sender
// 	outMsgSender := models.OutboundMessage{Type: "conversationCreated", Conversation: existingConvo}
// 	h.sendToClient(sender, outMsgSender)

//		// Notify recipient if they are online
//		h.clientsMux.RLock()
//		recipientClient, isRecipientOnline := h.clients[strconv.FormatInt(recipientAccID, 10)]
//		h.clientsMux.RUnlock()
//		if isRecipientOnline {
//			outMsgRecipient := models.OutboundMessage{Type: "conversationCreated", Conversation: existingConvo}
//			h.sendToClient(recipientClient, outMsgRecipient)
//		}
//	}

func (h *Hub) handleMarkSeenMessage(client *Client, msg models.InboundMessage) {
	if msg.ConversationID == 0 || msg.ReadTime == nil {
		h.sendErrorToClient(client, "Missing conversation ID or read time for mark seen message.")
		return
	}

	// Update the message seen status in the repository
	err := repo.UpdateMessageSeenStatus(client.Ctx, msg.ConversationID, msg.PartnerAccID, *msg.ReadTime)

	if err != nil {
		log.Printf("Hub: Error marking messages as seen for client %s in conversation %d: %v", client.ID, msg.ConversationID, err)
		h.sendErrorToClient(client, "Failed to mark messages as seen.")
		return
	}

	outMsg := models.OutboundMessage{Type: "seenMessage", ConversationID: msg.ConversationID, ReadTime: msg.ReadTime}
	h.sendToClient(client, outMsg)
	log.Printf("Hub: Client %s marked messages as seen in conversation %d at %s", client.ID, msg.ConversationID, msg.ReadTime)

	h.clientsMux.RLock()
	recipientClient, isRecipientOnline := h.clients[strconv.FormatInt(msg.PartnerAccID, 10)]
	h.clientsMux.RUnlock()

	if isRecipientOnline {
		h.sendToClient(recipientClient, outMsg)
	} else {
		log.Printf("Hub: Recipient %d for message in convo %d is offline.", msg.PartnerAccID, msg.ConversationID)
		// Here you might implement push notifications or unread counts later
	}
}

// --- Helper methods for sending messages ---
func (h *Hub) sendToClient(client *Client, msg models.OutboundMessage) {
	jsonData, err := json.Marshal(msg)
	if err != nil {
		log.Printf("Hub: Error marshalling message for client %s: %v", client.ID, err)
		return
	}
	select {
	case client.Send <- jsonData:
	default: // If client's send buffer is full, log and potentially disconnect
		log.Printf("Hub: Client %s send channel full. Dropping message of type %s.", client.ID, msg.Type)
		// Consider closing client.Send here or letting ReadPump/WritePump handle timeouts
	}
}

// func (h *Hub) broadcastToAll(msg models.OutboundMessage) {
// 	jsonData, err := json.Marshal(msg)
// 	if err != nil {
// 		log.Printf("Hub: Error marshalling broadcast message: %v", err)
// 		return
// 	}
// 	h.clientsMux.RLock()
// 	for _, client := range h.clients {
// 		select {
// 		case client.Send <- jsonData:
// 		default:
// 			log.Printf("Hub: Broadcast to client %s send channel full. Dropping message.", client.ID)
// 		}
// 	}
// 	h.clientsMux.RUnlock()
// }

func (h *Hub) sendErrorToClient(client *Client, errorText string) {
	errMsg := models.OutboundMessage{Type: "error", Text: errorText}
	h.sendToClient(client, errMsg)
}

// --- Service Interface Methods ---
// These methods allow the API layer (or other parts of the system) to interact with the Hub.

// Register is called by the API layer when a new WebSocket connection is established and authenticated.
func (h *Hub) Register(client *Client) {
	h.registerCh <- client
}

// Unregister is called by a Client's ReadPump when its connection closes.
func (h *Hub) Unregister(client *Client) {
	h.unregisterCh <- client
}

// ProcessMessage is called by a Client's ReadPump to forward a message to the Hub.
// This is not typically called directly from outside, but is part of the internal flow.
// The API layer would call HandleNewConnection, which creates the client and starts its pumps.
// func (h *Hub) ProcessMessage(envelope *ClientMessageEnvelope) {
// 	h.processMsgCh <- envelope
// }
