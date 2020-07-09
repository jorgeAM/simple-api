package service

import "github.com/jorgeAM/api/models"

// Service interface had all methods
type Service interface {
	NewUser(user *models.User) (*models.User, error)
	GetUsers() ([]*models.User, error)
	GetUser(id int) (*models.User, error)
	UpdateUser(user *models.User) (*models.User, error)
	DeleteUser(id int) error
}
