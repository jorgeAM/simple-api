package bus

import (
	"context"

	"github.com/jorgeAM/simple-api/kit/command"
)

type CommandBus struct {
	commands map[command.Type]command.Handler
}

func NewCommandBus() command.Bus {
	return &CommandBus{
		commands: make(map[command.Type]command.Handler),
	}
}

func (b *CommandBus) Dispatch(ctx context.Context, cmd command.Command) error {
	t := cmd.Type()
	handler, ok := b.commands[t]

	if !ok {
		return nil
	}

	return handler.Handle(ctx, cmd)
}

func (b *CommandBus) Register(t command.Type, handler command.Handler) {
	b.commands[t] = handler
}
