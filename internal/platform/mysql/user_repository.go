package mysql

import (
	"context"
	"fmt"

	"github.com/jinzhu/gorm"
	"github.com/jorgeAM/simple-api/internal/user/domain"
)

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) domain.Repository {
	return &userRepository{db}
}

func (u *userRepository) NewUser(_ context.Context, user *domain.User) error {
	err := u.db.Create(&userSQL{
		ID:        user.ID.String(),
		Username:  user.Username.String(),
		FirstName: user.FirstName.String(),
		LastName:  user.LastName.String(),
	}).Error

	if err != nil {
		return err
	}

	return nil
}

func (u *userRepository) GetUsers(_ context.Context) ([]*domain.User, error) {
	var usersSQL []*userSQL

	err := u.db.Find(&usersSQL).Error

	if err != nil {
		return nil, err
	}

	var users []*domain.User

	for _, userSQL := range usersSQL {
		user, _ := userSQL.parseToUser()

		users = append(users, user)
	}

	return users, nil
}

func (u *userRepository) GetUser(_ context.Context, id string) (*domain.User, error) {
	var user userSQL

	err := u.db.Where("id = ?", id).First(&user).Error

	if err != nil {
		return nil, err
	}

	fmt.Println(user)

	return user.parseToUser()
}

func (u *userRepository) DeleteUser(_ context.Context, id string) error {
	err := u.db.Where("id = ?", id).Delete(userSQL{}).Error

	if err != nil {
		return err
	}

	return nil
}
