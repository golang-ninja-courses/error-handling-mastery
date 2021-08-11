package main

import (
	"bufio"
	"fmt"
	"os"
)

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
