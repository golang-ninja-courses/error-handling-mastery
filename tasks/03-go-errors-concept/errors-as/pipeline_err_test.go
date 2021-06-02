package pipe

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestPipelineErr_As(t *testing.T) {
	for _, tt := range []struct {
		pipelineErr     *PipelineErr
		expectedUserErr *UserError
	}{
		{
			pipelineErr:     &PipelineErr{User: "Bob", Name: "bitcoin calculation", FailedSteps: []string{"step 1", "step 2"}},
			expectedUserErr: &UserError{User: "Bob", Operation: "bitcoin calculation"},
		},
		{
			pipelineErr:     &PipelineErr{User: "Alex", Name: "file downloading"},
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
