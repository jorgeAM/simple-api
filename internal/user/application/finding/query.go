package finding

import (
	"github.com/jorgeAM/simple-api/kit/query"
)

const FindUserByIDQueryType query.Type = "query.find.user"

type FindUserByIDQuery struct {
	id string
}

func (c FindUserByIDQuery) Type() query.Type {
	return FindUserByIDQueryType
}

func NewFindUserByIDQuery(id string) FindUserByIDQuery {
	return FindUserByIDQuery{
		id: id,
	}
}
