package server

import (
	"net/http"
	"os"

	_ "github.com/joho/godotenv/autoload"

	"github.com/jorgeAM/api/internal/platform/server/handler"
	"github.com/jorgeAM/api/internal/platform/server/route"
)

func Run(handler handler.Handler) error {
	port := ":" + os.Getenv("PORT")

	s := http.Server{
		Addr:    port,
		Handler: route.InitializeRoutes(handler),
	}

	return s.ListenAndServe()
}
