package main

import (
	"bufio"
	"fmt"
	"os"
)

//nolint:errcheck
func main() {
	b := bufio.NewWriter(os.Stdout)

	b.WriteString("Hello ")
	b.WriteString("World")
	b.WriteString("!")

	if err := b.Flush(); err != nil {
		fmt.Println(err)
	}

	// Hello World!
}
