package handler

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
)

func (h *Handler) DeleteUser(c *fiber.Ctx) error {
	id := c.Params("id")

	err := h.Removing.RemoveUserByID(c.Context(), id)

	if err != nil {
		return c.Status(http.StatusNotFound).JSON(fiber.Map{"message": "user with id " + id + " does not exist"})
	}

	return c.Status(http.StatusOK).JSON(fiber.Map{"message": "user was deleted successfully"})
}
