package handler

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/jorgeAM/simple-api/internal/user/application/removing"
)

func (h *Handler) DeleteUser(c *fiber.Ctx) error {
	id := c.Params("id")

	removeUserCmd := removing.NewRemoveUserCommand(id)

	err := h.CommandBus.Dispatch(c.Context(), removeUserCmd)

	if err != nil {
		return c.Status(http.StatusNotFound).JSON(fiber.Map{"message": "user with id " + id + " does not exist"})
	}

	return c.Status(http.StatusOK).JSON(fiber.Map{"message": "user was deleted successfully"})
}
