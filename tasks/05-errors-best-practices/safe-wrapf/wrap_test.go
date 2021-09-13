package errors

import (
	"io"
	"testing"

	"github.com/pkg/errors"
)

// Тесты ниже всегда проходят. Они для самостоятельной проверки глазами.

func TestWrapf_MultipleWraps(t *testing.T) {
	err := io.EOF
	for i := 0; i < 10; i++ {
		err = Wrapf(err, "wrap %d", i)
	}

	t.Logf("%+v", err)
	/*
	 wrap_example_test.go:14: EOF
	        wrap 0
	        .../tasks/05-errors-best-practices/safe-wrapf.Wrapf
	        	.../tasks/05-errors-best-practices/safe-wrapf/wrap.go:13
	        .../tasks/05-errors-best-practices/safe-wrapf.TestWrapf
	        	.../tasks/05-errors-best-practices/safe-wrapf/wrap_example_test.go:15
	        testing.tRunner
	        	/usr/local/go/src/testing/testing.go:1193
	        runtime.goexit
	        	/usr/local/go/src/runtime/asm_arm64.s:1130
	        wrap 1
	        wrap 2
	        wrap 3
	        wrap 4
	        wrap 5
	        wrap 6
	        wrap 7
	        wrap 8
	        wrap 9
	*/
}

func TestWrapf_AlreadyWithStack(t *testing.T) {
	err := errors.WithStack(io.EOF)
	for i := 0; i < 10; i++ {
		err = Wrapf(err, "wrap %d", i)
	}

	t.Logf("%+v", err)
	/*
	   wrap_example_test.go:48: EOF
	       .../tasks/05-errors-best-practices/safe-wrapf.TestWrapf_AlreadyWithStack
	       	.../tasks/05-errors-best-practices/safe-wrapf/wrap_example_test.go:43
	       testing.tRunner
	       	/usr/local/go/src/testing/testing.go:1193
	       runtime.goexit
	       	/usr/local/go/src/runtime/asm_arm64.s:1130
	       wrap 0
	       wrap 1
	       wrap 2
	       wrap 3
	       wrap 4
	       wrap 5
	       wrap 6
	       wrap 7
	       wrap 8
	       wrap 9
	*/
}
