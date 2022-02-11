package tmpl

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"testing"
	"text/template"
	"unsafe"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestIsExecUnexportedFieldError(t *testing.T) {
	t.Run("no error", func(t *testing.T) {
		assert.False(t, IsExecUnexportedFieldError(nil))
	})

	t.Run("not `unexported field` error", func(t *testing.T) {
		err := parseAndExecuteTemplate(bytes.NewBuffer(nil), "example",
			`{{ with .Name }}`,
			struct{ Name unsafe.Pointer }{},
		)
		require.Error(t, err)
		assert.False(t, IsExecUnexportedFieldError(err))
		assert.False(t, IsExecUnexportedFieldError(fmt.Errorf("parse and execute tmpl: %w", err)))
	})

	t.Run("partial assertion is not working", func(t *testing.T) {
		for _, err := range []error{
			errors.New("is an unexported field of struct type"),
			errors.New("template"),
		} {
			assert.False(t, IsExecUnexportedFieldError(err))
			assert.False(t, IsExecUnexportedFieldError(fmt.Errorf("%w", err)))
		}
	})

	t.Run("not template.Exec error", func(t *testing.T) {
		errFraudulent := errors.New(`template: example:1:3: executing "example" at <.name>: name is an unexported field of struct type`)
		assert.False(t, IsExecUnexportedFieldError(errFraudulent))
		assert.False(t, IsExecUnexportedFieldError(fmt.Errorf("%w", errFraudulent)))
	})

	t.Run("`unexported field` error", func(t *testing.T) {
		err := parseAndExecuteTemplate(bytes.NewBuffer(nil), "example",
			`{{ .name }}`,
			struct{ name string }{name: "Bob"},
		)
		require.Error(t, err)
		assert.True(t, IsExecUnexportedFieldError(err))
		assert.True(t, IsExecUnexportedFieldError(fmt.Errorf("parse and execute tmpl: %w", err)))
	})
}

func TestIsFunctionNotDefinedError(t *testing.T) {
	t.Run("no error", func(t *testing.T) {
		assert.False(t, IsFunctionNotDefinedError(nil))
	})

	t.Run("not `function not defined` error", func(t *testing.T) {
		err := parseAndExecuteTemplate(bytes.NewBuffer(nil), "example", `{{ with`, nil)
		require.Error(t, err)
		assert.False(t, IsFunctionNotDefinedError(err))
		assert.False(t, IsFunctionNotDefinedError(fmt.Errorf("parse and execute tmpl: %w", err)))
	})

	t.Run("partial assertion is not working", func(t *testing.T) {
		for _, err := range []error{
			errors.New("function"),
			errors.New("not defined"),
			errors.New("template"),
		} {
			assert.False(t, IsFunctionNotDefinedError(err))
			assert.False(t, IsFunctionNotDefinedError(fmt.Errorf("%w", err)))
		}
	})

	t.Run("`function not defined` error", func(t *testing.T) {
		err := parseAndExecuteTemplate(bytes.NewBuffer(nil), "example", `{{ call XXX }}`, nil)
		require.Error(t, err)
		assert.True(t, IsFunctionNotDefinedError(err))
		assert.True(t, IsFunctionNotDefinedError(fmt.Errorf("parse and execute tmpl: %w", err)))
	})
}

func parseAndExecuteTemplate(wr io.Writer, name, text string, data interface{}) error {
	t, err := template.New(name).Parse(text)
	if err != nil {
		return err
	}
	return t.Execute(wr, data)
}
