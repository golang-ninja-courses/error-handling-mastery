package requests

import (
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func ExampleSearchRequest_Validate() {
	req := SearchRequest{
		Exp:      "(.*golang.*",
		Page:     -1,
		PageSize: 3000,
	}

	err := req.Validate()
	if err == nil {
		panic("invalid Validate() realization")
	}

	fmt.Println(errorMsg(err))
	// Output:
	// validation errors:
	//     exp is not regexp: error parsing regexp: missing closing ): `(.*golang.*`
	//     invalid page: -1
	//     invalid page size: 3000 > 100
}

func errorMsg(err error) string {
	return strings.TrimSpace(strings.ReplaceAll(err.Error(), "\t", "    "))
}

func TestSearchRequest_Validate(t *testing.T) {
	cases := []struct {
		name        string
		req         SearchRequest
		expectedErr error
	}{
		{
			name:        "no error",
			req:         SearchRequest{Exp: ".*golang.*", Page: 3, PageSize: 10},
			expectedErr: nil,
		},
		{
			name:        "invalid regexp",
			req:         SearchRequest{Exp: "(.*golang.*", Page: 3, PageSize: 10},
			expectedErr: errIsNotRegexp,
		},
		{
			name:        "invalid page",
			req:         SearchRequest{Exp: ".*golang.*", Page: -3, PageSize: 10},
			expectedErr: errInvalidPage,
		},
		{
			name:        "invalid page size (too small)",
			req:         SearchRequest{Exp: ".*golang.*", Page: 3, PageSize: -1},
			expectedErr: errInvalidPageSize,
		},
		{
			name:        "invalid page size (too big)",
			req:         SearchRequest{Exp: ".*golang.*", Page: 3, PageSize: 1000},
			expectedErr: errInvalidPageSize,
		},
	}

	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.req.Validate()
			assert.ErrorIs(t, err, tt.expectedErr)
		})
	}
}

func TestSearchRequest_Validate_multipleErr(t *testing.T) {
	req := SearchRequest{
		Exp:      "(.*golang.*",
		Page:     -1,
		PageSize: 3000,
	}

	err := req.Validate()
	require.Error(t, err)

	assert.ErrorIs(t, err, errIsNotRegexp)
	assert.Contains(t, err.Error(), "(.*golang.*")

	assert.ErrorIs(t, err, errInvalidPage)
	assert.Contains(t, err.Error(), "-1")

	assert.ErrorIs(t, err, errInvalidPageSize)
	assert.Contains(t, err.Error(), "3000 > 100")

	// Negative PageSize.
	req.PageSize = -1
	err = req.Validate()
	require.Error(t, err)
	assert.ErrorIs(t, err, errInvalidPageSize)
	assert.Contains(t, err.Error(), "-1 < 0")
}
