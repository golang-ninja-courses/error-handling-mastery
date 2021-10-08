package server_test

import (
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/www-golang-courses-ru/advanced-dealing-with-errors-in-go/tasks/05-errors-best-practices/json-writer/server"
)

func TestServer_HandleIndex(t *testing.T) {
	t.Run("no error", func(t *testing.T) {
		result := struct {
			Result int `json:"result"`
		}{Result: 42}

		l := new(statefulLogger)
		s := server.New(l, dataProviderMock{v: result})

		r := httptest.NewRequest(http.MethodGet, "http://localhost:8080/", nil)
		w := httptest.NewRecorder()
		s.HandleIndex(w, r)

		resp := w.Result()
		body, err := io.ReadAll(resp.Body)
		require.NoError(t, err)

		assert.Equal(t, http.StatusOK, resp.StatusCode)
		assert.Equal(t, "application/json", resp.Header.Get("content-type"))
		assert.JSONEq(t, `{"result": 42}`, string(body))
		assert.Empty(t, l.lastMsg)
	})

	t.Run("json encoding error", func(t *testing.T) {
		type Result struct {
			Result *Result `json:"result"`
		}
		var result Result
		result.Result = &result // Циклическая структура.

		l := new(statefulLogger)
		s := server.New(l, dataProviderMock{v: result})

		r := httptest.NewRequest(http.MethodGet, "http://localhost:8080/", nil)
		w := httptest.NewRecorder()
		s.HandleIndex(w, r)

		resp := w.Result()
		body, err := io.ReadAll(resp.Body)
		require.NoError(t, err)

		assert.Equal(t, http.StatusInternalServerError, resp.StatusCode)
		assert.NotEqual(t, "application/json", resp.Header.Get("content-type"))
		assert.Contains(t, string(body), "json: unsupported value")
		assert.Contains(t, l.lastMsg, "json: unsupported value")
	})
}

type statefulLogger struct {
	lastMsg string
}

func (s *statefulLogger) Error(msg string) {
	s.lastMsg = msg
}

type dataProviderMock struct {
	v interface{}
}

func (d dataProviderMock) Data() interface{} {
	return d.v
}
