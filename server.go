package main

import (
	"log"
	"net/http"

	_ "github.com/joho/godotenv/autoload"
	"github.com/jorgeAM/api/db"
	"github.com/jorgeAM/api/routes"
)

type user struct {
	Username string `json:"username"`
}

func main() {
	db, err := db.GetConnection()
	defer db.Close()

	if err != nil {
		log.Fatal(err)
		return
	}

	r := routes.InitializeRoutes()
	s := http.Server{
		Addr:    ":3000",
		Handler: r,
	}

	log.Println("server is running on port 3000")
	log.Fatal(s.ListenAndServe())
}
