package retrieve

import "github.com/jorgeAM/simple-api/internal/user/domain"

type userResponse struct {
	ID        string `json:"id"`
	Username  string `json:"username"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
}

func NewUserResponseFromAggregate(user *domain.User) *userResponse {
	return &userResponse{
		ID:        user.ID.String(),
		Username:  user.Username.String(),
		FirstName: user.FirstName.String(),
		LastName:  user.LastName.String(),
	}
}
