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

func (u *UserError) Redirect(t *PipelineError) {
	//	*(&u.Operation) = t.Name
	//	*(&u.User) = t.User
	// x := *t
	// uO := u.Operation
	// u.User = x.User
	// uO = x.Name
	// u.Operation = x.Name
	// u.User = t.User
	// t.Name = u.Operation // wrong
	// newHead := &PipelineError{u.User, u.Operation, nil}
	// *u = newHead
}

func (p *PipelineError) As(target any) bool {
	// if target == nil {
	// 	return false
	// }
	f2, ok := target.(**UserError)
	if ok {
		*f2 = &UserError{
			Operation: p.Name,
			User:      p.User,
		}
		return true
	}
	return false
}
