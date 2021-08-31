package main

import (
	"errors"
	"fmt"
	"io"
	"strings"

	"github.com/hashicorp/go-multierror"
)

func main() {
	err1 := errors.New("an error 1")
	err2 := errors.New("an error 2")

	err := multierror.Append(io.EOF, err1, err2)

	fmt.Println(errors.Is(err, io.EOF)) // true
	fmt.Println(errors.Is(err, err1))   // true
	fmt.Println(errors.Is(err, err2))   // true
	fmt.Println()

	fmt.Println(err)
	/*
		3 errors occurred:
		        * EOF
		        * an error 1
		        * an error 2
	*/

	err.ErrorFormat = func(errors []error) string {
		var b strings.Builder
		b.WriteString("MY ERRORS:\n")
		for _, err := range errors {
			b.WriteString("\t - " + err.Error() + "\n")
		}
		return b.String()
	}
	fmt.Println(err)
	/*
		MY ERRORS:
		         - EOF
		         - an error 1
		         - an error 2
	*/
}
