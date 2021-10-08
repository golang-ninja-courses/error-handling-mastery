package commandexecutor

import "errors"

var (
	ErrUnsupportedCommand = errors.New("unsupported command")
	ErrCommandTimeout     = errors.New("command timeout")
)
