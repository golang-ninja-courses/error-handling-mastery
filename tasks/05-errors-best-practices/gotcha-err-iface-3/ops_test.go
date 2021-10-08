package ops

import (
	"errors"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func ExampleHandle() {
	if err := Handle(successOperation{}); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("no operations errors")
	}

	// Output:
	// no operations errors
}

func TestHandle(t *testing.T) {
	t.Run("no operations - no errors", func(t *testing.T) {
		err := Handle()
		assert.NoError(t, err)
	})

	t.Run("all operations are success - no errors", func(t *testing.T) {
		err := Handle(successOperation{}, successOperation{})
		assert.NoError(t, err)
	})

	t.Run("has error operations", func(t *testing.T) {
		err := Handle(successOperation{}, errOperation{}, errOperation{}, successOperation{})
		assert.Error(t, err)

		var opsErrs OperationsErrors
		assert.ErrorAs(t, err, &opsErrs)
		assert.Len(t, opsErrs, 2)
	})
}

type successOperation struct{}

func (successOperation) Do() error {
	return nil
}

type errOperation struct{}

func (errOperation) Do() error {
	return errors.New("something wrong")
}
