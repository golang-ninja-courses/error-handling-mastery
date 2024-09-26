package errctx_test

import (
	"fmt"
	"io"
	"testing"

	"github.com/stretchr/testify/assert"

	errctx "github.com/golang-ninja-courses/error-handling-mastery/tasks/05-errors-best-practices/error-context"
)

func TestAppendToNilIsNil(t *testing.T) {
	err := errctx.AppendTo(nil, errctx.Fields{"uid": "Stepan"})
	assert.Nil(t, err)
}

func TestNoErrorNoContext(t *testing.T) {
	ctx := errctx.From(nil)
	assert.Empty(t, ctx)
}

func TestWrapping(t *testing.T) {
	err := fmt.Errorf("read file: %w", io.EOF)
	err = errctx.AppendTo(err, errctx.Fields{"key1": "value1", "key2": "value2"})

	err = fmt.Errorf("do foo: %w", err)
	err = errctx.AppendTo(err, errctx.Fields{"key1": 11, "key3": 3})

	err = fmt.Errorf("do bar: %w", err)
	err = errctx.AppendTo(err, errctx.Fields{"key4": "value4"})

	assert.ErrorIs(t, err, io.EOF)
	assert.EqualError(t, err, "do bar: do foo: read file: EOF")

	ctx := errctx.From(err)
	assert.Equal(t, errctx.Fields{
		"key1": "value1",
		"key2": "value2",
		"key3": 3,
		"key4": "value4",
	}, ctx)
}

func TestImmutableCtx(t *testing.T) {
	ctx := errctx.Fields{"key1": "value1", "key2": "value2"}
	err := errctx.AppendTo(io.EOF, ctx)
	extractedCtx := errctx.From(err)
	ctx["new_key"] = "new_value"
	assert.NotEqual(t, ctx, extractedCtx)

	extractedCtx["new_key"] = "new_value"
	newExtractedCtx := errctx.From(err)
	assert.NotEqual(t, extractedCtx, newExtractedCtx)
}
