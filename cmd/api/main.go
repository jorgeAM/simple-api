package main

import (
	"log"
	"os"

	"github.com/jinzhu/gorm"

	_ "github.com/go-sql-driver/mysql"
	_ "github.com/joho/godotenv/autoload"

	"github.com/jorgeAM/api/internal/platform/mysql"
	"github.com/jorgeAM/api/internal/platform/server"
	"github.com/jorgeAM/api/internal/platform/server/handler"
	"github.com/jorgeAM/api/internal/user/application/creating"
	"github.com/jorgeAM/api/internal/user/application/removing"
	"github.com/jorgeAM/api/internal/user/application/retrieve"
	"github.com/jorgeAM/api/internal/user/application/updating"
)

var (
	engine = os.Getenv("DB_ENGINE")
	dsn    = os.Getenv("JAWSDB_URL")
)

func main() {
	db, err := gorm.Open(engine, dsn)

	if err != nil {
		log.Fatalf("something got wrong to connect with database %v", err)
	}

	repository := mysql.NewUserRepository(db)

	creating := creating.NewUserCreatingService(repository)
	retrieving := retrieve.NewUserRetrieveAllService(repository)
	finding := retrieve.NewUserRetrieveOneService(repository)
	updating := updating.NewUserUpdatingService(repository)
	removing := removing.NewUserRemovingService(repository)

	handler := handler.Handler{
		Creating:   creating,
		Retrieving: retrieving,
		Finding:    finding,
		Updating:   updating,
		Removing:   removing,
	}

	log.Println("server is running ...")

	if err := server.Run(handler); err != nil {
		log.Fatalf("something got wrong when we try to run web server %v", err)
	}
}
