package pkgb

import "github.com/www-golang-courses-ru/advanced-dealing-with-errors-in-go/examples/05-errors-best-practices/constant-errors-diff-pkgs/common"

type err string

func (e err) Error() string { return string(e) }

const (
	ErrInvalidHost = err("invalid host")
	ErrUnknownData = common.Error("unknown data")
)
