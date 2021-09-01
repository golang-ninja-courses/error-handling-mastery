package index

import (
	"fmt"
	"math"
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestGetIndexFromFileName(t *testing.T) {
	// Заполните кейсы ниже так, чтобы тест проходил.
	cases := []struct {
		fileName      string
		expectedIndex int
		expectedError error
	}{
		{
			fileName:      "parsed_page",
			expectedIndex: 0,
			expectedError: nil,
		},
		{
			fileName:      "parsedpage",
			expectedIndex: 0,
			expectedError: nil,
		},
		{
			fileName:      "parsed_page_",
			expectedIndex: 0,
			expectedError: nil,
		},
		{
			fileName:      "parsed_page_100_suffix",
			expectedIndex: 0,
			expectedError: nil,
		},
		{
			fileName:      "parsed_page_-1",
			expectedIndex: 0,
			expectedError: nil,
		},
		{
			fileName:      "parsed_page_0",
			expectedIndex: 0,
			expectedError: nil,
		},
		{
			fileName:      "parsed_page_15.5",
			expectedIndex: 0,
			expectedError: nil,
		},
		{
			fileName:      "parsed_page_1000",
			expectedIndex: 0,
			expectedError: nil,
		},
		{
			fileName:      fmt.Sprintf("parsed_page_%d", math.MaxInt32+1),
			expectedIndex: 0,
			expectedError: nil,
		},
	}

	for _, tt := range cases {
		index, err := GetIndexFromFileName(tt.fileName)
		require.ErrorIs(t, err, tt.expectedError)
		assert.Equal(t, tt.expectedIndex, index)

		_, ok := tt.expectedError.(*strconv.NumError)
		assert.False(t, ok, "do not use *strconv.NumError directly, look for more specific errors")
	}
}
