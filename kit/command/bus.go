package command

import "context"

type Bus interface {
	Dispatch(context.Context, Command) error
	Register(Type, Handler)
}
