package errctx_test

import (
	"fmt"
	"io"
	"testing"

	"github.com/stretchr/testify/assert"
	errctx "github.com/www-golang-courses-ru/advanced-dealing-with-errors-in-go/tasks/05-errors-best-practices/error-context"
)

func TestAppendToNilIsNil(t *testing.T) {
	assert.Nil(t, errctx.AppendTo(nil, map[string]interface{}{"uid": "Stepan"}))
}

func TestNoErrorNoContext(t *testing.T) {
	assert.Empty(t, errctx.From(nil))
}

func TestWrapping(t *testing.T) {
	err := fmt.Errorf("read file: %w", io.EOF)
	err = errctx.AppendTo(err, map[string]interface{}{"key1": "value1", "key2": "value2"})

	err = fmt.Errorf("do foo: %w", err)
	err = errctx.AppendTo(err, map[string]interface{}{"key1": 11, "key3": 3})

	err = fmt.Errorf("do bar: %w", err)
	err = errctx.AppendTo(err, map[string]interface{}{"key4": "value4"})

	assert.ErrorIs(t, err, io.EOF)
	assert.EqualError(t, err, "do bar: do foo: read file: EOF")
	assert.Equal(t, map[string]interface{}{
		"key1": "value1",
		"key2": "value2",
		"key3": 3,
		"key4": "value4",
	}, errctx.From(err))
}

func TestImmutableCtx(t *testing.T) {
	ctx := map[string]interface{}{"key1": "value1", "key2": "value2"}
	err := errctx.AppendTo(io.EOF, ctx)
	extractedCtx := errctx.From(err)
	ctx["new_key"] = "new_value"
	assert.NotEqual(t, ctx, extractedCtx)

	extractedCtx["new_key"] = "new_value"
	newExtractedCtx := errctx.From(err)
	assert.NotEqual(t, extractedCtx, newExtractedCtx)
}
