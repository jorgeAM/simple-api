package mysql

import "github.com/jorgeAM/simple-api/internal/user/domain"

type userSQL struct {
	ID        string
	Username  string
	FirstName string
	LastName  string
}

func (userSQL) TableName() string {
	return "users"
}

func (u userSQL) parseToUser() (*domain.User, error) {
	user, err := domain.NewUser(u.ID, u.Username, u.FirstName, u.LastName)

	if err != nil {
		return nil, err
	}

	return user, nil
}
