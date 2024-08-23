package http

import (
	"github.com/gofiber/fiber/v2"
	"github.com/raamaj/chat-app/internal/model"
	"github.com/raamaj/chat-app/internal/usecase"
	"github.com/sirupsen/logrus"
	"strconv"
)

type MessageController struct {
	Log     *logrus.Logger
	UseCase *usecase.MessageUseCase
}

func NewMessageController(useCase *usecase.MessageUseCase, log *logrus.Logger) *MessageController {
	return &MessageController{
		Log:     log,
		UseCase: useCase,
	}
}

// SendMessage godoc
// @Summary      Send a message in a conversation.
// @Description  Send a message in a conversation.
// @Tags         Message
// @Accept       json
// @Produce      json
// @Param		 request body   model.MessageRequest true "Message Request"
// @Param		 conversationId path string true "Conversation ID"
// @Success      200  {object}  model.WebResponse[model.MessageResponse]
// @Failure      400  {object}  model.ErrorResponse "Bad Request"
// @Failure      500  {object}  model.ErrorResponse "Internal Server Error"
// @Router       /conversations/{conversationId}/messages [post]
func (mc *MessageController) SendMessage(ctx *fiber.Ctx) error {
	conversationIdStr := ctx.Params("conversationId")
	conversationId, _ := strconv.Atoi(conversationIdStr)

	request := new(model.MessageRequest)
	err := ctx.BodyParser(request)
	if err != nil {
		mc.Log.Warnf("Failed to parse request body : %+v", err)
		return fiber.ErrBadRequest
	}

	request.ConversationId = int64(conversationId)

	response, err := mc.UseCase.Create(ctx.Context(), request)
	if err != nil {
		return err
	}

	return ctx.JSON(model.WebResponse[*model.MessageResponse]{Data: response})
}

// GetMessages godoc
// @Summary      Retrieve all messages in a conversation.
// @Description  Retrieve all messages in a conversation.
// @Tags         Message
// @Accept       json
// @Produce      json
// @Param		 conversationId path string true "Conversation ID"
// @Success      200  {object}  model.WebResponse[model.MessageResponse]
// @Failure      400  {object}  model.ErrorResponse "Bad Request"
// @Failure      500  {object}  model.ErrorResponse "Internal Server Error"
// @Router       /conversations/{conversationId}/messages [get]
func (mc *MessageController) GetMessages(ctx *fiber.Ctx) error {
	conversationIdStr := ctx.Params("conversationId")
	conversationId, _ := strconv.Atoi(conversationIdStr)

	response, err := mc.UseCase.List(ctx.Context(), int64(conversationId))
	if err != nil {
		return err
	}

	return ctx.JSON(model.WebResponse[*[]model.MessageResponse]{Data: response})
}
