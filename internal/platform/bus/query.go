package bus

import (
	"context"

	"github.com/jorgeAM/simple-api/kit/query"
)

type QueryBus struct {
	queries map[query.Type]query.Handler
}

func NewQueryBus() query.Bus {
	return &QueryBus{
		queries: make(map[query.Type]query.Handler),
	}
}

func (b *QueryBus) Dispatch(ctx context.Context, cmd query.Command) (interface{}, error) {
	t := cmd.Type()
	handler, ok := b.queries[t]

	if !ok {
		return nil, nil
	}

	return handler.Handle(ctx, cmd)
}

func (b *QueryBus) Register(t query.Type, handler query.Handler) {
	b.queries[t] = handler
}
