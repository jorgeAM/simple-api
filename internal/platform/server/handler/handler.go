package handler

import (
	"github.com/jorgeAM/simple-api/internal/user/application/finding"
	"github.com/jorgeAM/simple-api/kit/command"
	"github.com/jorgeAM/simple-api/kit/query"
)

// Handler handles all endpoint for user
type Handler struct {
	Finding    finding.UserRetrieveOneService
	CommandBus command.Bus
	QueryBus   query.Bus
}
