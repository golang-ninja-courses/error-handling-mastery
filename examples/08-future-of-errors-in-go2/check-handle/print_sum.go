package example

import (
	"fmt"
	"strconv"
)

func printSum(a, b string) error {
	x, err := strconv.Atoi(a)
	if err != nil {
		return fmt.Errorf("printSum(%q + %q): %v", a, b, err)
	}

	y, err := strconv.Atoi(b)
	if err != nil {
		return fmt.Errorf("printSum(%q + %q): %v", a, b, err)
	}

	fmt.Println("result:", x+y)
	return nil
}

// Несуществующий синтаксис.

/*
func printSum(a, b string) error {
	handle err {
		return fmt.Errorf("printSum(%q + %q): %v", a, b, err)
	}
	x := check strconv.Atoi(a)
	y := check strconv.Atoi(b)
	fmt.Println("result:", x + y)
	return nil
}
*/
