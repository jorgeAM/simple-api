package mysql

import (
	"github.com/jinzhu/gorm"
	"github.com/jorgeAM/api/internal/user/domain"
)

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) domain.Repository {
	return &userRepository{db}
}

func (u *userRepository) NewUser(user *domain.User) (*domain.User, error) {
	err := u.db.Create(user).Error

	if err != nil {
		return nil, err
	}

	return user, nil
}

func (u *userRepository) GetUsers() ([]*domain.User, error) {
	var users []*domain.User
	err := u.db.Find(&users).Error

	if err != nil {
		return nil, err
	}

	return users, nil
}

func (u *userRepository) GetUser(id string) (*domain.User, error) {
	user := new(domain.User)
	err := u.db.First(user, id).Error

	if err != nil {
		return nil, err
	}

	return user, nil
}

func (u *userRepository) UpdateUser(user *domain.User) (*domain.User, error) {
	err := u.db.Save(user).Error

	if err != nil {
		return nil, err
	}

	return user, nil
}

func (u *userRepository) DeleteUser(id string) error {
	err := u.db.Where("id = ?", id).Delete(domain.User{}).Error

	if err != nil {
		return err
	}

	return nil
}
