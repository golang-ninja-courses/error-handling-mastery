package ops

import (
	"fmt"
)

type OperationsErrors []error

func (oe OperationsErrors) Error() string {
	return fmt.Sprintf("operations errors: %v", []error(oe))
}

type IOperation interface {
	Do() error
}

func Handle(ops ...IOperation) error {
	var opsErrs OperationsErrors

	for _, op := range ops {
		if err := op.Do(); err != nil {
			opsErrs = append(opsErrs, err)
		}
	}

	return opsErrs
}
