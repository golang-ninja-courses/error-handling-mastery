package pipe

import "fmt"

type UserError struct {
	Operation string
	User      string
}

func (u *UserError) Error() string {
	return fmt.Sprintf("user %s cannot do op %s", u.User, u.Operation)
}

type PipelineError struct {
	User        string
	Name        string
	FailedSteps []string
}

func (p *PipelineError) Error() string {
	return fmt.Sprintf("pipeline %q error", p.Name)
}

// Добавь методов для структуры PipelineError.
