package server

import (
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"

	_ "github.com/joho/godotenv/autoload"

	"github.com/jorgeAM/simple-api/internal/platform/server/handler"
	"github.com/jorgeAM/simple-api/internal/platform/server/route"
)

type Server struct {
	engine *fiber.App
	port   string
}

func NewServer(handler handler.Handler) *Server {
	server := &Server{
		engine: fiber.New(),
		port:   ":" + os.Getenv("PORT"),
	}

	server.engine.Use(logger.New())

	server.registerRoutes(handler)

	return server
}

func (s *Server) registerRoutes(handler handler.Handler) {
	router := s.engine.Group("/users")

	route.InitializeUsersRoutes(router, handler)
}

func (s *Server) Run() error {
	return s.engine.Listen(s.port)
}
