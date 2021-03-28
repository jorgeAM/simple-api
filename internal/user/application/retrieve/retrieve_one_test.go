package retrieve

import (
	"context"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/jorgeAM/simple-api/internal/platform/repositorymock"
	"github.com/jorgeAM/simple-api/internal/user/domain"
)

func TestFindUserByID(t *testing.T) {
	user, _ := domain.NewUser("47a0f027-15e6-47cc-a5d2-64183281087e", "jorgeAM", "jorge", "alfaro")

	tests := []struct {
		name  string
		input struct {
			id string
		}
		output struct {
			user *domain.User
			err  error
		}
	}{
		{
			name: "get user by ID without error",
			input: struct{ id string }{
				id: "47a0f027-15e6-47cc-a5d2-64183281087e",
			},
			output: struct {
				user *domain.User
				err  error
			}{
				user: user,
				err:  nil,
			},
		},
		{
			name: "get user by ID with error",
			input: struct{ id string }{
				id: "28",
			},
			output: struct {
				user *domain.User
				err  error
			}{
				user: nil,
				err:  errors.New("Something got wrong"),
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			mockRepository := new(repositorymock.UserMockRepository)

			mockRepository.On("GetUser", tt.input.id).Return(tt.output.user, tt.output.err)

			retrieving := NewUserRetrieveOneService(mockRepository)

			user, err := retrieving.FindUserByID(context.Background(), tt.input.id)

			assert.Equal(t, tt.output.err, err)
			assert.Equal(t, tt.output.user, user)

			mockRepository.AssertNumberOfCalls(t, "GetUser", 1)
			mockRepository.AssertExpectations(t)
		})
	}
}
