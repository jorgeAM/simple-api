package retrieve

import (
	"context"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/jorgeAM/api/internal/platform/repositorymock"
	"github.com/jorgeAM/api/internal/user/domain"
)

func TestGetAllUser(t *testing.T) {
	tests := []struct {
		name   string
		output struct {
			users []*domain.User
			err   error
		}
	}{
		{
			name: "get all user without error",
			output: struct {
				users []*domain.User
				err   error
			}{
				users: []*domain.User{},
				err:   nil,
			},
		},
		{
			name: "get all user with error",
			output: struct {
				users []*domain.User
				err   error
			}{
				users: nil,
				err:   errors.New("Something got wrong"),
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			mockRepository := new(repositorymock.UserMockRepository)

			mockRepository.On("GetUsers").Return(tt.output.users, tt.output.err)

			retrieving := NewUserRetrieveAllService(mockRepository)

			users, err := retrieving.GetAllUser(context.Background())

			assert.Equal(t, tt.output.err, err)
			assert.Equal(t, tt.output.users, users)

			mockRepository.AssertNumberOfCalls(t, "GetUsers", 1)
			mockRepository.AssertExpectations(t)
		})
	}
}
