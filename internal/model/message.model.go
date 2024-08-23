package model

import "time"

type MessageRequest struct {
	ConversationId int64  `json:"-"`
	SenderID       int64  `json:"sender_id"`
	Content        string `json:"content"`
}

type MessageResponse struct {
	ID             int64     `json:"id"`
	ConversationID int64     `json:"conversation_id"`
	SenderID       int64     `json:"sender_id"`
	Content        string    `json:"content"`
	SentAt         time.Time `json:"sent_at"`
}
