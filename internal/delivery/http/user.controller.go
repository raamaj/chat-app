package http

import (
	"github.com/gofiber/fiber/v2"
	"github.com/raamaj/chat-app/internal/model"
	"github.com/raamaj/chat-app/internal/usecase"
	"github.com/sirupsen/logrus"
	"strconv"
)

type UserController struct {
	Log     *logrus.Logger
	UseCase *usecase.UserUseCase
}

func NewUserController(useCase *usecase.UserUseCase, logger *logrus.Logger) *UserController {
	return &UserController{
		Log:     logger,
		UseCase: useCase,
	}
}

// Register godoc
// @Summary      Create a new user.
// @Description  Create a new user.
// @Tags         Users
// @Accept       json
// @Produce      json
// @Param		 request body   model.RegisterUserRequest true "User Data"
// @Success      200  {object}  model.WebResponse[model.UserResponse]
// @Failure      400  {object}  model.ErrorResponse "Bad Request"
// @Failure      500  {object}  model.ErrorResponse "Internal Server Error"
// @Router       /users [post]
func (c *UserController) Register(ctx *fiber.Ctx) error {
	request := new(model.RegisterUserRequest)
	err := ctx.BodyParser(request)
	if err != nil {
		c.Log.Warnf("Failed to parse request body : %+v", err)
		return fiber.ErrBadRequest
	}

	response, err := c.UseCase.Create(ctx.UserContext(), request)
	if err != nil {
		return err
	}

	return ctx.JSON(model.WebResponse[*model.UserResponse]{Data: response})
}

// View godoc
// @Summary      Retrieve user details by ID.
// @Description  Retrieve user details by ID.
// @Tags         Users
// @Accept       json
// @Produce      json
// @Param		 id path string true "User ID"
// @Success      200  {object}  model.WebResponse[model.UserResponse]
// @Failure      400  {object}  model.ErrorResponse "Bad Request"
// @Failure      500  {object}  model.ErrorResponse "Internal Server Error"
// @Router       /users/{id} [get]
func (c *UserController) View(ctx *fiber.Ctx) error {
	userId := ctx.Params("id")
	userIdInt, err := strconv.Atoi(userId)
	if err != nil {
		c.Log.Warnf("Failed to convert userId to int : %+v", err)
		return fiber.ErrInternalServerError
	}

	response, err := c.UseCase.View(ctx.Context(), int64(userIdInt))
	if err != nil {
		return err
	}

	return ctx.JSON(model.WebResponse[*model.UserResponse]{Data: response})
}
