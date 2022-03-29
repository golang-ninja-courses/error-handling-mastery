package queue

import (
	"fmt"
	"io"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestErrors(t *testing.T) {
	assert.EqualError(t, ErrAlreadyDone, "job is already done")
	assert.EqualError(t, ErrInconsistentData, "job payload is corrupted")
	assert.EqualError(t, ErrInvalidID, "invalid job id")
	assert.EqualError(t, ErrNotFound, "job wasn't found")
	assert.EqualError(t, ErrNotReady, "job is not ready to be performed")
}

func TestHandler_Handle(t *testing.T) {
	cases := []struct {
		job              Job
		expectedErr      error
		expectedPostpone time.Duration
	}{
		{
			job:         Job{ID: 0},
			expectedErr: nil,
		},
		{
			job:         Job{ID: 1},
			expectedErr: nil,
		},
		{
			job:              Job{ID: 2},
			expectedErr:      nil,
			expectedPostpone: time.Second,
		},
		{
			job:         Job{ID: 3},
			expectedErr: nil,
		},
		{
			job:         Job{ID: 4},
			expectedErr: nil,
		},
		{
			job:         Job{ID: 5},
			expectedErr: nil,
		},
		{
			job:         Job{ID: 6},
			expectedErr: io.EOF,
		},
	}

	for _, tt := range cases {
		t.Run(fmt.Sprintf("handle job #%d", tt.job.ID), func(t *testing.T) {
			var h Handler
			p, err := h.Handle(tt.job)
			assert.Equal(t, tt.expectedErr, err)
			assert.Equal(t, tt.expectedPostpone, p)
		})
	}
}
