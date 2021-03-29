package finding

import (
	"context"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"

	"github.com/jorgeAM/simple-api/internal/platform/repositorymock"
	"github.com/jorgeAM/simple-api/internal/user/domain"
)

func TestFindUserByID(t *testing.T) {
	user, _ := domain.NewUser("47a0f027-15e6-47cc-a5d2-64183281087e", "jorgeAM", "jorge", "alfaro")

	var tests []struct {
		name   string
		input  domain.UserID
		output struct {
			user *domain.User
			err  error
		}
	}

	userID, err := domain.NewUserID("47a0f027-15e6-47cc-a5d2-64183281087e")
	require.NoError(t, err)

	tests = append(tests, struct {
		name   string
		input  domain.UserID
		output struct {
			user *domain.User
			err  error
		}
	}{
		name:  "get user by ID without error",
		input: userID,
		output: struct {
			user *domain.User
			err  error
		}{
			user: user,
			err:  nil,
		},
	}, struct {
		name   string
		input  domain.UserID
		output struct {
			user *domain.User
			err  error
		}
	}{name: "get user by ID without error",
		input: userID,
		output: struct {
			user *domain.User
			err  error
		}{
			user: nil,
			err:  errors.New("Something got wrong"),
		},
	})

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			mockRepository := new(repositorymock.UserMockRepository)

			mockRepository.On("GetUser", mock.Anything, tt.input.String()).Return(tt.output.user, tt.output.err)

			retrieving := NewUserRetrieveOneService(mockRepository)

			_, err := retrieving.FindUserByID(context.Background(), tt.input)

			assert.Equal(t, tt.output.err, err)

			mockRepository.AssertNumberOfCalls(t, "GetUser", 1)
			mockRepository.AssertExpectations(t)
		})
	}
}
