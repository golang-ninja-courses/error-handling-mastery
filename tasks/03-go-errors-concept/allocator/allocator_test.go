package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAllocate_UserID(t *testing.T) {
	_, err := Allocate(Admin, MinMemoryBlock)
	assert.NoError(t, err)

	_, err = Allocate(123, MinMemoryBlock)
	assert.Error(t, err)
	assert.EqualError(t, err, "operation not permitted")
}

func TestAllocate_Size(t *testing.T) {
	_, err := Allocate(Admin, MinMemoryBlock)
	assert.NoError(t, err)

	_, err = Allocate(Admin, 512)
	assert.Error(t, err)
	assert.EqualError(t, err, "math argument out of domain of func")

	_, err = Allocate(Admin, 2048)
	assert.NoError(t, err)
}

func TestAllocate(t *testing.T) {
	buffer, err := Allocate(Admin, MinMemoryBlock)
	assert.NoError(t, err)
	assert.Len(t, buffer, MinMemoryBlock)

	buffer, err = Allocate(Admin, 2048)
	assert.NoError(t, err)
	assert.Len(t, buffer, 2048)
}
