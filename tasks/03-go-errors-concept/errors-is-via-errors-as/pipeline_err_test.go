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

					isEqual := (u1 == u2) && (op1 == op2)
					assert.Equal(t, isEqual, IsPipelineError(err, u2, op2),
						"err: %#v, user: %v, pipeline name: %v", err, u2, op2)
				}
			}
		}
	}
}

func TestIsPipelineError_DifferentTypes(t *testing.T) {
	for i, err := range []error{
		io.EOF,
		&os.PathError{Op: "parse", Path: "/tmp/file.txt"},
		nil,
		net.UnknownNetworkError("tdp"),
	} {
		t.Run(fmt.Sprintf("case %d", i), func(t *testing.T) {
			assert.False(t, IsPipelineError(err, "parse", "/tmp/file.txt"))
		})
	}
}
