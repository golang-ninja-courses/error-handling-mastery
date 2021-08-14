package main

import (
	"errors"
	"fmt"
	"os"
)

func main() {
	var err error
	var osErr *os.PathError
	if errors.As(err, osErr) { // Казалось бы, уже указатель!
		fmt.Println(osErr.Path) // До этой строчки не дойдёт :(
	}
}
