package fetcher

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"testing"

	"github.com/stretchr/testify/require"
)

var errUnknownNetErr = errors.New("cannot do request")

func TestFetchURL(t *testing.T) {
	cases := []struct {
		client       clientMock
		expectedData []byte
		expectedErr  error
	}{
		{
			client:      newClientMock(nil, newNetErrMock(false, true)),
			expectedErr: ErrFetchTimeout,
		},
		{
			client:      newClientMock(nil, newNetErrMock(true, false)),
			expectedErr: ErrFetchTmp,
		},
		{
			client:      newClientMock(nil, errUnknownNetErr),
			expectedErr: errUnknownNetErr,
		},
		{
			client: newClientMock(&http.Response{
				Body: ioutil.NopCloser(bytes.NewReader([]byte("hello"))),
			}, nil),
			expectedData: []byte("hello"),
		},
	}

	for i, tt := range cases {
		t.Run(fmt.Sprintf("case %d", i), func(t *testing.T) {
			ctx, cancel := context.WithCancel(context.Background())
			defer cancel()

			data, err := FetchURL(ctx, tt.client, "test-url")
			require.ErrorIs(t, err, tt.expectedErr)
			require.Equal(t, tt.expectedData, data)
		})
	}
}

type clientMock struct {
	doResponse *http.Response
	doErr      error
}

func newClientMock(resp *http.Response, err error) clientMock {
	return clientMock{resp, err}
}

func (c clientMock) Do(_ *http.Request) (*http.Response, error) {
	return c.doResponse, c.doErr
}

type netErrMock struct { // Реализуй меня.
}

func newNetErrMock(isTemporary, isTimeout bool) *netErrMock {
	// Реализуй меня.
	return nil
}
