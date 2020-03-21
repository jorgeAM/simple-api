package main

import (
	"log"
	"net/http"

	_ "github.com/joho/godotenv/autoload"
	"github.com/jorgeAM/api/routes"
)

func main() {
	s := http.Server{
		Addr:    ":3000",
		Handler: routes.InitializeRoutes(),
	}

	log.Println("server is running on port 3000")
	log.Fatal(s.ListenAndServe())
}
