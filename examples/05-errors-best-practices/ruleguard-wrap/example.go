package example

import "github.com/pkg/errors"

// $ ruleguard -rules gorules/rules.go example.go

func foo() error {
	if err := bar(); err != nil {
		return err
	}
	return nil
}

func bar() error {
	return errors.Wrap(errors.New("bad request"), "do request")
}
