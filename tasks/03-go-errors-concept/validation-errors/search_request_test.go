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
	type or = []string

	type expected struct {
		err  error
		text or
	}

	cases := []struct {
		name string
		req  SearchRequest
		exp  []expected
	}{
		{
			name: "regexp  negative page  big page size",
			req: SearchRequest{
				Exp:      "(.*golang.*",
				Page:     -1,
				PageSize: 3000,
			},
			exp: []expected{
				{err: errIsNotRegexp, text: or{"(.*golang.*"}},
				{err: errInvalidPage, text: or{"-1"}},
				{err: errInvalidPageSize, text: or{"3000 > 100"}},
			},
		},
		{
			name: "regexp  negative page  zero page size",
			req: SearchRequest{
				Exp:      "(.*golang.*",
				Page:     -1,
				PageSize: 0,
			},
			exp: []expected{
				{err: errIsNotRegexp, text: or{"(.*golang.*"}},
				{err: errInvalidPage, text: or{"-1"}},
				{err: errInvalidPageSize, text: or{"0 < 1", "0 <= 0"}},
			},
		},
		{
			name: "regexp  negative page  negative page size",
			req: SearchRequest{
				Exp:      "(.*golang.*",
				Page:     -1,
				PageSize: -1,
			},
			exp: []expected{
				{err: errIsNotRegexp, text: or{"(.*golang.*"}},
				{err: errInvalidPage, text: or{"-1"}},
				{err: errInvalidPageSize, text: or{"-1 < 1", "-1 <= 0"}},
			},
		},
		{
			name: "regexp  big page size",
			req: SearchRequest{
				Exp:      "(.*golang.*",
				Page:     10,
				PageSize: 101,
			},
			exp: []expected{
				{err: errIsNotRegexp, text: or{"(.*golang.*"}},
				{err: errInvalidPageSize, text: or{"101 > 100"}},
			},
		},
		{
			name: "negative page  big page size",
			req: SearchRequest{
				Page:     -2,
				PageSize: 101,
			},
			exp: []expected{
				{err: errInvalidPage, text: or{"-2"}},
				{err: errInvalidPageSize, text: or{"101 > 100"}},
			},
		},
	}

	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.req.Validate()
			require.Error(t, err)

			for _, e := range tt.exp {
				assert.ErrorIs(t, err, e.err)
				assert.True(t, containsAnyExclusively(err.Error(), e.text),
					"err msg %q do not contains exclusively one of %#v", err, e.text)
			}
		})
	}
}

func containsAnyExclusively(s string, subs []string) bool {
	var result bool
	for _, substr := range subs {
		c := strings.Contains(s, substr)
		if result && c {
			return false
		}
		result = result || c
	}
	return result
}
