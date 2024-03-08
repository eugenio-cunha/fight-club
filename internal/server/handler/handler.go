package handler

import (
	"fight-club/internal/domain"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

type Payload struct {
	Value       int    `json:"valor"`
	Kind        string `json:"tipo"`
	Description string `json:"descricao"`
}

func Balance(c *fiber.Ctx) error {
	idStr := c.Params("id")
	id, err := strconv.ParseInt(idStr, 10, 0)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).Send(nil)
	}

	if id < 1 || id > 5 {
		return c.Status(fiber.StatusNotFound).Send(nil)
	}

	output, err := domain.Balance(id)
	if err != nil {
		return c.Status(fiber.StatusUnprocessableEntity).Send(nil)
	}

	return c.Status(fiber.StatusOK).Send(*output)
}

func Transaction(c *fiber.Ctx) error {
	idStr := c.Params("id")
	id, err := strconv.ParseInt(idStr, 10, 0)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).Send(nil)
	}

	if id < 1 || id > 5 {
		return c.Status(fiber.StatusNotFound).Send(nil)
	}

	var payload Payload
	if err := c.BodyParser(&payload); err != nil {
		return c.Status(fiber.StatusBadRequest).Send(nil)
	}

	if payload.Kind != "C" && payload.Kind != "c" && payload.Kind != "D" && payload.Kind != "d" {
		return c.Status(fiber.StatusBadRequest).Send(nil)
	}

	if len(payload.Description) == 0 || len(payload.Description) > 10 {
		return c.Status(fiber.StatusBadRequest).Send(nil)
	}

	output, err := domain.Transaction(id, payload.Value, payload.Kind, payload.Description)
	if err != nil {
		return c.Status(fiber.StatusUnprocessableEntity).Send(nil)
	}

	return c.Status(fiber.StatusOK).Send(*output)
}
