package router

import (
	"github.com/gofiber/fiber/v2"

	"fight-club/internal/server/handler"
)

func SetupRoutes(app *fiber.App) {
	api := app.Group("/clientes")

	api.Get("/:id/extrato", handler.Balance)
	api.Post("/:id/transacoes", handler.Transaction)
}
