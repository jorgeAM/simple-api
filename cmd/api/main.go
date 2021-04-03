package main

import (
	"log"
	"os"

	"github.com/jinzhu/gorm"

	_ "github.com/go-sql-driver/mysql"
	_ "github.com/joho/godotenv/autoload"

	"github.com/jorgeAM/simple-api/internal/platform/bus"
	"github.com/jorgeAM/simple-api/internal/platform/mysql"
	"github.com/jorgeAM/simple-api/internal/platform/server"
	"github.com/jorgeAM/simple-api/internal/platform/server/handler"
	"github.com/jorgeAM/simple-api/internal/user/application/creating"
	"github.com/jorgeAM/simple-api/internal/user/application/finding"
	"github.com/jorgeAM/simple-api/internal/user/application/removing"
	"github.com/jorgeAM/simple-api/internal/user/application/retrieve"
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

	commandBus := bus.NewCommandBus()
	queryBus := bus.NewQueryBus()

	creatingService := creating.NewUserCreatingService(repository)
	retrievingService := retrieve.NewUserRetrieveAllService(repository)
	findingService := finding.NewUserRetrieveOneService(repository)
	removingService := removing.NewUserRemovingService(repository)

	createNewUserHandler := creating.NewCreateNewUserHandler(creatingService)
	removeUserHandler := removing.NewRemoveUserHandler(removingService)

	// register commands
	commandBus.Register(creating.CreateNewUserCommandType, createNewUserHandler)
	commandBus.Register(removing.RemoveUserCommandType, removeUserHandler)

	getAllUserHandler := retrieve.NewGetAllUsersHandler(retrievingService)
	findUserByIDHandler := finding.NewFindUserByIDHandler(findingService)

	// register queries
	queryBus.Register(retrieve.GetAllUsersQueryType, getAllUserHandler)
	queryBus.Register(finding.FindUserByIDQueryType, findUserByIDHandler)

	handler := handler.Handler{
		CommandBus: commandBus,
		QueryBus:   queryBus,
	}

	log.Println("server is running ...")

	srv := server.NewServer(handler)

	if err := srv.Run(); err != nil {
		log.Fatalf("something got wrong when we try to run web server %v", err)
	}
}
