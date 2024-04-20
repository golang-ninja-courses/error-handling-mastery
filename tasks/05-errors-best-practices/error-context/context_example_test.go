package errctx_test

import (
	"io"
	"sort"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"

	errctx "github.com/golang-ninja-courses/error-handling-mastery/tasks/05-errors-best-practices/error-context"
)

func ExampleAppendTo() {
	bar := func() error {
		return errctx.AppendTo(io.EOF, errctx.Fields{
			"uid": "f51be77e-f60e-11eb-b47f-1e00d13a7870",
		})
	}

	foo := func() error {
		if err := bar(); err != nil {
			return errctx.AppendTo(err, errctx.Fields{
				"uid":      "unknown",
				"filename": "/etc/hosts",
			})
		}
		return nil
	}

	c := zap.Config{
		Level:    zap.NewAtomicLevelAt(zap.InfoLevel),
		Encoding: "json",
		EncoderConfig: zapcore.EncoderConfig{
			LevelKey:    "level",
			NameKey:     "logger",
			MessageKey:  "msg",
			LineEnding:  zapcore.DefaultLineEnding,
			EncodeLevel: zapcore.LowercaseLevelEncoder,
		},
		OutputPaths:      []string{"stdout"},
		ErrorOutputPaths: []string{"stdout"},
	}
	l, _ := c.Build() // No options - no error.

	if err := foo(); err != nil {
		l.With(errorCtxToZapFields(err)...).Error("cannot do operation")
	}

	// Output:
	// {"level":"error","msg":"cannot do operation","error":"EOF","filename":"/etc/hosts","uid":"f51be77e-f60e-11eb-b47f-1e00d13a7870"}
}

func errorCtxToZapFields(err error) []zap.Field {
	ctx := errctx.From(err)
	if len(ctx) == 0 {
		return nil
	}

	fields := make([]zap.Field, 0, len(ctx)+1)
	for k, v := range ctx {
		fields = append(fields, zap.Any(k, v))
	}
	fields = append(fields, zap.Error(err))

	sort.Sort(byKey(fields))
	return fields
}

type byKey []zap.Field

func (b byKey) Len() int           { return len(b) }
func (b byKey) Less(i, j int) bool { return b[i].Key < b[j].Key }
func (b byKey) Swap(i, j int)      { b[i], b[j] = b[j], b[i] }
