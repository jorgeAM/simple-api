package main

import (
	"log"
	"net/http"

	_ "github.com/joho/godotenv/autoload"
	"github.com/jorgeAM/api/db"
	"github.com/jorgeAM/api/repository"
	"github.com/jorgeAM/api/routes"
)

func main() {
	repository := &repository.UserRepository{DB: db.GetConnection()}

	s := http.Server{
		Addr:    ":3000",
		Handler: routes.InitializeRoutes(repository),
	}

	log.Println("server is running on port 3000")
	log.Fatal(s.ListenAndServe())
}
