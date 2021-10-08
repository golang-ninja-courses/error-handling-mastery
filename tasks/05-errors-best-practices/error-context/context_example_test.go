package errctx_test

import (
	"io"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"

	errctx "github.com/www-golang-courses-ru/advanced-dealing-with-errors-in-go/tasks/05-errors-best-practices/error-context"
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
		l.With(errorCtxToZapFields(err)...).Error("cannot do operation", zap.Error(err))
	}

	// Output:
	// {"level":"error","msg":"cannot do operation","uid":"f51be77e-f60e-11eb-b47f-1e00d13a7870","filename":"/etc/hosts","error":"EOF"}
}

func errorCtxToZapFields(err error) []zap.Field {
	ctx := errctx.From(err)
	if len(ctx) == 0 {
		return nil
	}

	fields := make([]zap.Field, 0, len(ctx))
	for k, v := range ctx {
		fields = append(fields, zap.Any(k, v))
	}
	return fields
}
