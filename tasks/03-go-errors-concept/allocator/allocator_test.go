package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestErrors(t *testing.T) {
	assert.Implements(t, (*error)(nil), new(NotPermittedError))
	assert.Implements(t, (*error)(nil), new(ArgOutOfDomainError))
}

func TestAllocate_UserID(t *testing.T) {
	_, err := Allocate(Admin, MinMemoryBlock)
	assert.NoError(t, err)

	errNotPermitted := &NotPermittedError{}

	_, err = Allocate(123, MinMemoryBlock)
	assert.Error(t, err)
	assert.ErrorAs(t, err, &errNotPermitted)
}

func TestAllocate_Size(t *testing.T) {
	_, err := Allocate(Admin, MinMemoryBlock)
	assert.NoError(t, err)

	errArgOutOfDomain := &ArgOutOfDomainError{}

	_, err = Allocate(Admin, 512)
	assert.Error(t, err)
	assert.ErrorAs(t, err, &errArgOutOfDomain)

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
