package mysql

import (
	"github.com/jinzhu/gorm"
	"github.com/jorgeAM/simple-api/internal/user/domain"
)

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) domain.Repository {
	return &userRepository{db}
}

func (u *userRepository) NewUser(user *domain.User) error {
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

func (u *userRepository) GetUsers() ([]*domain.User, error) {
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

func (u *userRepository) GetUser(id string) (*domain.User, error) {
	var user userSQL

	err := u.db.Where("id = ?", id).First(user).Error

	if err != nil {
		return nil, err
	}

	return user.parseToUser()
}

func (u *userRepository) UpdateUser(user *domain.User) (*domain.User, error) {
	userDB := &userSQL{
		ID:        user.ID.String(),
		Username:  user.Username.String(),
		FirstName: user.FirstName.String(),
		LastName:  user.LastName.String(),
	}

	err := u.db.Save(userDB).Error

	if err != nil {
		return nil, err
	}

	return userDB.parseToUser()
}

func (u *userRepository) DeleteUser(id string) error {
	err := u.db.Where("id = ?", id).Delete(userSQL{}).Error

	if err != nil {
		return err
	}

	return nil
}
