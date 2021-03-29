package mysql

import (
	"context"

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
	err := u.db.Create(&userPrimitive{
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
	var usersSQL []*userPrimitive

	err := u.db.Find(&usersSQL).Error

	if err != nil {
		return nil, err
	}

	var users []*domain.User

	for _, userSQL := range usersSQL {
		user, _ := userSQL.UnmarshalAggregate()

		users = append(users, user)
	}

	return users, nil
}

func (u *userRepository) GetUser(_ context.Context, id string) (*domain.User, error) {
	var user userPrimitive

	err := u.db.Where("id = ?", id).First(&user).Error

	if err != nil {
		return nil, err
	}

	return user.UnmarshalAggregate()
}

func (u *userRepository) DeleteUser(_ context.Context, id string) error {
	err := u.db.Where("id = ?", id).Delete(userPrimitive{}).Error

	if err != nil {
		return err
	}

	return nil
}
