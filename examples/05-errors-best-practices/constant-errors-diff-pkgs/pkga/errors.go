package pkga

import "github.com/golang-ninja-courses/error-handling-mastery/examples/05-errors-best-practices/constant-errors-diff-pkgs/common"

type err string

func (e err) Error() string { return string(e) }

const (
	ErrInvalidHost = err("invalid host")
	ErrUnknownData = common.Error("unknown data")
)
