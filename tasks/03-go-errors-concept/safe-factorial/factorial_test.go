package factorial

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCalculate(t *testing.T) {
	cases := []struct {
		name                string
		n                   int
		expectedErr         error
		expectedErrContains string
		expectedFactorial   int
	}{
		{
			name:        "negative n",
			n:           -1,
			expectedErr: ErrNegativeN,
		},
		{
			name:                "too deep",
			n:                   1e3,
			expectedErr:         ErrTooDeep,
			expectedErrContains: "256",
		},
		{
			name:              "factorial of 0 is 1",
			n:                 0,
			expectedFactorial: 1,
		},
		{
			name:              "factorial of 10 is 3628800",
			n:                 10,
			expectedFactorial: 3628800,
		},
	}

	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			f, err := Calculate(tt.n)
			require.ErrorIs(t, err, tt.expectedErr)
			if tt.expectedErrContains != "" {
				require.Error(t, err)
				require.Contains(t, err.Error(), tt.expectedErrContains)
			}
			require.Equal(t, tt.expectedFactorial, f)
		})
	}
}
