package pipeline

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHash_String(t *testing.T) {
	h := newHash([]byte("hello"))
	assert.Equal(t, "9595c9df90075148eb06860365df33584b75bff782a510c6cd4883a419833d50", h.String())
}
