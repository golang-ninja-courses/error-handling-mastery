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
		capacity              int
		isNotPermittedError   bool
		isArgOutOfDomainError bool
	}{
		{
			name:                  "invalid capacity from admin",
			uid:                   Admin,
			capacity:              1023,
			isArgOutOfDomainError: true,
		},
		{
			name:     "min valid capacity from admin",
			uid:      Admin,
			capacity: 1024,
		},
		{
			name:     "valid capacity from admin",
			uid:      Admin,
			capacity: 2048,
		},
		{
			name:                "invalid capacity from unknown user",
			uid:                 42,
			capacity:            1023,
			isNotPermittedError: true,
		},
		{
			name:                "min valid capacity from unknown user",
			uid:                 42,
			capacity:            1024,
			isNotPermittedError: true,
		},
		{
			name:                "valid capacity from unknown user",
			uid:                 42,
			capacity:            2048,
			isNotPermittedError: true,
		},
	}

	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			buffer, err := Allocate(tt.uid, tt.capacity)
			if tt.isNotPermittedError {
				assert.True(t, isNotPermittedError(err))
				assert.Nil(t, buffer)
			} else if tt.isArgOutOfDomainError {
				assert.True(t, isArgOutOfDomainError(err))
				assert.Nil(t, buffer)
			} else {
				require.NoError(t, err)
				assert.Equal(t, cap(buffer), tt.capacity)
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
