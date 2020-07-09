package main

import (
	"log"
	"net/http"

	_ "github.com/joho/godotenv/autoload"
	"github.com/jorgeAM/api/db"
	"github.com/jorgeAM/api/repository"
	"github.com/jorgeAM/api/routes"
	"github.com/jorgeAM/api/service"
)

func main() {
	repository := &repository.UserRepository{DB: db.GetConnection()}
	service := service.UserService{Repository: repository}

	s := http.Server{
		Addr:    ":3000",
		Handler: routes.InitializeRoutes(service),
	}

	log.Println("server is running on port 3000")
	log.Fatal(s.ListenAndServe())
}
