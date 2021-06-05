package main

import (
	"io"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestHandler_Handle(t *testing.T) {
	h := Handler{}

	postpone, err := h.Handle(Job{ID: 1})
	assert.NoError(t, err)
	assert.Equal(t, time.Second.Milliseconds(), postpone)

	postpone, err = h.Handle(Job{ID: 2})
	assert.NoError(t, err)
	assert.Empty(t, postpone)

	postpone, err = h.Handle(Job{ID: 3})
	assert.NoError(t, err)
	assert.Empty(t, postpone)

	postpone, err = h.Handle(Job{ID: 4})
	assert.Error(t, err)
	assert.Empty(t, postpone)
	assert.Equal(t, io.EOF, err)
}

func Test_ErrorsImplementInterfaces(t *testing.T) {
	assert.Implements(t, (*temporary)(nil), &NotReadyError{})
	assert.Implements(t, (*shouldBeSkipped)(nil), &InconsistentDataError{})
}
