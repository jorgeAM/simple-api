package repository

import "github.com/jorgeAM/api/models"

// Repository use case to user entity
type Repository interface {
	NewUser(user *models.User) (*models.User, error)
	GetUsers() ([]*models.User, error)
	GetUser(id int) (*models.User, error)
	UpdateUser(user *models.User) (*models.User, error)
	DeleteUser(id int) error
}
