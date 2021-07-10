package wrapping_opaque_errors

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsTemporary(t *testing.T) {
	errNetwork := &NetworkError{}
	assert.True(t, IsTemporary(errNetwork))

	errWrapped := fmt.Errorf("simple wrap: %w", errNetwork)
	assert.True(t, IsTemporary(errWrapped))

	errWithMessage := &WithMessageError{
		message: "any message",
		err:     errWrapped,
	}
	assert.True(t, IsTemporary(errWithMessage))
}
