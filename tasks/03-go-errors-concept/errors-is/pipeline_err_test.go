package pipe

import (
	"errors"
	"fmt"
	"io"
	"math/rand"
	"net"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestPipelineError_Is(t *testing.T) {
	users := []string{"Bob", "John"}
	operations := []string{"bitcoin calculation", "file downloading"}
	steps := []string{"init", "hash", "download", "save"}

	for _, u1 := range users {
		for _, u2 := range users {
			for _, op1 := range operations {
				for _, op2 := range operations {
					err := fmt.Errorf("wrap: %w", &PipelineError{
						User:        u1,
						Name:        op1,
						FailedSteps: steps[:rand.Intn(len(steps))],
					})

					target := &PipelineError{
						User:        u2,
						Name:        op2,
						FailedSteps: steps[:rand.Intn(len(steps))],
					}

					isEqual := (u1 == u2) && (op1 == op2)

					var msg string
					if isEqual {
						msg = "errors %#v and %#v must be equal"
					} else {
						msg = "errors %#v and %#v must be not equal"
					}
					assert.Equal(t, isEqual, errors.Is(err, target), msg, err, target)
				}
			}
		}
	}
}

func TestPipelineError_Is_DifferentTypes(t *testing.T) {
	for i, err := range []error{
		io.EOF,
		&os.PathError{Op: "parse", Path: "/tmp/file.txt"},
		nil,
		net.UnknownNetworkError("tdp"),
		errors.New("integration error"),
	} {
		t.Run(fmt.Sprintf("case %d", i), func(t *testing.T) {
			pipeLineErr := &PipelineError{User: "parse", Name: "/tmp/file.txt"}
			require.False(t, errors.Is(pipeLineErr, err))
			require.False(t, errors.Is(err, pipeLineErr))
		})
	}
}
