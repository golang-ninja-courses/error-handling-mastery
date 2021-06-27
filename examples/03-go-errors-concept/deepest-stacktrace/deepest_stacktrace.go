package main

import (
	"fmt"
	"strings"

	"github.com/cockroachdb/sentry-go"
	"github.com/pkg/errors"
)

func getError(depth int) error {
	var err error
	if depth != 0 {
		err = getError(depth - 1)
		return errors.Wrap(err, fmt.Sprintf("%d wrap", depth-1))
	}
	return errors.New("ooops, an error")
}

func GetDeepestStackTrace(err error) *sentry.Stacktrace {
	chain := make([]error, 0, 16)
	for err != nil {
		chain = append(chain, err)
		err = errors.Unwrap(err)
	}

	for i := len(chain) - 1; i >= 0; i-- {
		if stacktrace := sentry.ExtractStacktrace(chain[i]); stacktrace != nil {
			return stacktrace
		}
	}

	return nil
}

func StackTraceString(stacktrace *sentry.Stacktrace) string {
	if stacktrace == nil {
		return ""
	}

	frames := stacktrace.Frames
	builder := strings.Builder{}

	for i := len(frames) - 1; i >= 0; i-- {
		builder.WriteString(fmt.Sprintf("%s.%s\n", frames[i].Module, frames[i].Function))
		builder.WriteString(fmt.Sprintf("\t%s:%d\n", frames[i].AbsPath, frames[i].Lineno))
	}
	return builder.String()
}

func main() {
	err := getError(10)
	stacktrace := GetDeepestStackTrace(err)

	fmt.Printf("%+v\n\n", err)                       // Будет выведен стектрейс по всем ошибкам
	fmt.Printf("%s\n", StackTraceString(stacktrace)) // Будет выведен стектрейс корневой ошибки
}
