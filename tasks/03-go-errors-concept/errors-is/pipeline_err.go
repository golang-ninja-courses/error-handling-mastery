package pipe

import "fmt"

type PipelineErr struct {
	User        string
	Name        string
	FailedSteps []string
}

func (p *PipelineErr) Error() string {
	return fmt.Sprintf("pipeline %q error", p.Name)
}

// Добавь методов для структуры PipelineErr.
