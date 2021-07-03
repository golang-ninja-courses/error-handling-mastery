package main

type Error string

func (e Error) Error() string { return string(e) }

const EOF = Error("EOF")

func main() {
	// cannot assign to EOF (declared const)
	// EOF = Error("EOF2")
}
