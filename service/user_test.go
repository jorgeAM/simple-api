package service

import (
	"errors"
	"testing"

	"github.com/jorgeAM/api/models"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type UserServiceMock struct {
	mock.Mock
}

func (u *UserServiceMock) NewUser(user *models.User) (*models.User, error) {
	args := u.Called(user)
	return args.Get(0).(*models.User), args.Error(1)
}

func (u *UserServiceMock) GetUsers() ([]*models.User, error) {
	args := u.Called()
	return args.Get(0).([]*models.User), args.Error(1)
}

func (u *UserServiceMock) GetUser(id int) (*models.User, error) {
	args := u.Called(id)
	return args.Get(0).(*models.User), args.Error(1)
}

func (u *UserServiceMock) UpdateUser(user *models.User) (*models.User, error) {
	args := u.Called(user)
	return args.Get(0).(*models.User), args.Error(1)
}

func (u *UserServiceMock) DeleteUser(id int) error {
	args := u.Called(id)
	return args.Error(0)
}

func TestNewUser(t *testing.T) {
	user := new(models.User)

	tests := []struct {
		input  *models.User
		output struct {
			user *models.User
			err  error
		}
	}{
		{
			user,
			struct {
				user *models.User
				err  error
			}{
				user,
				nil,
			},
		},
		{
			nil,
			struct {
				user *models.User
				err  error
			}{
				nil,
				errors.New("Something got wrong"),
			},
		},
	}

	for _, tt := range tests {
		t.Run("TestNewUser", func(t *testing.T) {
			t.Parallel()

			testObj := new(UserServiceMock)

			testObj.On("NewUser", tt.input).Return(tt.output.user, tt.output.err)

			u, err := testObj.NewUser(tt.input)

			assert.Equal(t, tt.output.user, u)
			assert.Equal(t, tt.output.err, err)

			testObj.AssertNumberOfCalls(t, "NewUser", 1)
			testObj.AssertExpectations(t)
		})
	}
}

func TestGetUsers(t *testing.T) {
	tests := []struct {
		output struct {
			users []*models.User
			err   error
		}
	}{
		{
			output: struct {
				users []*models.User
				err   error
			}{
				users: []*models.User{},
				err:   nil,
			},
		},
		{
			output: struct {
				users []*models.User
				err   error
			}{
				users: nil,
				err:   errors.New("Something got wrong"),
			},
		},
	}

	for _, tt := range tests {
		t.Run("TestGetUsers", func(t *testing.T) {
			t.Parallel()

			testObj := new(UserServiceMock)

			testObj.On("GetUsers").Return(tt.output.users, tt.output.err)

			users, err := testObj.GetUsers()

			assert.Equal(t, tt.output.users, users)
			assert.Equal(t, tt.output.err, err)

			testObj.AssertNumberOfCalls(t, "GetUsers", 1)
			testObj.AssertExpectations(t)
		})
	}
}

func TestGetUser(t *testing.T) {
	tests := []struct {
		input  int
		output struct {
			user *models.User
			err  error
		}
	}{
		{
			input: 1,
			output: struct {
				user *models.User
				err  error
			}{
				user: &models.User{
					ID:        1,
					FirstName: "Jorge",
					LastName:  "Alfaro",
					Username:  "JorgeAM",
				},
				err: nil,
			},
		},
		{
			input: 2,
			output: struct {
				user *models.User
				err  error
			}{
				user: nil,
				err:  errors.New("There is no user with id: 2"),
			},
		},
	}

	for _, tt := range tests {
		t.Run("TestGetUser", func(t *testing.T) {
			t.Parallel()

			testObj := new(UserServiceMock)

			testObj.On("GetUser", tt.input).Return(tt.output.user, tt.output.err)

			user, err := testObj.GetUser(tt.input)

			assert.Equal(t, tt.output.user, user)
			assert.Equal(t, tt.output.err, err)

			testObj.AssertNumberOfCalls(t, "GetUser", 1)
			testObj.AssertExpectations(t)
		})
	}
}

func TestUpdateUser(t *testing.T) {
	user := new(models.User)

	tests := []struct {
		input  *models.User
		output struct {
			user *models.User
			err  error
		}
	}{
		{
			user,
			struct {
				user *models.User
				err  error
			}{
				user,
				nil,
			},
		},
		{
			nil,
			struct {
				user *models.User
				err  error
			}{
				nil,
				errors.New("Something got wrong"),
			},
		},
	}

	for _, tt := range tests {
		t.Run("TestUpdateUser", func(t *testing.T) {
			t.Parallel()

			testObj := new(UserServiceMock)

			testObj.On("UpdateUser", tt.input).Return(tt.output.user, tt.output.err)

			u, err := testObj.UpdateUser(tt.input)

			assert.Equal(t, tt.output.user, u)
			assert.Equal(t, tt.output.err, err)

			testObj.AssertNumberOfCalls(t, "UpdateUser", 1)
			testObj.AssertExpectations(t)
		})
	}
}

func TestDeleteUser(t *testing.T) {
	tests := []struct {
		input  int
		output error
	}{
		{
			input:  1,
			output: nil,
		},
		{
			input:  2,
			output: errors.New("There is no user with id: 2"),
		},
	}

	for _, tt := range tests {
		t.Run("TestDeleteUser", func(t *testing.T) {
			t.Parallel()

			testObj := new(UserServiceMock)

			testObj.On("DeleteUser", tt.input).Return(tt.output)

			err := testObj.DeleteUser(tt.input)

			assert.Equal(t, tt.output, err)

			testObj.AssertNumberOfCalls(t, "DeleteUser", 1)
			testObj.AssertExpectations(t)
		})
	}
}
