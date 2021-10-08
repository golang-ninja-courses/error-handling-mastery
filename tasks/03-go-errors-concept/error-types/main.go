package main

import (
	"context"
	"fmt"
	"io"
	"io/fs"
	"net"
	"syscall"
)

func main() {
	for _, e := range []error{
		net.UnknownNetworkError("usp"),
		fmt.Errorf("cannot read from file: %v", io.EOF),
		io.EOF,
		syscall.EDOM,
		fmt.Errorf("cannot save HTML page: %v",
			fmt.Errorf("cannot fetch URL: %v", context.Canceled)),
		&fs.PathError{Op: "seek", Path: "/etc/hosts", Err: fs.ErrInvalid},
	} {
		fmt.Printf("%-30T%q\n", e, e)
	}
}
