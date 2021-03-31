package handler

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/jorgeAM/simple-api/internal/user/application/creating"
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

	createNewUserCmd := creating.NewCreateNewUserComand(req.ID, req.Username, req.FirstName, req.LastName)

	err = h.CommandBus.Dispatch(c.Context(), createNewUserCmd)

	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"message": err.Error()})
	}

	return c.Status(http.StatusCreated).JSON(fiber.Map{"message": "user created successfully"})
}
