package handler

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
)

type createUserRequest struct {
	ID        string `json:"id"`
	Username  string `json:"username"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
}

func (h *Handler) NewUser(c *fiber.Ctx) error {
	var req createUserRequest

	err := c.BodyParser(&req)

	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"message": "something got wrong to parse request body"})
	}

	err = h.Creating.CreateNewUser(c.Context(), req.ID, req.Username, req.FirstName, req.LastName)

	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"message": err.Error()})
	}

	return c.Status(http.StatusCreated).JSON(fiber.Map{"message": "user created successfully"})
}
