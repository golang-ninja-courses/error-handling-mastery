package pipeline

import (
	"context"
	"errors"
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSource(t *testing.T) {
	cases := []struct {
		name         string
		input        []string
		output       []string
		errcExpected error
		errExpected  error
	}{
		{
			name:         "ok",
			input:        []string{"1", "2", "3"},
			output:       []string{"1", "2", "3"},
			errcExpected: nil,
			errExpected:  nil,
		},
		{
			name:        "empty input",
			input:       []string{},
			errExpected: errEmptyInput,
		},
		{
			name:         "invalid input",
			input:        []string{"1", "", "3"},
			output:       []string{"1"},
			errcExpected: errEmptyStringInInput,
			errExpected:  nil,
		},
	}

	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			out, errc, err := Source(context.Background(), tt.input...)

			if tt.errExpected != nil {
				assert.ErrorIs(t, err, tt.errExpected)
				return
			}

			ss := make([]string, 0, len(tt.input))
			for s := range out {
				ss = append(ss, s)
			}
			assert.Equal(t, tt.output, ss)

			assert.ErrorIs(t, tt.errcExpected, <-errc)
		})
	}
}

func TestSource_CancelledContext(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	cancel()

	out, errc, err := Source(ctx, []string{"1", "2", "3"}...)

	assert.NoError(t, err)

	var ss []string
	for s := range out {
		ss = append(ss, s)
	}

	switch len(ss) {
	case 0:
	case 1:
		assert.Equal(t, "1", ss[0])
	default:
		t.Error("out should be empty or contain single value")
	}

	assert.Empty(t, <-errc)
}

func TestParseInt(t *testing.T) {
	src := make(chan string, 3)
	go func() {
		defer close(src)
		for _, s := range []string{"1", "2", "3"} {
			src <- s
		}
	}()

	out, errc, err := ParseInt(context.Background(), 10, src)
	assert.NoError(t, err)

	var ii []int64
	for i := range out {
		ii = append(ii, i)
	}

	assert.Equal(t, []int64{1, 2, 3}, ii)
	assert.NoError(t, <-errc)
}

func TestParseInt_InvalidBaseError(t *testing.T) {
	out, errc, err := ParseInt(context.Background(), 1, nil)
	assert.ErrorIs(t, err, errInvalidBase)
	assert.Nil(t, out)
	assert.Nil(t, errc)
}

func TestParseInt_ParseError(t *testing.T) {
	src := make(chan string, 2)
	go func() {
		for _, s := range []string{"1", "a"} {
			src <- s
		}
	}()

	out, errc, err := ParseInt(context.Background(), 10, src)
	assert.NoError(t, err, errInvalidBase)
	assert.Equal(t, int64(1), <-out)
	errTarget := &strconv.NumError{}
	assert.ErrorAs(t, <-errc, &errTarget)
}

func TestParseInt_CancelledContext(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	cancel()

	src := make(chan string, 1)
	src <- "1"

	out, errc, err := ParseInt(ctx, 10, src)

	assert.NoError(t, err)

	var ii []int64
	for i := range out {
		ii = append(ii, i)
	}

	switch len(ii) {
	case 0:
	case 1:
		assert.Equal(t, "1", ii[0])
	default:
		t.Error("out should be empty or contain single value")
	}

	assert.Empty(t, <-errc)
}

func TestSink(t *testing.T) {
	cases := []struct {
		name         string
		srcFunc      func() <-chan int64
		errcExpected error
	}{
		{
			name: "ok",
			srcFunc: func() <-chan int64 {
				src := make(chan int64, 3)
				go func() {
					defer close(src)
					for _, i := range []int64{1, 2, 3} {
						src <- i
					}
				}()
				return src
			},
			errcExpected: nil,
		},
		{
			name: "",
			srcFunc: func() <-chan int64 {
				src := make(chan int64, 3)
				go func() {
					defer close(src)
					for _, i := range []int64{maxInt + 1, maxInt + 2, maxInt + 3} {
						src <- i
					}
				}()
				return src
			},
			errcExpected: errTooLargeNumber,
		},
	}

	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			errc, err := Sink(context.Background(), maxInt, tt.srcFunc())
			assert.NoError(t, err)
			assert.ErrorIs(t, <-errc, tt.errcExpected)
		})
	}
}

func TestMerge(t *testing.T) {
	errs := []error{errors.New("1"), errors.New("2"), errors.New("3")}
	errcs := make([]<-chan error, 0, 3)
	for _, err := range errs {
		errc := make(chan error, 1)
		errc <- err
		close(errc)

		errcs = append(errcs, errc)
	}

	out := Merge(errcs...)

	for err := range out {
		assert.Contains(t, errs, err)
	}
}

func TestRun(t *testing.T) {
	cases := []struct {
		name        string
		ss          []string
		expectError bool
	}{
		{
			name:        "ok",
			ss:          []string{"1", "2", "3", "4", "5"},
			expectError: false,
		},
		{
			name:        "error",
			ss:          []string{"1", "2", "a", "4", "5"},
			expectError: true,
		},
	}

	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			err := Pipeline(tt.ss...)
			if tt.expectError {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}
