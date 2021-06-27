package main

import (
	"fmt"
	"strconv"

	"github.com/pkg/errors"
)

func GimmeDeepError(depth int) error {
	var err error
	if depth != 0 {
		err = GimmeDeepError(depth - 1)
		return errors.Wrap(err, "error happened on level "+strconv.Itoa(depth-1))
	}
	return errors.New("ooops, an error")
}

func main() {
	fmt.Printf("%+v", GimmeDeepError(2))
}
