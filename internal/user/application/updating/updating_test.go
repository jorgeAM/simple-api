package updating

import (
	"context"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/jorgeAM/api/internal/platform/repositorymock"
	"github.com/jorgeAM/api/internal/user/domain"
)

func TestUpdateUser(t *testing.T) {
	tests := []struct {
		name  string
		input struct {
			user *domain.User
		}
		output struct {
			user *domain.User
			err  error
		}
	}{
		{
			name: "Update user without error",
			input: struct{ user *domain.User }{
				user: &domain.User{
					ID:        "47a0f027-15e6-47cc-a5d2-64183281087e",
					Username:  "jorgeAM",
					FirstName: "Jorge",
					LastName:  "Alfaro",
				},
			},
			output: struct {
				user *domain.User
				err  error
			}{
				user: &domain.User{
					ID:        "47a0f027-15e6-47cc-a5d2-64183281087e",
					Username:  "jorgeAM",
					FirstName: "Jorge",
					LastName:  "Alfaro",
				},
				err: nil,
			},
		},
		{
			name: "Update user with error",
			input: struct{ user *domain.User }{
				user: &domain.User{
					ID:        "12",
					Username:  "jorgeAM",
					FirstName: "Jorge",
					LastName:  "Alfaro",
				},
			},
			output: struct {
				user *domain.User
				err  error
			}{
				user: nil,
				err:  errors.New("Invalid ID"),
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			mockRepository := new(repositorymock.UserMockRepository)

			mockRepository.On("UpdateUser", tt.input.user).Return(tt.output.user, tt.output.err)

			updating := NewUserUpdatingService(mockRepository)

			updatedUser, err := updating.UpdateUser(context.Background(), tt.input.user)

			assert.Equal(t, tt.output.err, err)
			assert.Equal(t, tt.output.user, updatedUser)

			mockRepository.AssertNumberOfCalls(t, "UpdateUser", 1)
			mockRepository.AssertExpectations(t)
		})
	}
}
