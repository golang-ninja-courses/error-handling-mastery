package main

import "fmt"

// go run -tags pkg ./...
// go run -tags cockroach ./...
func main() {
	fmt.Printf("%+v", GimmeDeepError(2))
}
