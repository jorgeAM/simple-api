package handler

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
)

func (h *Handler) GetUsers(c *fiber.Ctx) error {
	users, err := h.Retrieving.GetAllUser(c.Context())

	if err != nil {
		return c.Status(http.StatusNotFound).JSON(fiber.Map{"message": err.Error()})
	}

	return c.Status(http.StatusOK).JSON(fiber.Map{"users": users})
}
