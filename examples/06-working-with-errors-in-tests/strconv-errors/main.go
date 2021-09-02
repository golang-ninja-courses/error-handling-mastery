package main

import (
	"errors"
	"fmt"
	"strconv"
)

func main() {
	// strconv.ParseInt: parsing "1": invalid base 123
	_, err := strconv.ParseInt("1", 123, 32)
	fmt.Println(err)

	var numErr *strconv.NumError
	if errors.As(err, &numErr) {
		fmt.Printf("%v (%T)\n", numErr.Err, numErr.Err) // invalid base 123 (*errors.errorString)
	}

	// strconv.ParseInt: parsing "1": invalid bit size -1
	_, err = strconv.ParseInt("1", 10, -1)
	fmt.Println(err)
}
