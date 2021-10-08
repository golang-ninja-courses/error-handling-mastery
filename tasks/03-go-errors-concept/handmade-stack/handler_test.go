package handmadestack

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_handler(t *testing.T) {
	t.Run("no errors", func(t *testing.T) {
		getEntity = func() (Entity, error) { return Entity{ID: "some-id"}, nil }
		updateEntity = func(e Entity) error { return nil }
		runInTransaction = func(f func() error) error { return f() }

		e, err := handler()
		require.NoError(t, err)
		require.Equal(t, Entity{ID: "some-id"}, e)
	})

	t.Run("check unique wrapping", func(t *testing.T) {
		var errs []string
		collectErr := func() {
			if e, err := handler(); err != nil {
				require.Equal(t, Entity{}, e)
				errs = append(errs, err.Error())
			}
		}

		// first transaction error
		runInTransaction = func(f func() error) error {
			return ErrInitTransaction
		}
		collectErr()

		// second transaction error
		var i int
		runInTransaction = func(f func() error) error {
			if i == 1 {
				return ErrInitTransaction
			}
			i++
			return f()
		}
		collectErr()

		// third transaction error
		i = 0
		runInTransaction = func(f func() error) error {
			if i == 2 {
				return ErrInitTransaction
			}
			i++
			return f()
		}
		collectErr()
		runInTransaction = func(f func() error) error { return f() }

		// getEntity error
		getEntity = func() (Entity, error) {
			return Entity{}, ErrExecSQL
		}
		collectErr()
		getEntity = func() (Entity, error) { return Entity{ID: "some-id"}, nil }

		// first updateEntity error
		updateEntity = func(e Entity) error {
			return ErrExecSQL
		}
		collectErr()

		// second updateEntity error
		var j int
		updateEntity = func(e Entity) error {
			if j == 1 {
				return ErrExecSQL
			}
			j++
			return nil
		}
		collectErr()

		// third updateEntity error
		j = 0
		updateEntity = func(e Entity) error {
			if j == 2 {
				return ErrExecSQL
			}
			j++
			return nil
		}
		collectErr()

		t.Log("Received errors:\n" + strings.Join(errs, "\n"))
		require.Len(t, errs, 7)

		uniqErrs := make(map[string]struct{}, len(errs))
		for _, s := range errs {
			uniqErrs[s] = struct{}{}
		}
		require.Len(t, uniqErrs, len(errs), "errors are not wrapped in a unique message")
	})
}
