package handler

import (
	"github.com/jorgeAM/simple-api/internal/user/application/creating"
	"github.com/jorgeAM/simple-api/internal/user/application/finding"
	"github.com/jorgeAM/simple-api/internal/user/application/removing"
	"github.com/jorgeAM/simple-api/internal/user/application/retrieve"
)

// Handler handles all endpoint for user
type Handler struct {
	Creating   *creating.UserCreatingService
	Retrieving *retrieve.UserRetrieveAllService
	Finding    *finding.UserRetrieveOneService
	Removing   *removing.UserRemovingService
}
