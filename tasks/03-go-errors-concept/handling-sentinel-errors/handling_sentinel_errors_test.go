package main

import (
	"io"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func Test_Errors(t *testing.T) {
	assert.EqualError(t, ErrInconsistentData, "job payload is corrupted")
	assert.EqualError(t, ErrNotReady, "job is not ready to be performed")
	assert.EqualError(t, ErrNotFound, "job wasn't found")
	assert.EqualError(t, ErrAlreadyDone, "job is already done")
	assert.EqualError(t, ErrInvalidId, "invalid job id")
}

func TestHandler_Handle(t *testing.T) {
	h := Handler{}

	postpone, err := h.Handle(Job{ID: 1})
	assert.Error(t, err)
	assert.Empty(t, postpone)
	assert.Equal(t, ErrInconsistentData, err)

	postpone, err = h.Handle(Job{ID: 2})
	assert.NoError(t, err)
	assert.Equal(t, time.Second.Milliseconds(), postpone)

	postpone, err = h.Handle(Job{ID: 3})
	assert.Error(t, err)
	assert.Empty(t, postpone)
	assert.Equal(t, ErrNotFound, err)

	postpone, err = h.Handle(Job{ID: 4})
	assert.NoError(t, err)
	assert.Empty(t, postpone)

	postpone, err = h.Handle(Job{ID: 5})
	assert.Error(t, err)
	assert.Empty(t, postpone)
	assert.Equal(t, ErrInvalidId, err)

	postpone, err = h.Handle(Job{ID: 6})
	assert.Error(t, err)
	assert.Empty(t, postpone)
	assert.Equal(t, io.EOF, err)
}
