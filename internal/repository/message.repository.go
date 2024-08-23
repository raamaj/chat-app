package repository

import (
	"github.com/raamaj/chat-app/internal/entity"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

//go:generate mockgen -destination=mock/mock_message_repository.go -source=message.repository.go
type MessageRepository interface {
	List(db *gorm.DB, conversationID int64) (*[]entity.Message, error)
	Create(db *gorm.DB, message *entity.Message) error
}

type messageRepository struct {
	Log *logrus.Logger
}

func NewMessageRepository(log *logrus.Logger) *messageRepository {
	return &messageRepository{
		Log: log,
	}
}

func (r *messageRepository) List(db *gorm.DB, conversationID int64) (*[]entity.Message, error) {
	var messages []entity.Message
	rows, err := db.Model(&messages).Select("messages.id,users.id,conversations.id,messages.content,messages.send_at").
		Joins("inner join users on users.id = messages.sender_id").
		Joins("inner join conversations on conversations.id = messages.conversation_id").
		Where("messages.conversation_id = ?", conversationID).Rows()
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var message entity.Message
		err = rows.Scan(&message.ID, &message.SenderID, &message.ConversationID, &message.Content, &message.SendAt)
		if err != nil {
			return nil, err
		}

		messages = append(messages, message)
	}

	return &messages, nil
}

func (r *messageRepository) Create(db *gorm.DB, message *entity.Message) error {
	return db.Create(message).Error
}
