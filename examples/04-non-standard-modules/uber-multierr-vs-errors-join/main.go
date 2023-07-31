package main

import (
	"context"
	"errors"
	"fmt"
	"io"

	"go.uber.org/multierr"
)

func main() {
	err1 := multierr.Append(step1(), step2())
	err2 := multierr.Combine(step1(), step2())
	err3 := errors.Join(step1(), step2())
	for _, err := range []error{err1, err2, err3} {
		fmt.Println(err)
		fmt.Printf("%+v\n", err)
		fmt.Println()
	}

	unpacked := multierr.Errors(err3)
	fmt.Println(len(unpacked))       // 1
	fmt.Println(unpacked[0] == err3) // true

	fmt.Println()
	fmt.Println(processCloserMock_AppendInvoke(closerMock{})) // close error
	fmt.Println(processCloserMock_Join(closerMock{}))         // close error
	fmt.Println(processCloserMock_Helper(closerMock{}))       // close error
}

var errAcquireTimeout = errors.New("timeout acquiring connection from pool")

func step1() error { return errAcquireTimeout }
func step2() error { return context.Canceled }

func processCloserMock_AppendInvoke(r io.ReadCloser) (err error) {
	defer multierr.AppendInvoke(&err, multierr.Close(r))
	return nil
}

func processCloserMock_Join(r io.ReadCloser) (err error) {
	defer func() {
		err = errors.Join(err, r.Close())
	}()
	return nil
}

func processCloserMock_Helper(r io.ReadCloser) (err error) {
	defer appendInvoke(&err, r.Close)
	return nil
}

func appendInvoke(into *error, fn func() error) {
	*into = errors.Join(*into, fn())
}

var _ io.ReadCloser = closerMock{}      //
type closerMock struct{ io.ReadCloser } //
func (c closerMock) Close() error       { return errors.New("close error") }
