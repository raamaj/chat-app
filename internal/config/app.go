package config

import (
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/raamaj/chat-app/internal/delivery/http"
	"github.com/raamaj/chat-app/internal/delivery/http/route"
	"github.com/raamaj/chat-app/internal/repository"
	"github.com/raamaj/chat-app/internal/usecase"
	"github.com/redis/go-redis/v9"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"gorm.io/gorm"
)

type BootstrapConfig struct {
	DB       *gorm.DB
	App      *fiber.App
	Log      *logrus.Logger
	Validate *validator.Validate
	Config   *viper.Viper
	Cache    *redis.Client
}

func Bootstrap(config *BootstrapConfig) {
	// setup repositories
	userRepository := repository.NewUserRepository(config.Log)
	messageRepository := repository.NewMessageRepository(config.Log)
	conversationRepository := repository.NewConversationRepository(config.Log)
	articleRepository := repository.NewArticleRepository(config.Log)

	// setup use cases
	userUseCase := usecase.NewUserUseCase(config.DB, config.Log, config.Validate, userRepository)
	messageUseCase := usecase.NewMessageUseCase(config.DB, config.Log, config.Validate, messageRepository, conversationRepository, userRepository)
	articleUseCase := usecase.NewArticleUseCase(config.Log, config.DB, articleRepository)

	// setup controller
	userController := http.NewUserController(userUseCase, config.Log)
	messageController := http.NewMessageController(messageUseCase, config.Log)
	articleController := http.NewArticleController(config.Log, articleUseCase)

	routeConfig := route.RouteConfig{
		App:               config.App,
		UserController:    userController,
		MessageController: messageController,
		ArticleController: articleController,
		Cache:             config.Cache,
		JWTSecret:         config.Config.GetString("jwt.secret"),
	}
	routeConfig.Setup()
}
