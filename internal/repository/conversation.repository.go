package repository

import (
	"github.com/raamaj/chat-app/internal/entity"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

//go:generate mockgen -destination=mock/mock_conversation_repository.go -source=conversation.repository.go
type ConversationRepository interface {
	FindById(db *gorm.DB, conversation *entity.Conversation, id int64) error
	Create(db *gorm.DB, conversation *entity.Conversation) error
}

type conversationRepository struct {
	Log *logrus.Logger
}

func NewConversationRepository(log *logrus.Logger) *conversationRepository {
	return &conversationRepository{
		Log: log,
	}
}

func (c *conversationRepository) FindById(db *gorm.DB, conversation *entity.Conversation, id int64) error {
	return db.Where("id = ?", id).First(conversation).Error
}

func (c *conversationRepository) Create(db *gorm.DB, conversation *entity.Conversation) error {
	return db.Create(conversation).Error
}
