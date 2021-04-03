package query

import "context"

type Bus interface {
	Dispatch(context.Context, Command) (interface{}, error)
	Register(Type, Handler)
}
