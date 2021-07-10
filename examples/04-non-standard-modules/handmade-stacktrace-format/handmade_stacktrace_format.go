package main

import (
	"fmt"
	"io"
	"runtime"
	"strconv"
)

const (
	maxStacktraceDepth = 32
	skipFrames         = 2
)

type Frame uintptr

func (f Frame) pc() uintptr {
	return uintptr(f) - 1
}

func (f Frame) function() string {
	fn := runtime.FuncForPC(f.pc())
	if fn == nil {
		return "unknown"
	}
	return fn.Name()
}

func (f Frame) file() string {
	fn := runtime.FuncForPC(f.pc())
	if fn == nil {
		return "unknown"
	}
	file, _ := fn.FileLine(f.pc())
	return file
}

func (f Frame) line() int {
	fn := runtime.FuncForPC(f.pc())
	if fn == nil {
		return 0
	}
	_, line := fn.FileLine(f.pc())
	return line
}

func (f Frame) Format(s fmt.State, verb rune) {
	switch verb {
	case 's':
		io.WriteString(s, f.function())
		io.WriteString(s, "()\n\t")
		io.WriteString(s, f.file())
	case 'd':
		io.WriteString(s, strconv.Itoa(f.line()))
	case 'v':
		f.Format(s, 's')
		io.WriteString(s, ":")
		f.Format(s, 'd')
	}
}

type StackTrace []Frame

func (s StackTrace) Format(st fmt.State, verb rune) {
	switch verb {
	case 'v':
		for i, pc := range s {
			fmt.Fprintf(st, "%v", pc)
			if i != len(s) {
				fmt.Fprint(st, "\n")
			}
		}
	}
}

func CallersFrames() StackTrace {
	var pcs [maxStacktraceDepth]uintptr
	n := runtime.Callers(skipFrames, pcs[:]) // пропускаем фреймы runtime.Callers() и CallersFrames()
	callers := pcs[0:n]

	frames := make([]Frame, len(callers))
	for i := range callers {
		frames[i] = Frame(callers[i])
	}
	return frames
}

func stubFunc1() {
	stubFunc2()
}

func stubFunc2() {
	fmt.Printf("%+v", CallersFrames())
}

func main() {
	stubFunc1()
}
