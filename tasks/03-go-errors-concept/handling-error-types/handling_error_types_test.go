package queue

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
	assert.Empty(t, postpone)

	postpone, err = h.Handle(Job{ID: 2})
	assert.NoError(t, err)
	assert.Equal(t, time.Second.Milliseconds(), postpone)

	postpone, err = h.Handle(Job{ID: 3})
	assert.NoError(t, err)
	assert.Empty(t, postpone)

	postpone, err = h.Handle(Job{ID: 4})
	assert.NoError(t, err)
	assert.Empty(t, postpone)

	postpone, err = h.Handle(Job{ID: 5})
	assert.NoError(t, err)
	assert.Empty(t, postpone)

	postpone, err = h.Handle(Job{ID: 6})
	assert.Error(t, err)
	assert.Empty(t, postpone)
	assert.Equal(t, io.EOF, err)
}
