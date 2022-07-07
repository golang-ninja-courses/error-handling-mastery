// https://github.com/cockroachdb/errors/issues/97
package main

import (
	"fmt"
	"github.com/cockroachdb/errors"
)

type SimpleWrapper struct {
	err error
}

func (w SimpleWrapper) Error() string {
	return "boom!"
}

func (w SimpleWrapper) Unwrap() error {
	return w.err
}

func main() {
	stack := errors.WithStack

	ref := stack(stack(SimpleWrapper{}))
	err := stack(stack(SimpleWrapper{err: stack(errors.New("boom!"))}))

	if errors.IsAny(err, ref) {
		fmt.Println("gotcha!")
	}

	/* panic: runtime error: index out of range [3] with length 3

	goroutine 1 [running]:
	github.com/cockroachdb/errors/markers.equalMarks(...)
	        github.com/cockroachdb/errors@v1.9.0/markers/markers.go:205
	github.com/cockroachdb/errors/markers.IsAny({0x102802528, 0x1400000e438}, {0x14000167f48, 0x1, 0x14000167f28?})
	        github.com/cockroachdb/errors@v1.9.0/markers/markers.go:186 +0x364
	github.com/cockroachdb/errors.IsAny(...)
	        github.com/cockroachdb/errors@v1.9.0/markers_api.go:64
	main.main()
	        examples/04-non-standard-modules/cockroach-is-any-bug/main.go:26 +0x318
	*/
}
