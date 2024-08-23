package converter

import (
	"github.com/raamaj/chat-app/internal/entity"
	"github.com/raamaj/chat-app/internal/model"
)

func UserToResponse(user *entity.User) *model.UserResponse {
	return &model.UserResponse{
		ID:        user.ID,
		Username:  user.Username,
		Email:     user.Email,
		CreatedAt: user.CreatedAt,
	}
}
