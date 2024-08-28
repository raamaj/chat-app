package http

import (
	"github.com/gofiber/fiber/v2"
	"github.com/raamaj/chat-app/internal/model"
	"github.com/raamaj/chat-app/internal/usecase"
	"github.com/sirupsen/logrus"
)

type ArticleController struct {
	Log     *logrus.Logger
	UseCase *usecase.ArticleUseCase
}

func NewArticleController(log *logrus.Logger, useCase *usecase.ArticleUseCase) *ArticleController {
	return &ArticleController{
		Log:     log,
		UseCase: useCase,
	}
}

func (c *ArticleController) GetArticles(ctx *fiber.Ctx) error {
	articles, err := c.UseCase.ListArticle(ctx.Context())
	if err != nil {
		return err
	}

	return ctx.JSON(model.WebResponse[*[]model.ArticleResponse]{Data: articles})
}
