//go:build pkg.msg.only
// +build pkg.msg.only

package main

import (
	stderrors "errors"

	"github.com/pkg/errors"
)

func GimmeDeepError(depth int) error {
	if depth == 1 {
		return stderrors.New("ooops, an error on level 1")
	}
	return errors.WithMessagef(GimmeDeepError(depth-1), "error happened on level %d", depth)
}
