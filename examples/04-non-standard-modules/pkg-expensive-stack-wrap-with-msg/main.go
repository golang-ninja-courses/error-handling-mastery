package main

import "fmt"

// go run -tags std ./...
// go run -tags pkg.msg.stack ./...
// go run -tags pkg.msg.only ./...
func main() {
	fmt.Printf("%+v", GimmeDeepError(2))
}
