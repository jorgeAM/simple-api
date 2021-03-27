package route

import (
	"github.com/gorilla/mux"
	"github.com/jorgeAM/api/internal/platform/server/handler"
)

// InitializeRoutes initialize all endpoints
func InitializeRoutes(handler handler.Handler) *mux.Router {
	r := mux.NewRouter()

	initializeUsersRoutes(r, handler)

	return r
}
