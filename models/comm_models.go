package models

import "time"

type ContactInfo struct {
	AccID    int    `json:"acc_id"`
	FullName string `json:"full_name"`
	Role     string `json:"role"`
}
type SendingMessage struct {
	ToAccID int    `json:"to_acc_id"`
	Content string `json:"content"`
}
type Message struct {
	Content  string    `json:"content"`
	SenderID int       `json:"sender_id"`
	SentAt   time.Time `json:"sent_at"`
}
