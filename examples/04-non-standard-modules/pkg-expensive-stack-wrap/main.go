package main

import "fmt"

// go run -tags std ./...
// go run -tags pkg ./...
func main() {
	fmt.Printf("%+v", GimmeDeepError(2))
}
