package pipe

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestPipelineError_As(t *testing.T) {
	for _, tt := range []struct {
		pipelineErr     *PipelineError
		expectedUserErr *UserError
	}{
		{
			pipelineErr:     &PipelineError{User: "Bob", Name: "bitcoin calculation", FailedSteps: []string{"step 1", "step 2"}},
			expectedUserErr: &UserError{User: "Bob", Operation: "bitcoin calculation"},
		},
		{
			pipelineErr:     &PipelineError{User: "Alex", Name: "file downloading"},
			expectedUserErr: &UserError{User: "Alex", Operation: "file downloading"},
		},
	} {
		userErr := &UserError{}
		require.True(t, errors.As(tt.pipelineErr, &userErr))
		require.Equal(t, tt.expectedUserErr, userErr)

		// Проверяем работоспособность для nil указателя.
		var urErr *UserError
		require.True(t, errors.As(tt.pipelineErr, &urErr))
		require.Equal(t, tt.expectedUserErr, urErr)
	}
}
