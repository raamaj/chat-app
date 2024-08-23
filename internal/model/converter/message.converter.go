package converter

import (
	"github.com/raamaj/chat-app/internal/entity"
	"github.com/raamaj/chat-app/internal/model"
)

func MessageToResponse(message *entity.Message) *model.MessageResponse {
	return &model.MessageResponse{
		ID:             message.ID,
		ConversationID: message.ConversationID,
		SenderID:       message.SenderID,
		Content:        message.Content,
		SentAt:         message.SendAt,
	}
}
