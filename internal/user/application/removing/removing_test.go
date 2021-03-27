package removing

import (
	"context"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/jorgeAM/api/internal/platform/repositorymock"
)

func TestRemoveUserByID(t *testing.T) {
	tests := []struct {
		name  string
		input struct {
			id string
		}
		output struct {
			err error
		}
	}{
		{
			name: "remove user without error",
			input: struct{ id string }{
				id: "47a0f027-15e6-47cc-a5d2-64183281087e",
			},
			output: struct{ err error }{err: nil},
		},
		{
			name: "remove user with error",
			input: struct{ id string }{
				id: "12",
			},
			output: struct{ err error }{err: errors.New("invalid ID")},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			mockRepository := new(repositorymock.UserMockRepository)

			mockRepository.On("DeleteUser", tt.input.id).Return(tt.output.err)

			removing := NewUserRemovingService(mockRepository)

			err := removing.RemoveUserByID(context.Background(), tt.input.id)

			assert.Equal(t, tt.output.err, err)

			mockRepository.AssertNumberOfCalls(t, "DeleteUser", 1)
			mockRepository.AssertExpectations(t)
		})
	}
}
