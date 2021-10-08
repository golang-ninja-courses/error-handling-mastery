package queue

import (
	"fmt"
	"io"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestHandler_Handle(t *testing.T) {
	cases := []struct {
		job              Job
		expectedErr      error
		expectedPostpone time.Duration
	}{
		{
			job:         Job{ID: 1},
			expectedErr: nil,
		},
		{
			job:              Job{ID: 2},
			expectedErr:      nil,
			expectedPostpone: defaultPostpone,
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

func TestErrorHelpers(t *testing.T) {
	cases := []struct {
		err             error
		isTemporary     bool
		shouldBeSkipped bool
	}{
		{
			err:             new(AlreadyDoneError),
			isTemporary:     false,
			shouldBeSkipped: true,
		},
		{
			err:             new(InconsistentDataError),
			isTemporary:     false,
			shouldBeSkipped: true,
		},
		{
			err:             new(InvalidIDError),
			isTemporary:     false,
			shouldBeSkipped: true,
		},
		{
			err:             new(NotFoundError),
			isTemporary:     false,
			shouldBeSkipped: true,
		},
		{
			err:             new(NotReadyError),
			isTemporary:     true,
			shouldBeSkipped: false,
		},
	}

	for _, tt := range cases {
		t.Run(fmt.Sprintf("%T", tt.err), func(t *testing.T) {
			assert.Equal(t, tt.isTemporary, isTemporary(tt.err))
			assert.Equal(t, tt.shouldBeSkipped, shouldBeSkipped(tt.err))
		})
	}
}
