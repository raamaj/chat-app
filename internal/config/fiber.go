package config

import (
	"github.com/gofiber/fiber/v2"
	"github.com/raamaj/chat-app/internal/util"
	"github.com/spf13/viper"
)

func NewFiber(config *viper.Viper) *fiber.App {
	var app = fiber.New(fiber.Config{
		AppName:      config.GetString("app.name"),
		ErrorHandler: util.NewErrorHandler(),
		Prefork:      config.GetBool("web.prefork"),
	})

	return app
}
