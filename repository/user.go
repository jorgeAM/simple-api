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
	err := u.DB.Create(user).Error

	if err != nil {
		return nil, err
	}

	return user, nil
}

// GetUsers retrieves all users
func (u *UserRepository) GetUsers() ([]*models.User, error) {
	var users []*models.User
	err := u.DB.Find(&users).Error

	if err != nil {
		return nil, err
	}

	return users, nil
}

// GetUser retrieve one user
func (u *UserRepository) GetUser(id int) (*models.User, error) {
	user := new(models.User)
	err := u.DB.First(user, id).Error

	if err != nil {
		return nil, err
	}

	return user, nil
}

// UpdateUser updates user
func (u *UserRepository) UpdateUser(user *models.User) (*models.User, error) {
	err := u.DB.Save(user).Error

	if err != nil {
		return nil, err
	}

	return user, nil
}

// DeleteUser deletes a user
func (u *UserRepository) DeleteUser(id int) error {
	err := u.DB.Where("id = ?", id).Delete(models.User{}).Error

	if err != nil {
		return err
	}

	return nil
}
