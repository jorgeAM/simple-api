package creating

import (
	"context"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	"github.com/jorgeAM/simple-api/internal/platform/repositorymock"
	"github.com/jorgeAM/simple-api/internal/user/domain"
)

func TestCreateNewUser(t *testing.T) {
	tests := []struct {
		name  string
		input struct {
			id        string
			username  string
			firstName string
			lastName  string
		}
		output struct {
			err error
		}
	}{
		{
			name: "Create new user without error",
			input: struct {
				id        string
				username  string
				firstName string
				lastName  string
			}{
				id:        "47a0f027-15e6-47cc-a5d2-64183281087e",
				username:  "jorgeAM",
				firstName: "Jorge",
				lastName:  "Alfaro",
			},
			output: struct{ err error }{err: nil},
		},
		{
			name: "Create new user with error",
			input: struct {
				id        string
				username  string
				firstName string
				lastName  string
			}{
				id:        "47a0f027-15e6-47cc-a5d2-64183281077e",
				username:  "jorgeAM",
				firstName: "Jorge",
				lastName:  "Alfaro",
			},
			output: struct{ err error }{err: errors.New("something got wrong")},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			mockRepository := new(repositorymock.UserMockRepository)

			user, _ := domain.NewUser(tt.input.id, tt.input.username, tt.input.firstName, tt.input.lastName)

			mockRepository.On("NewUser", mock.Anything, user).Return(tt.output.err)

			creating := NewUserCreatingService(mockRepository)

			err := creating.CreateNewUser(context.Background(), tt.input.id, tt.input.username, tt.input.firstName, tt.input.lastName)

			assert.Equal(t, tt.output.err, err)

			mockRepository.AssertNumberOfCalls(t, "NewUser", 1)
			mockRepository.AssertExpectations(t)
		})
	}
}
