package pipe

import (
	"fmt"
	"io"
	"math/rand"
	"net"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsPipelineError(t *testing.T) {
	t.Run("different types", func(t *testing.T) {
		for _, err := range []error{
			io.EOF,
			&os.PathError{Op: "parse", Path: "/tmp/file.txt"},
			nil,
			net.UnknownNetworkError("tdp"),
		} {
			assert.False(t, IsPipelineError(err, "parse", "/tmp/file.txt"))
		}
	})

	t.Run("equal errors", func(t *testing.T) {
		users := []string{"Bob", "John"}
		operations := []string{"bitcoin calculation", "file downloading"}
		steps := []string{"init", "hash", "download", "save"}

		for _, u := range users {
			for _, op := range operations {
				err := fmt.Errorf("wrap: %w", &PipelineError{
					User:        u,
					Name:        op,
					FailedSteps: steps[:rand.Intn(len(steps))],
				})
				assert.True(t, IsPipelineError(err, u, op))
			}
		}
	})
}
