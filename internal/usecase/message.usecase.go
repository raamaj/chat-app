package usecase

import (
	"context"
	"errors"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/raamaj/chat-app/internal/entity"
	"github.com/raamaj/chat-app/internal/model"
	"github.com/raamaj/chat-app/internal/model/converter"
	"github.com/raamaj/chat-app/internal/repository"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
	"time"
)

type MessageUseCase struct {
	DB                     *gorm.DB
	Log                    *logrus.Logger
	Validate               *validator.Validate
	MessageRepository      repository.MessageRepository
	ConversationRepository repository.ConversationRepository
	UserRepository         repository.UserRepository
}

func NewMessageUseCase(db *gorm.DB, logger *logrus.Logger, validate *validator.Validate,
	messageRepo repository.MessageRepository, conRepo repository.ConversationRepository, userRepo repository.UserRepository) *MessageUseCase {
	return &MessageUseCase{
		DB:                     db,
		Log:                    logger,
		Validate:               validate,
		MessageRepository:      messageRepo,
		ConversationRepository: conRepo,
		UserRepository:         userRepo,
	}
}

func (muc *MessageUseCase) Create(ctx context.Context, request *model.MessageRequest) (*model.MessageResponse, error) {
	tx := muc.DB.WithContext(ctx).Begin()
	defer tx.Rollback()
	if err := muc.Validate.Struct(request); err != nil {
		muc.Log.Warnf("Invalid request body : %+v", err)
		return nil, fiber.ErrBadRequest
	}

	sender := &entity.User{}
	err := muc.UserRepository.FindById(tx, sender, request.SenderID)
	if err != nil {
		muc.Log.Warnf("User FindById Error : %+v", err)
		return nil, fiber.ErrInternalServerError
	}

	conversation := &entity.Conversation{}
	// Check existing conversation
	err = muc.ConversationRepository.FindById(tx, conversation, request.ConversationId)
	if err != nil {
		// Create if not exist
		conversation.CreatedAt = time.Now()
		err = muc.ConversationRepository.Create(tx, conversation)
		if err != nil {
			muc.Log.Warnf("Failed create conversation : %+v", err)
			return nil, fiber.ErrInternalServerError
		}
	}

	message := &entity.Message{
		Conversation: *conversation,
		Sender:       *sender,
		Content:      request.Content,
		SendAt:       time.Now(),
	}

	err = muc.MessageRepository.Create(tx, message)
	if err != nil {
		muc.Log.Warnf("Failed create conversation : %+v", err)
		return nil, fiber.ErrInternalServerError
	}

	if err = tx.Commit().Error; err != nil {
		muc.Log.Warnf("Failed commit transaction : %+v", err)
		return nil, fiber.ErrInternalServerError
	}

	return converter.MessageToResponse(message), nil
}

func (muc *MessageUseCase) List(ctx context.Context, conversationID int64) (*[]model.MessageResponse, error) {
	tx := muc.DB.WithContext(ctx).Begin()
	defer tx.Rollback()

	// Get Conversation
	conversation := &entity.Conversation{}
	err := muc.ConversationRepository.FindById(tx, conversation, conversationID)
	if err != nil {
		muc.Log.Warnf("Conversation FindById Error : %+v", err)
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, fiber.ErrNotFound
		}
		return nil, fiber.ErrInternalServerError
	}

	// Get Messages
	messages, err := muc.MessageRepository.List(tx, conversationID)
	if err != nil {
		muc.Log.Warnf("Failed list conversation : %+v", err)
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, fiber.ErrNotFound
		}
		return nil, fiber.ErrInternalServerError
	}

	if err = tx.Commit().Error; err != nil {
		muc.Log.Warnf("Failed commit transaction : %+v", err)
		return nil, fiber.ErrInternalServerError
	}

	var responses []model.MessageResponse
	for _, message := range *messages {
		responses = append(responses, *converter.MessageToResponse(&message))
	}

	return &responses, nil
}
