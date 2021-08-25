// +build cockroach

package main

import (
	"github.com/cockroachdb/errors"
)

func GimmeDeepError(depth int) error {
	if depth == 1 {
		return errors.New("ooops, an error on level 1")
	}
	return errors.Wrapf(GimmeDeepError(depth-1), "error happened on level %d", depth)
}
