package entity

import (
	"encoding/json"
	"time"
)

type Conversation struct {
	ID        int64     `gorm:"column:id;primaryKey"`
	CreatedAt time.Time `gorm:"column:created_at;autoCreateTime:milli"`
}

func (c *Conversation) TableName() string {
	return "conversations"
}

func (c *Conversation) MarshalBinary() ([]byte, error) {
	p, err := json.Marshal(c)
	if err != nil {
		return nil, err
	}

	return p, nil
}

func (c *Conversation) UnmarshalBinary(data []byte) error {
	err := json.Unmarshal(data, &c)
	if err != nil {
		return err
	}

	return nil
}
