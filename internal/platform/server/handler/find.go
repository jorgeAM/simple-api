package handler

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/jorgeAM/simple-api/internal/user/domain"
)

func (h *Handler) GetUser(c *fiber.Ctx) error {
	id := c.Params("id")

	userID, err := domain.NewUserID(id)

	if err != nil {
		return c.Status(http.StatusNotFound).JSON(fiber.Map{"message": err.Error()})
	}

	user, err := h.Finding.FindUserByID(c.Context(), userID)

	if err != nil {
		return c.Status(http.StatusNotFound).JSON(fiber.Map{"message": "user with id " + id + " does not exist"})
	}

	return c.Status(http.StatusOK).JSON(fiber.Map{"user": user})
}
