package main

import (
	"context"
	"errors"
	"fmt"

	"github.com/hashicorp/go-multierror"
)

func main() {
	err1 := errorsJoin()
	fmt.Printf("%#v\n", err1)
	fmt.Println(err1)

	err2 := hashicorpMultierr()
	fmt.Printf("%#v\n", err2)
	fmt.Println(err2)

	err2.(*multierror.Error).ErrorFormat = StdFormat
	fmt.Println(err2)

	fmt.Println(errors.Is(err1, errAcquireTimeout))
	fmt.Println(errors.Is(err2, errAcquireTimeout))
}

func errorsJoin() error {
	// Без врапинга:
	// return errors.Join(step1(), step2())

	var result error
	if err := step1(); err != nil {
		result = errors.Join(result, fmt.Errorf("step 1: %w", err))
	}
	if err := step2(); err != nil {
		result = errors.Join(result, fmt.Errorf("step 2: %w", err))
	}
	return result
}

func hashicorpMultierr() error {
	var result error
	if err := step1(); err != nil {
		result = multierror.Append(result, fmt.Errorf("step 1: %w", err))
	}
	if err := step2(); err != nil {
		result = multierror.Append(result, fmt.Errorf("step 2: %w", err))
	}
	return result
}

func Process(input []string) error {
	var result error
	for _, s := range input {
		if err := process(s); err != nil {
			result = multierror.Append(result, err)
		}
	}
	return result
}

var errAcquireTimeout = errors.New("timeout acquiring connection from pool")

func step1() error { return errAcquireTimeout }
func step2() error { return context.Canceled }

var StdFormat = multierror.ErrorFormatFunc(func(errs []error) string {
	var b []byte
	for i, err := range errs {
		if i > 0 {
			b = append(b, '\n')
		}
		b = append(b, err.Error()...)
	}
	return string(b)
})
