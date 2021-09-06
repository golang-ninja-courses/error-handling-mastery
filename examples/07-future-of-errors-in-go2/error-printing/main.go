package main

import (
	"errors"
	"fmt"
	"io"
)

var КонецФайла = errors.New("конец файла")

func main() {
	var verbs = []string{"%s", "%q", "%+q", "%v", "%#v"}

	for _, err := range []error{КонецФайла, io.EOF} {
		for _, f := range verbs {
			fmt.Printf("%3s - \t"+f, f, err)
			fmt.Println()
		}
		fmt.Println()
	}
}
