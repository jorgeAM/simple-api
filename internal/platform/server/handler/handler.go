package handler

import (
	"github.com/jorgeAM/simple-api/internal/user/application/finding"
	"github.com/jorgeAM/simple-api/internal/user/application/removing"
	"github.com/jorgeAM/simple-api/internal/user/application/retrieve"
	"github.com/jorgeAM/simple-api/kit/command"
)

// Handler handles all endpoint for user
type Handler struct {
	Retrieving retrieve.UserRetrieveAllService
	Finding    finding.UserRetrieveOneService
	Removing   removing.UserRemovingService
	CommandBus command.Bus
}
