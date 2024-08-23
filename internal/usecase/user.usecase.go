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
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	"time"
)

type UserUseCase struct {
	DB             *gorm.DB
	Log            *logrus.Logger
	Validate       *validator.Validate
	UserRepository repository.UserRepository
}

func NewUserUseCase(db *gorm.DB, logger *logrus.Logger, validate *validator.Validate,
	userRepository repository.UserRepository) *UserUseCase {
	return &UserUseCase{
		DB:             db,
		Log:            logger,
		Validate:       validate,
		UserRepository: userRepository,
	}
}

func (c *UserUseCase) Create(ctx context.Context, request *model.RegisterUserRequest) (*model.UserResponse, error) {
	tx := c.DB.WithContext(ctx).Begin()
	defer tx.Rollback()

	err := c.Validate.Struct(request)
	if err != nil {
		c.Log.Warnf("Invalid request body : %+v", err)
		return nil, fiber.ErrBadRequest
	}

	total, err := c.UserRepository.CountByUsername(tx, request.Username)
	if err != nil {
		c.Log.Warnf("Failed count user from database : %+v", err)
		return nil, fiber.ErrInternalServerError
	}

	if total > 0 {
		c.Log.Warnf("User already exists : %+v", err)
		return nil, fiber.ErrConflict
	}

	password, err := bcrypt.GenerateFromPassword([]byte(request.Password), bcrypt.DefaultCost)
	if err != nil {
		c.Log.Warnf("Failed to generate bcrype hash : %+v", err)
		return nil, fiber.ErrInternalServerError
	}

	user := &entity.User{
		Password:  string(password),
		Username:  request.Username,
		Email:     request.Email,
		CreatedAt: time.Now(),
	}

	if err := c.UserRepository.Create(tx, user); err != nil {
		c.Log.Warnf("Failed create user to database : %+v", err)
		return nil, fiber.ErrInternalServerError
	}

	if err := tx.Commit().Error; err != nil {
		c.Log.Warnf("Failed commit transaction : %+v", err)
		return nil, fiber.ErrInternalServerError
	}

	return converter.UserToResponse(user), nil
}

func (c *UserUseCase) View(ctx context.Context, id int64) (*model.UserResponse, error) {
	tx := c.DB.WithContext(ctx).Begin()
	defer tx.Rollback()

	user := &entity.User{}

	err := c.UserRepository.FindById(tx, user, id)
	if err != nil {
		c.Log.Warnf("Failed find user by id : %+v", err)
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, fiber.ErrNotFound
		}
		return nil, fiber.ErrInternalServerError
	}

	if err := tx.Commit().Error; err != nil {
		c.Log.Warnf("Failed commit transaction : %+v", err)
		return nil, fiber.ErrInternalServerError
	}

	return converter.UserToResponse(user), nil
}
