package errors

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestErrors(t *testing.T) {
	assert.EqualError(t, ErrAlreadyDone, "job is already done")
	assert.EqualError(t, ErrInconsistentData, "job payload is corrupted")
	assert.EqualError(t, ErrInvalidID, "invalid job id")
	assert.EqualError(t, ErrNotFound, "job wasn't found")
	assert.EqualError(t, ErrNotReady, "job is not ready to be performed")
}
