package query

import "context"

type Handler interface {
	Handle(context.Context, Command) (interface{}, error)
}
