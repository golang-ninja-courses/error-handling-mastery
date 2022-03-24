package chain

import (
	"errors"
	"fmt"
	"io"
	"testing"

	"github.com/stretchr/testify/assert"
)

func ExampleProcessMessage() {
	if errors.Is(ProcessMessage(), io.ErrShortWrite) {
		fmt.Println("chain is not broken")
	} else {
		fmt.Println("chain is broken")
	}

	// Output:
	// chain is not broken
}

func TestProcessMessage(t *testing.T) {
	err := ProcessMessage()
	assert.Error(t, err)
	assert.ErrorIs(t, err, io.ErrShortWrite)
	assert.NotErrorIs(t, err, io.EOF)
	assert.EqualError(t, err,
		`cannot process msg: cannot write data: save msg "8fbad38c-c5c5-11eb-b876-1e00d13a7870" error: short write`)
}
