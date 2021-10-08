package pipeline

import (
	"bytes"
	"context"
	"io"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func Test_sink(t *testing.T) {
	t.Run("positive scenario", func(t *testing.T) {
		ctx, cancel := context.WithCancel(context.Background())
		defer cancel()

		blockc := make(chan Block)

		buf := bytes.NewBuffer(nil)
		errc, err := sink(ctx, blockc, buf)
		require.NoError(t, err)

		for _, v := range []string{"1", "2", "3"} {
			blockc <- Block{Hash: newHash([]byte(v))}
		}
		close(blockc)

		err = <-errc
		assert.Nil(t, err)
		_, ok := <-errc
		assert.False(t, ok)

		assert.Equal(t, `9c2e4d8fe97d881430de4e754b4205b9c27ce96715231cffc4337340cb110280
0c08173828583fc6ecd6ecdbcca7b6939c49c242ad5107e39deb7b0a5996b903
80903da4e6bbdf96e8ff6fc3966b0cfd355c7e860bdd1caa8e4722d9230e40ac`, strings.TrimSpace(buf.String()))
	})

	t.Run("nil input", func(t *testing.T) {
		ctx, cancel := context.WithCancel(context.Background())
		defer cancel()

		errc, err := sink(ctx, nil, bytes.NewBuffer(nil))
		require.ErrorIs(t, err, errNilChannel)
		assert.Nil(t, errc)
	})

	t.Run("error from writer", func(t *testing.T) {
		ctx, cancel := context.WithCancel(context.Background())
		defer cancel()

		blockc := make(chan Block, 3)
		for _, v := range []string{"1", "2", "3"} {
			blockc <- Block{Hash: newHash([]byte(v))}
		}
		close(blockc)

		writeErr := io.ErrClosedPipe
		errc, err := sink(ctx, blockc, writerMock{err: writeErr})
		require.NoError(t, err)

		err = <-errc
		assert.ErrorIs(t, err, writeErr)
		_, ok := <-errc
		assert.False(t, ok)
	})
}

type writerMock struct {
	err error
}

func (w writerMock) Write(d []byte) (int, error) {
	return len(d), w.err
}
