package main

import "io"

var _ io.Closer = (*CloserMock)(nil)

type CloserMock struct {
	err error
}

func (c *CloserMock) SetErr(e error) {
	c.err = e
}

func (c *CloserMock) Close() error {
	return c.err
}
