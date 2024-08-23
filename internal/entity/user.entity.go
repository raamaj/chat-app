package entity

import (
	"encoding/json"
	"time"
)

// User is a struct that represents a user entity
type User struct {
	ID        int64     `gorm:"column:id;primaryKey"`
	Username  string    `gorm:"column:username"`
	Email     string    `gorm:"column:email"`
	Password  string    `gorm:"column:password"`
	CreatedAt time.Time `gorm:"column:created_at;autoCreateTime:milli"`
}

func (u *User) TableName() string {
	return "users"
}

func (u *User) MarshalBinary() ([]byte, error) {
	p, err := json.Marshal(u)
	if err != nil {
		return nil, err
	}

	return p, nil
}

func (u *User) UnmarshalBinary(data []byte) error {
	err := json.Unmarshal(data, &u)
	if err != nil {
		return err
	}

	return nil
}
