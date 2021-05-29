package main

import (
	"fmt"
	"io"
)

type MyAwesomeStructWriter struct {
	Value string
}

func (m MyAwesomeStructWriter) Write(p []byte) (n int, err error) {
	panic("implement me")
}

func main() {
	var w io.Writer

	w = MyAwesomeStructWriter{Value: "value"}

	m := w.(MyAwesomeStructWriter) // 1. Прямое приведение интерфейса к типу
	fmt.Println(m.Value)

	switch m2 := w.(type) { // 2. switch по типу интерфейса
	case MyAwesomeStructWriter:
		fmt.Println(m2.Value)
	default:
		fmt.Println("have no idea what it is")
	}
}
