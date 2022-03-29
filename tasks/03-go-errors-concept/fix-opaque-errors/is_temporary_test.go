package errors

import (
	"errors"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsTemporary(t *testing.T) {
	cases := []struct {
		err         error
		isTemporary bool
	}{
		{
			err:         errors.New("integrity error"),
			isTemporary: false,
		},
		{
			err:         newNetErrorMock(true),
			isTemporary: true,
		},
		{
			err:         fmt.Errorf("wrap 1: %w", newNetErrorMock(true)),
			isTemporary: true,
		},
		{
			err:         fmt.Errorf("wrap 2: %w", fmt.Errorf("wrap 1: %w", newNetErrorMock(true))),
			isTemporary: true,
		},
		{
			err:         newNetErrorMock(false),
			isTemporary: false,
		},
		{
			err:         fmt.Errorf("wrap 1: %w", newNetErrorMock(false)),
			isTemporary: false,
		},
		{
			err:         fmt.Errorf("wrap 2: %w", fmt.Errorf("wrap 1: %w", newNetErrorMock(false))),
			isTemporary: false,
		},
	}

	for i, tt := range cases {
		t.Run(fmt.Sprintf("case %d", i), func(t *testing.T) {
			assert.Equal(t, tt.isTemporary, IsTemporary(tt.err))
		})
	}
}

type netErrorMock struct { //nolint:errname
	isTmp bool
}

func newNetErrorMock(isTmp bool) *netErrorMock {
	return &netErrorMock{isTmp: isTmp}
}

func (n *netErrorMock) Error() string {
	return fmt.Sprintf("network error (temporary: %t)", n.isTmp)
}

func (n *netErrorMock) IsTemporary() bool {
	return n.isTmp
}
