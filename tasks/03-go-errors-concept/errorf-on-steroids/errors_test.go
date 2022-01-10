package errs

import (
	"errors"
	"fmt"
	"io"
	"net"
	"syscall"
	"testing"

	"github.com/stretchr/testify/require"
)

func ExampleErrorf() {
	err := Errorf("cannot load file %q: %w through %w", "file.txt",
		&net.AddrError{Err: "err", Addr: "0.0.0.0:4242"},
		syscall.ECONNREFUSED)

	if err == nil {
		panic("invalid realization of Errorf")
	}
	fmt.Println(err)
	fmt.Println(errors.Is(err, syscall.ECONNREFUSED))

	var addrErr *net.AddrError
	fmt.Println(errors.As(err, &addrErr))
	fmt.Println(addrErr.Addr)

	// Output:
	// cannot load file "file.txt": address 0.0.0.0:4242: err through connection refused
	// true
	// true
	// 0.0.0.0:4242
}

func TestErrorf(t *testing.T) {
	t.Run("error msg formatting", func(t *testing.T) {
		err := Errorf("worker %d: cannot load file %q from %s to %v: error: %t: %w through %w",
			42,
			"file.txt",
			"localhost:4242/file",
			"/tmp/file.txt",
			true,
			io.ErrClosedPipe,
			syscall.EACCES,
		)
		require.Error(t, err)
		require.EqualError(t, err, `worker 42: cannot load file "file.txt" from localhost:4242/file to /tmp/file.txt: `+
			"error: true: io: read/write on closed pipe through permission denied")
	})

	t.Run("errors.Is is working", func(t *testing.T) {
		err := Errorf("cannot sync files in %s: %w through %w", "/tmp", io.ErrClosedPipe, syscall.EACCES)
		require.Error(t, err)
		require.ErrorIs(t, err, io.ErrClosedPipe)
		require.ErrorIs(t, err, syscall.EACCES)
	})

	t.Run("errors.Is is not working for not-w errors", func(t *testing.T) {
		err := Errorf("cannot sync files: %v (or %s) through %w", io.ErrClosedPipe, io.ErrUnexpectedEOF, syscall.EACCES)
		require.Error(t, err)
		require.EqualError(t, err,
			`cannot sync files: io: read/write on closed pipe (or unexpected EOF) through permission denied`)
		require.NotErrorIs(t, err, io.ErrClosedPipe)
		require.NotErrorIs(t, err, io.ErrUnexpectedEOF)
		require.ErrorIs(t, err, syscall.EACCES)
	})

	t.Run("errors.As is working", func(t *testing.T) {
		err := Errorf("cannot load file %q: %w through %w", "file.txt",
			&net.AddrError{Addr: "0.0.0.0:4242"},
			syscall.ECONNREFUSED)

		require.Error(t, err)
		require.ErrorIs(t, err, syscall.ECONNREFUSED)

		var addrErr *net.AddrError
		require.ErrorAs(t, err, &addrErr)
		require.Equal(t, "0.0.0.0:4242", addrErr.Addr)
	})

	t.Run("errors.As is not working for not-w errors", func(t *testing.T) {
		err := Errorf("cannot load file: %v (or %s) through %w",
			&net.AddrError{Addr: "0.0.0.0:4242", Err: "err"},
			&net.DNSError{Name: "0.0.0.0", Err: "no such host"},
			syscall.ECONNREFUSED)

		require.Error(t, err)
		require.EqualError(t, err,
			`cannot load file: address 0.0.0.0:4242: err (or lookup 0.0.0.0: no such host) through connection refused`)
		require.ErrorIs(t, err, syscall.ECONNREFUSED)

		var addrErr *net.AddrError
		require.False(t, errors.As(err, &addrErr))

		var dnsErr *net.DNSError
		require.False(t, errors.As(err, &dnsErr))
	})

	t.Run("errors.As returns the first suitable err", func(t *testing.T) {
		err := Errorf("cannot load file %q: %w through %w", "file.txt",
			&net.AddrError{Addr: "0.0.0.0:4242"},
			&net.AddrError{Addr: "0.0.0.0:4243"})

		var addrErr *net.AddrError
		require.ErrorAs(t, err, &addrErr)
		require.Equal(t, "0.0.0.0:4242", addrErr.Addr)
	})

	t.Run("multiple Errorf", func(t *testing.T) {
		err := Errorf("cannot read file: %w", io.EOF)
		err = Errorf("cannot load config: %w", err)
		err = Errorf("cannot start app: %w", err)

		require.Error(t, err)
		require.ErrorIs(t, err, io.EOF)
		require.EqualError(t, err, "cannot start app: cannot load config: cannot read file: EOF")
	})

	t.Run("std wrapping", func(t *testing.T) {
		err := fmt.Errorf("cannot read file: %w", io.EOF)
		err = Errorf("cannot load config: %w through %w", err, syscall.EACCES)
		err = fmt.Errorf("cannot start app: %w", err)

		require.Error(t, err)
		require.ErrorIs(t, err, io.EOF)
		require.ErrorIs(t, err, syscall.EACCES)
		require.EqualError(t, err,
			"cannot start app: cannot load config: cannot read file: EOF through permission denied")
	})

	t.Run("not error argument", func(t *testing.T) {
		err := Errorf("cannot sync files in %s: %w through %w", "/tmp",
			io.ErrClosedPipe.Error(),
			syscall.EACCES,
		)
		require.Error(t, err)
		require.False(t, errors.Is(err, io.ErrClosedPipe))
		require.True(t, errors.Is(err, syscall.EACCES))
		require.EqualError(t, err,
			"cannot sync files in /tmp: io: read/write on closed pipe through permission denied")
	})

	t.Run("nil arguments", func(t *testing.T) {
		err := Errorf("cannot sync files in %s: %w through %w", "/tmp", nil, nil)
		require.NoError(t, err)

		err = Errorf("cannot sync files in %s: %w through %w", "/tmp", nil, io.ErrClosedPipe)
		require.Error(t, err)
		require.ErrorIs(t, err, io.ErrClosedPipe)
		require.EqualError(t, err,
			"cannot sync files in /tmp: <nil> through io: read/write on closed pipe")
	})
}
