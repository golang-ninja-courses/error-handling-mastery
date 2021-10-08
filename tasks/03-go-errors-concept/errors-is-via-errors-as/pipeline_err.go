package pipe

import (
	"fmt"
)

type PipelineError struct {
	User        string
	Name        string
	FailedSteps []string
}

func (p *PipelineError) Error() string {
	return fmt.Sprintf("pipeline %q error", p.Name)
}

func IsPipelineError(err error, user, pipelineName string) bool {
	// Реализуй меня.
	return false
}
