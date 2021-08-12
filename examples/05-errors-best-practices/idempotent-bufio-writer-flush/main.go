package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	f, err := os.Open("/etc/hosts")
	if err != nil {
		panic(err)
	}
	if err := f.Close(); err != nil {
		panic(err)
	}

	w := bufio.NewWriter(f)
	w.WriteString("bad")

	for i := 0; i < 5; i++ {
		if err := w.Flush(); err != nil {
			fmt.Println(err)
		}
	}

	/*
		write /etc/hosts: file already closed
		write /etc/hosts: file already closed
		write /etc/hosts: file already closed
		write /etc/hosts: file already closed
		write /etc/hosts: file already closed
	*/
}
