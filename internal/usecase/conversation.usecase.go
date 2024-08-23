package usecase

import (
	"context"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/raamaj/chat-app/internal/entity"
	"github.com/raamaj/chat-app/internal/repository"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
	"time"
)

type ConversationUseCase struct {
	DB                     *gorm.DB
	Log                    *logrus.Logger
	Validate               *validator.Validate
	ConversationRepository repository.ConversationRepository
}

func NewConversationUseCase(db *gorm.DB, log *logrus.Logger, validate *validator.Validate, repo repository.ConversationRepository) *ConversationUseCase {
	return &ConversationUseCase{
		DB:                     db,
		Log:                    log,
		Validate:               validate,
		ConversationRepository: repo,
	}
}

func (uc *ConversationUseCase) Create(ctx context.Context, id int64) (*entity.Conversation, error) {
	tx := uc.DB.WithContext(ctx).Begin()
	defer tx.Rollback()

	conversation := &entity.Conversation{}

	// Check existing conversation
	err := uc.ConversationRepository.FindById(tx, conversation, id)
	if err != nil {
		// Create if not exist
		conversation.CreatedAt = time.Now()
		err = uc.ConversationRepository.Create(tx, conversation)
		if err != nil {
			uc.Log.Warnf("Failed create conversation : %+v", err)
			return nil, fiber.ErrInternalServerError
		}
	}

	if err = tx.Commit().Error; err != nil {
		uc.Log.Warnf("Failed commit transaction : %+v", err)
		return nil, fiber.ErrInternalServerError
	}

	return conversation, nil
}
