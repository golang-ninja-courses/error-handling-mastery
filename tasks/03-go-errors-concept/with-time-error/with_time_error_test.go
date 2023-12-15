package errs

import (
	"context"
	"fmt"
	"io"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func ExampleWithTimeError() {
	err := NewWithTimeError(context.Canceled)
	fmt.Println(err)
	// context canceled, occurred at: 2021-06-07T08:15:48.518835+03:00
}

func TestWithTimeError(t *testing.T) {
	var s int64
	timeNowMock := func() time.Time {
		s++
		return time.Unix(s, 0)
	}

	newWithTimeErrorMock := func(err error) error {
		return newWithTimeError(err, timeNowMock)
	}

	var timed interface {
		Time() time.Time
	}

	err := newWithTimeErrorMock(context.Canceled)
	require.ErrorIs(t, err, context.Canceled)
	require.ErrorAs(t, err, &timed)
	require.Equal(t, 1, timed.Time().Second())

	err = fmt.Errorf("cannot read file: %w", newWithTimeErrorMock(io.ErrUnexpectedEOF))
	require.ErrorIs(t, err, io.ErrUnexpectedEOF)
	require.ErrorAs(t, err, &timed)
	require.Equal(t, 2, timed.Time().Second())
}

func TestWithTimeError_FrozenTime(t *testing.T) {
	var timed interface {
		Time() time.Time
	}

	err := NewWithTimeError(context.Canceled)
	require.ErrorAs(t, err, &timed)

	t1 := timed.Time()
	time.Sleep(time.Millisecond)
	t2 := timed.Time()
	require.True(t, t1.Equal(t2),
		"Time() must return the time the error was created, not the time the method was called")
}
