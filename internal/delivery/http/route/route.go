package route

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/gofiber/swagger"
	"github.com/raamaj/chat-app/internal/delivery/http"
	"github.com/redis/go-redis/v9"
)

type RouteConfig struct {
	App               *fiber.App
	UserController    *http.UserController
	MessageController *http.MessageController
	ArticleController *http.ArticleController
	AuthMiddleware    fiber.Handler
	Cache             *redis.Client
	JWTSecret         string
}

func (c *RouteConfig) Setup() {
	c.App.Use(cors.New())
	c.App.Use(recover.New())
	c.SetupGuestRoute()
}

func (c *RouteConfig) SetupGuestRoute() {
	c.App.Get("/api/docs/*", swagger.HandlerDefault)

	c.App.Post("/api/users", c.UserController.Register)
	c.App.Get("/api/users/:id", c.UserController.View)

	c.App.Post("/api/conversations/:conversationId/messages", c.MessageController.SendMessage)
	c.App.Get("/api/conversations/:conversationId/messages", c.MessageController.GetMessages)

	c.App.Get("/api/articles", c.ArticleController.GetArticles)
}
