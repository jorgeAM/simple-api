package handler

import (
	"github.com/jorgeAM/simple-api/kit/command"
	"github.com/jorgeAM/simple-api/kit/query"
)

// Handler handles all endpoint for user
type Handler struct {
	CommandBus command.Bus
	QueryBus   query.Bus
}
