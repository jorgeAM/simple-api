package repository

import (
	"github.com/jinzhu/gorm"
	"github.com/jorgeAM/api/models"
)

// UserRepository implements Repository interface
type UserRepository struct {
	DB *gorm.DB
}

// NewUser record new user in database
func (u *UserRepository) NewUser(user *models.User) (*models.User, error) {
	err := u.DB.Table("Users").Create(user).Error

	if err != nil {
		return nil, err
	}

	return user, nil
}

// GetUsers retrieves all users
func (u *UserRepository) GetUsers() ([]*models.User, error) {
	var users []*models.User
	err := u.DB.Table("Users").Find(&users).Error

	if err != nil {
		return nil, err
	}

	return users, nil
}

// GetUser retrieve one user
func (u *UserRepository) GetUser(id int) (*models.User, error) {
	user := new(models.User)
	err := u.DB.Table("Users").First(user, id).Error

	if err != nil {
		return nil, err
	}

	return user, nil
}

// UpdateUser updates user
func (u *UserRepository) UpdateUser(user *models.User) (*models.User, error) {
	err := u.DB.Table("Users").Save(user).Error

	if err != nil {
		return nil, err
	}

	return user, nil
}

// DeleteUser deletes a user
func (u *UserRepository) DeleteUser(id int) error {
	err := u.DB.Table("Users").Where("id = ?", id).Delete(models.User{}).Error

	if err != nil {
		return err
	}

	return nil
}
