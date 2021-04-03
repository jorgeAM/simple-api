package retrieve

import (
	"github.com/jorgeAM/simple-api/kit/query"
)

const GetAllUsersQueryType query.Type = "query.get.all.user"

type GetAllUsersQuery struct{}

func (c GetAllUsersQuery) Type() query.Type {
	return GetAllUsersQueryType
}

func NewGetAllUsersQuery() GetAllUsersQuery {
	return GetAllUsersQuery{}
}
