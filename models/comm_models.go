package models

import "time"

type ContactInfo struct {
	AccID    int    `json:"acc_id"`
	FullName string `json:"full_name"`
}
type SendingMessage struct {
	ToAccID int    `json:"to_acc_id"`
	Content string `json:"content"`
}
type Message struct {
	Content string    `json:"content"`
	SentAt  time.Time `json:"sent_at"`
}
