package main

import (
	"fmt"
	"os"

	"github.com/pkg/errors"
)

func main() {
	wrappedErr := os.ErrNotExist

	cases := []struct {
		title string
		err   error
	}{
		{
			title: "only std errors",
			err:   fmt.Errorf("msg 2: %w", fmt.Errorf("msg 1: %w", wrappedErr)),
		},
		{
			title: "only pkg/errors errors",
			err:   errors.Wrap(errors.Wrap(os.ErrNotExist, "msg 1"), " msg 2"),
		},
		{
			title: "combined 1",
			err:   errors.Wrap(fmt.Errorf("msg 1: %w", os.ErrNotExist), "msg 2"),
		},
		{
			title: "combined 2",
			err:   fmt.Errorf("msg 2: %w", errors.Wrap(os.ErrNotExist, "msg 1")),
		},
	}

	for _, c := range cases {
		fmt.Println(c.title)
		fmt.Println("\terrors.Is:", errors.Is(c.err, wrappedErr))
		fmt.Println("\terrors.Cause:", errors.Cause(c.err) == wrappedErr)
		fmt.Println()
	}
}
