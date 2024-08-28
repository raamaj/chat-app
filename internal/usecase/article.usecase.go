package usecase

import (
	"context"
	"github.com/gofiber/fiber/v2"
	"github.com/raamaj/chat-app/internal/model"
	"github.com/raamaj/chat-app/internal/model/converter"
	"github.com/raamaj/chat-app/internal/repository"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type ArticleUseCase struct {
	Log        *logrus.Logger
	DB         *gorm.DB
	Repository repository.ArticleRepository
}

func NewArticleUseCase(log *logrus.Logger, db *gorm.DB, repo repository.ArticleRepository) *ArticleUseCase {
	return &ArticleUseCase{
		Log:        log,
		DB:         db,
		Repository: repo,
	}
}

func (a *ArticleUseCase) ListArticle(ctx context.Context) (*[]model.ArticleResponse, error) {
	tx := a.DB.WithContext(ctx).Begin()
	defer tx.Rollback()

	articles, err := a.Repository.List(tx)
	if err != nil {
		a.Log.Warnf("Failed Get List Article : %+v", err)
		return nil, fiber.ErrInternalServerError
	}

	if err = tx.Commit().Error; err != nil {
		a.Log.Warnf("Failed commit transaction : %+v", err)
		return nil, fiber.ErrInternalServerError
	}

	return converter.ArticleToResponseList(articles), err
}
