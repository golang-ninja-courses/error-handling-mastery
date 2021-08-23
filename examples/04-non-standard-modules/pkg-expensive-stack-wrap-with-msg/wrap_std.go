// +build std

package main

import (
	"errors"
	"fmt"
)

func GimmeDeepError(depth int) error {
	if depth == 1 {
		return errors.New("ooops, an error on level 1")
	}
	return fmt.Errorf("error happened on level %d: %w", depth, GimmeDeepError(depth-1))
}
