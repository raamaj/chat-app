package repository

import (
	"github.com/raamaj/chat-app/internal/entity"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type ArticleRepository interface {
	List(db *gorm.DB) (*[]entity.Article, error)
}

type articleRepository struct {
	Log *logrus.Logger
}

func NewArticleRepository(log *logrus.Logger) *articleRepository {
	return &articleRepository{
		Log: log,
	}
}

func (arr *articleRepository) List(db *gorm.DB) (*[]entity.Article, error) {
	var articles []entity.Article
	err := db.Find(&articles).Error
	return &articles, err
}
