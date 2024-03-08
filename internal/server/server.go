package server

import (
	"fight-club/internal/server/router"

	"github.com/bytedance/sonic"
	"github.com/gofiber/fiber/v2"
)

func NewServer() *fiber.App {
	app := fiber.New(fiber.Config{
		AppName:       "Fight Club",
		Prefork:       true,
		CaseSensitive: true,
		JSONEncoder:   sonic.Marshal,
		JSONDecoder:   sonic.Unmarshal,
	})

	router.SetupRoutes(app)

	return app
}
