package main

import (
	"errors"
	"fmt"
	"net"
)

type PipelineError struct {
	User        string
	Name        string
	FailedSteps []string
}

func (p *PipelineError) Error() string {
	return fmt.Sprintf("pipeline %q error", p.Name)
}

func main() {
	var p error = &PipelineError{}
	// p := &PipelineError{}
	if errors.As(net.UnknownNetworkError("tdp"), &p) {
		fmt.Println(p)
	} else {
		fmt.Println("As() return false")
	}
}
