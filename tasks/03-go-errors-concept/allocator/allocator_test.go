package memory

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestErrors(t *testing.T) {
	assert.EqualError(t, new(NotPermittedError), "operation not permitted")
	assert.EqualError(t, new(ArgOutOfDomainError), "numerical argument out of domain of func")
}

func TestAllocateUserID(t *testing.T) {
	t.Run("allocate from admin", func(t *testing.T) {
		buffer, err := Allocate(Admin, MinMemoryBlock)
		assert.NoError(t, err)
		assert.NotEmpty(t, buffer)
	})

	t.Run("allocate from unknown user", func(t *testing.T) {
		buffer, err := Allocate(123, MinMemoryBlock)
		assert.Error(t, err)
		assert.True(t, isNotPermittedError(err))
		assert.Nil(t, buffer)
	})
}

func TestAllocateSize(t *testing.T) {
	t.Run("allocate min memory block", func(t *testing.T) {
		buffer, err := Allocate(Admin, MinMemoryBlock)
		assert.NoError(t, err)
		assert.Len(t, buffer, MinMemoryBlock)
	})

	t.Run("allocate too low memory block", func(t *testing.T) {
		buffer, err := Allocate(Admin, 512)
		assert.Error(t, err)
		assert.True(t, isArgOutOfDomainError(err))
		assert.Nil(t, buffer)
	})

	t.Run("allocate memory block with valid size", func(t *testing.T) {
		buffer, err := Allocate(Admin, 2048)
		assert.NoError(t, err)
		assert.Len(t, buffer, 2048)
	})
}

func isNotPermittedError(err error) bool {
	_, ok := err.(*NotPermittedError)
	return ok
}

func isArgOutOfDomainError(err error) bool {
	_, ok := err.(*ArgOutOfDomainError)
	return ok
}
