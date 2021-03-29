package route

import (
	"github.com/gofiber/fiber/v2"
	"github.com/jorgeAM/simple-api/internal/platform/server/handler"
)

func InitializeUsersRoutes(r fiber.Router, handler handler.Handler) {
	r.Post("", handler.NewUser)
	r.Get("", handler.GetUsers)
	r.Get("/:id", handler.GetUser)
	r.Delete("/:id", handler.DeleteUser)
}
