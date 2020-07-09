package service

import (
	"github.com/jorgeAM/api/models"
	"github.com/jorgeAM/api/repository"
)

// UserService use userRepository
type UserService struct {
	Repository repository.Repository
}

// NewUser creates new user
func (u *UserService) NewUser(user *models.User) (*models.User, error) {
	return u.Repository.NewUser(user)
}

// GetUsers retrieves all users
func (u *UserService) GetUsers() ([]*models.User, error) {
	return u.Repository.GetUsers()
}

// GetUser retrieves one user
func (u *UserService) GetUser(id int) (*models.User, error) {
	return u.Repository.GetUser(id)
}

// UpdateUser updates a user
func (u *UserService) UpdateUser(user *models.User) (*models.User, error) {
	_, err := u.Repository.GetUser(user.ID)

	if err != nil {
		return nil, err
	}

	return u.Repository.UpdateUser(user)
}

// DeleteUser removes a user
func (u *UserService) DeleteUser(id int) error {
	_, err := u.Repository.GetUser(id)

	if err != nil {
		return err
	}

	return u.Repository.DeleteUser(id)
}
