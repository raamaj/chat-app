package entity

import (
	"encoding/json"
	"time"
)

type Message struct {
	ID             int64 `gorm:"column:id;primaryKey"`
	ConversationID int64
	Conversation   Conversation `gorm:"foreignKey:ConversationID;references:ID"`
	SenderID       int64
	Sender         User      `gorm:"foreignKey:SenderID;references:ID"`
	Content        string    `gorm:"column:content"`
	SendAt         time.Time `gorm:"column:send_at;autoCreateTime:milli"`
}

func (m *Message) TableName() string {
	return "messages"
}

func (m *Message) MarshalBinary() ([]byte, error) {
	p, err := json.Marshal(m)
	if err != nil {
		return nil, err
	}

	return p, nil
}

func (m *Message) UnmarshalBinary(data []byte) error {
	err := json.Unmarshal(data, &m)
	if err != nil {
		return err
	}

	return nil
}
