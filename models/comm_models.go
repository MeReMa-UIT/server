package models

import "time"

type Conversation struct {
	ConversationID int64      `json:"conversation_id"`
	PartnerAccID   int64      `json:"partner_acc_id"`
	PartnerName    string     `json:"partner_name"`
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
	IsSeen      bool      `json:"is_seen"`
	SentAt      time.Time `json:"sent_at"`
}

type InboundMessage struct {
	Type           string     `json:"type"`                      //  "sendMessage", "loadHistory", "markSeenMessage"
	ConversationID int64      `json:"conversation_id,omitempty"` // For "sendMessage", "loadHistory", "markSeenMessage"
	PartnerAccID   int64      `json:"partner_acc_id,omitempty"`  // For "sendMessage", "markSeenMessage"
	Text           string     `json:"text,omitempty"`            // For "sendMessage"
	ReadTime       *time.Time `json:"read_time,omitempty"`       // For "markSeenMessage"
}

type OutboundMessage struct {
	Type           string         `json:"type"`                      // e.g., "yourID", "newMessage", "messageHistory", "conversationList",  "error"
	ID             string         `json:"id,omitempty"`              // For "yourID"
	Message        *Message       `json:"message,omitempty"`         // For "newMessage"
	Messages       []Message      `json:"messages,omitempty"`        // For "messageHistory"
	Conversations  []Conversation `json:"conversations,omitempty"`   // For "conversationList"
	Text           string         `json:"text,omitempty"`            // For "error"
	ConversationID int64          `json:"conversation_id,omitempty"` // for "newMessage", "seenMessage", "messageHistory"
	ReadTime       *time.Time     `json:"read_time,omitempty"`       // For "seenMessage"
}
