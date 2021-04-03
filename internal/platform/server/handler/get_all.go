package handler

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/jorgeAM/simple-api/internal/user/application/retrieve"
)

func (h *Handler) GetUsers(c *fiber.Ctx) error {
	getAllUserQuery := retrieve.NewGetAllUsersQuery()

	users, err := h.QueryBus.Dispatch(c.Context(), getAllUserQuery)

	if err != nil {
		return c.Status(http.StatusNotFound).JSON(fiber.Map{"message": err.Error()})
	}

	return c.Status(http.StatusOK).JSON(fiber.Map{"users": users})
}
