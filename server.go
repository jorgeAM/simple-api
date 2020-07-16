package main

import (
	"log"
	"net/http"
	"os"

	_ "github.com/joho/godotenv/autoload"
	"github.com/jorgeAM/api/db"
	"github.com/jorgeAM/api/repository"
	"github.com/jorgeAM/api/routes"
	"github.com/jorgeAM/api/service"
)

func main() {
	port := ":" + os.Getenv("PORT")
	repository := &repository.UserRepository{DB: db.GetConnection()}
	service := service.UserService{Repository: repository}

	s := http.Server{
		Addr:    port,
		Handler: routes.InitializeRoutes(service),
	}

	log.Println("server is running on port " + port)
	log.Fatal(s.ListenAndServe())
}
