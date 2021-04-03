package handler

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/jorgeAM/simple-api/internal/user/application/finding"
)

func (h *Handler) GetUser(c *fiber.Ctx) error {
	id := c.Params("id")

	findUserByIDQuery := finding.NewFindUserByIDQuery(id)

	user, err := h.QueryBus.Dispatch(c.Context(), findUserByIDQuery)

	if err != nil {
		return c.Status(http.StatusNotFound).JSON(fiber.Map{"message": err.Error()})
	}

	return c.Status(http.StatusOK).JSON(fiber.Map{"user": user})
}
