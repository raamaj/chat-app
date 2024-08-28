package converter

import (
	"github.com/raamaj/chat-app/internal/entity"
	"github.com/raamaj/chat-app/internal/model"
)

func ArticleToResponseList(entity *[]entity.Article) *[]model.ArticleResponse {
	var responses []model.ArticleResponse
	for _, article := range *entity {
		responses = append(responses, model.ArticleResponse{
			ID:           article.ID,
			Title:        article.Title,
			ImageUrl:     article.ImageUrl,
			Content:      article.Content,
			Author:       article.Author,
			LikeCount:    article.LikeCount,
			CommentCount: article.CommentCount,
			CreatedAt:    article.CreatedAt,
		})
	}

	return &responses
}
