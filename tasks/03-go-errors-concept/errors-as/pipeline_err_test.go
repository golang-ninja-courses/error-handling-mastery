package pipe

import (
	"errors"
	"fmt"
	"io"
	"net"
	"os"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestPipelineError_As(t *testing.T) {
	for i, tt := range []struct {
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
		t.Run(fmt.Sprintf("case %d", i), func(t *testing.T) {
			userErr := &UserError{}
			require.True(t, errors.As(tt.pipelineErr, &userErr))
			require.Equal(t, tt.expectedUserErr, userErr)

			// Проверяем работоспособность для nil указателя.
			var urErr *UserError
			require.True(t, errors.As(tt.pipelineErr, &urErr))
			require.Equal(t, tt.expectedUserErr, urErr)
		})
	}
}

func TestPipelineError_As_DifferentTypes(t *testing.T) {
	for i, err := range []error{
		io.EOF,
		&os.PathError{Op: "parse", Path: "/tmp/file.txt"},
		nil,
		net.UnknownNetworkError("tdp"),
	} {
		t.Run(fmt.Sprintf("case %d", i), func(t *testing.T) {
			pipeLineErr := &PipelineError{User: "parse", Name: "/tmp/file.txt"}
			require.False(t, errors.Is(pipeLineErr, err))
		})
	}
}
