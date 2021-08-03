package secondary_error_format

import (
	"errors"
	"fmt"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCombineErrors(t *testing.T) {
	for _, test := range []struct {
		name           string
		doSomething    func() error
		handleError    func() error
		expectedFormat string
	}{
		{
			name:           "primary and secondary",
			doSomething:    func() error { return os.ErrExist },
			handleError:    func() error { return errors.New("ah shit, here we go again") },
			expectedFormat: "file already exists\nsecondary: ah shit, here we go again",
		},
		{
			name:           "only primary",
			doSomething:    func() error { return os.ErrExist },
			handleError:    func() error { return nil },
			expectedFormat: "file already exists",
		},
		{
			name:           "nil errors",
			doSomething:    func() error { return nil },
			handleError:    func() error { return nil },
			expectedFormat: "<nil>",
		},
	} {
		t.Run(test.name, func(t *testing.T) {
			err := test.doSomething()
			if err != nil {
				if err2 := test.handleError(); err2 != nil {
					err = CombineErrors(err, err2)
				}
			}

			actualOutput := fmt.Sprintf("%+v", err)

			assert.Equal(t, test.expectedFormat, actualOutput)
		})
	}
}
