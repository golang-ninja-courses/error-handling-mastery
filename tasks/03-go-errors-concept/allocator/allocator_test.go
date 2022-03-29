package allocator

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestErrors(t *testing.T) {
	assert.EqualError(t, new(NotPermittedError), "operation not permitted")
	assert.EqualError(t, new(ArgOutOfDomainError), "numerical argument out of domain of func")
}

func TestAllocate(t *testing.T) {
	cases := []struct {
		name                  string
		uid                   int
		size                  int
		isNotPermittedError   bool
		isArgOutOfDomainError bool
	}{
		{
			name:                  "invalid size from admin",
			uid:                   Admin,
			size:                  1023,
			isArgOutOfDomainError: true,
		},
		{
			name: "min valid size from admin",
			uid:  Admin,
			size: 1024,
		},
		{
			name: "valid size from admin",
			uid:  Admin,
			size: 2048,
		},
		{
			name:                "invalid size from unknown user",
			uid:                 42,
			size:                1023,
			isNotPermittedError: true,
		},
		{
			name:                "min valid size from unknown user",
			uid:                 42,
			size:                1024,
			isNotPermittedError: true,
		},
		{
			name:                "valid size from unknown user",
			uid:                 42,
			size:                2048,
			isNotPermittedError: true,
		},
	}

	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			buffer, err := Allocate(tt.uid, tt.size)
			if tt.isNotPermittedError {
				assert.True(t, isNotPermittedError(err))
				assert.Nil(t, buffer)
			} else if tt.isArgOutOfDomainError {
				assert.True(t, isArgOutOfDomainError(err))
				assert.Nil(t, buffer)
			} else {
				require.NoError(t, err)
				assert.Len(t, buffer, tt.size)
			}
		})
	}
}

func isNotPermittedError(err error) bool {
	_, ok := err.(*NotPermittedError)
	return ok
}

func isArgOutOfDomainError(err error) bool {
	_, ok := err.(*ArgOutOfDomainError)
	return ok
}
