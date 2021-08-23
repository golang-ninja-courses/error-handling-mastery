// +build pkg.msg.stack

package main

import (
	"github.com/pkg/errors"
)

func GimmeDeepError(depth int) error {
	if depth == 1 {
		return errors.New("ooops, an error on level 1")
	}
	return errors.WithMessagef(GimmeDeepError(depth-1), "error happened on level %d", depth)
}
