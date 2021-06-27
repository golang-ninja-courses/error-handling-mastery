package main

import (
	"errors"
	"fmt"
)

func Wrap(err *error, f string, v ...interface{}) {
	if *err != nil {
		*err = fmt.Errorf(f+": %w", append(v, *err)...)
	}
}

func main() {
	var err error
	Wrap(&err, "oops")
	fmt.Println(err)

	err = errors.New("error happened")
	Wrap(&err, "oops")
	fmt.Println(err)
}
