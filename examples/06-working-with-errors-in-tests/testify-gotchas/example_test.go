package main

import (
	"context"
	"errors"
	"io"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestEqualErrors(t *testing.T) {
	MyEOF := errors.New(io.EOF.Error())
	require.Equal(t, MyEOF, io.EOF) // Хотелось бы, чтобы тест не прошёл.
}

func TestErrorInsteadOfErrorIs(t *testing.T) {
	someOperation := func() error {
		// Попробуйте:
		// return nil
		return io.EOF
	}

	err := someOperation()
	require.Error(t, err, context.DeadlineExceeded) // Хотелось бы, чтобы тест не прошёл.
}

func TestErrorIsAtHome(t *testing.T) {
	someOperation := func() error {
		return io.EOF
	}

	err := someOperation()
	// Обратите внимание на сообщение об ошибке!
	// Без него будет сложно понять, почему errors.Is вернул false.
	require.True(t, errors.Is(err, context.DeadlineExceeded), "actual err: %v", err)
}
