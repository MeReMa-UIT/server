package models

import "time" // For mock repository structs

// --- Mock Repository Structs (to avoid direct dependency on actual repository package) ---
// These would normally be in your repository package. We define simplified versions here.
type Conversation struct {
	ConversationID int64      `json:"conversation_id"`
	AccID1         int64      `json:"acc_id_1"`
	AccID2         int64      `json:"patient_acc_id"` // Example field
	LastMessageAt  *time.Time `json:"last_message_at"`
}

type NewMessage struct {
	ConversationID int64  `json:"conversation_id"`
	Content        string `json:"content"`
}

type Message struct {
	NewMessage
	MessageID   int64     `json:"message_id"`
	SenderAccID int64     `json:"sender_acc_id"`
	IsRead      bool      `json:"is_read"`
	SentAt      time.Time `json:"sent_at"`
}

// InboundMessage represents messages coming from the client TO the server.
type InboundMessage struct {
	Type           string `json:"type"`                     // e.g., "sendMessage", "loadHistory", "startConversation"
	ConversationID int64  `json:"conversationId,omitempty"` // For sendMessage, loadHistory
	RecipientAccID int64  `json:"recipientAccId,omitempty"` // For startConversation (target user's acc_id as string)
	Text           string `json:"text,omitempty"`           // For sendMessage
	Limit          int    `json:"limit,omitempty"`          // For loadHistory
	Offset         int    `json:"offset,omitempty"`         // For loadHistory
}

// OutboundMessage represents messages going FROM the server TO the client.
type OutboundMessage struct {
	Type           string         `json:"type"`                     // e.g., "yourId", "newMessage", "messageHistory", "conversationList", "userList", "notification", "error", "conversationCreated"
	ID             string         `json:"id,omitempty"`             // For "yourId" (client's acc_id)
	Message        *Message       `json:"message,omitempty"`        // For "newMessage"
	Messages       []Message      `json:"messages,omitempty"`       // For "messageHistory"
	Conversations  []Conversation `json:"conversations,omitempty"`  // For "conversationList"
	Conversation   *Conversation  `json:"conversation,omitempty"`   // For "conversationCreated"
	UserList       []string       `json:"userList,omitempty"`       // For online user list (acc_ids as strings)
	Text           string         `json:"text,omitempty"`           // For "notification" or "error"
	ConversationID int64          `json:"conversationId,omitempty"` // Context for some messages like newMessage
	SenderIsSelf   bool           `json:"senderIsSelf,omitempty"`   // UI hint for "newMessage"
}

// --- End Mock Repository Structs ---
