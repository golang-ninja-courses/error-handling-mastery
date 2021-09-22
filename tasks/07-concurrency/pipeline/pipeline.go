package pipeline

import (
	"context"
	"errors"
	_ "fmt"
	_ "strconv"
	_ "sync"
)

const (
	maxInt = 100
)

var (
	errEmptyInput         = errors.New("input strings couldn't be empty")
	errEmptyStringInInput = errors.New("input shouldn't contain empty strings")
	errInvalidBase        = errors.New("invalid base")
	errTooLargeNumber     = errors.New("too large number")
)

func Source(ctx context.Context, input ...string) (<-chan string, <-chan error, error) {
	// Реализовать
	return nil, nil, nil
}

func ParseInt(ctx context.Context, base int, in <-chan string) (<-chan int64, <-chan error, error) {
	// Реализовать
	return nil, nil, nil
}

func Sqr(ctx context.Context, in <-chan int64) (<-chan int64, <-chan error, error) {
	// Реализовать
	return nil, nil, nil
}

func Sink(_ context.Context, maxInt int64, in <-chan int64) (<-chan error, error) {
	// Реализовать
	return nil, nil
}

func Merge(errcs ...<-chan error) <-chan error {
	// Реализовать
	return nil
}

func Pipeline(ss ...string) error {
	// Реализовать
	return nil
}
