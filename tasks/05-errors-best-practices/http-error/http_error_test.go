package httperr

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestHTTPError(t *testing.T) {
	cases := []struct {
		err  error
		code int
		text string
	}{
		{
			err:  ErrStatusOK,
			code: http.StatusOK,
			text: "200 OK",
		},
		{
			err:  ErrStatusBadRequest,
			code: http.StatusBadRequest,
			text: "400 Bad Request",
		},
		{
			err:  ErrStatusNotFound,
			code: http.StatusNotFound,
			text: "404 Not Found",
		},
		{
			err:  ErrStatusUnprocessableEntity,
			code: http.StatusUnprocessableEntity,
			text: "422 Unprocessable Entity",
		},
		{
			err:  ErrStatusInternalServerError,
			code: http.StatusInternalServerError,
			text: "500 Internal Server Error",
		},
	}

	for _, tt := range cases {
		_, ok := tt.err.(HTTPError)
		assert.True(t, ok)

		var c interface {
			Code() int
		}
		require.ErrorAs(t, tt.err, &c)
		assert.Equal(t, tt.code, c.Code())
		assert.Equal(t, tt.err.Error(), tt.text)
	}
}

func TestHTTPErrorIsInt(t *testing.T) {
	_ = HTTPError(200)
}
