package pipeline

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func Test_hashTransactions(t *testing.T) {
	cases := []struct {
		name                   string
		batches                [][]Transaction
		expectedErrFromChannel error
		expectedBlocksHashes   []string
	}{
		{
			name: "positive scenario",
			batches: [][]Transaction{
				{{ID: 1}, {ID: 2}, {ID: 3}},
				{{ID: 4}, {ID: 5}, {ID: 6}},
				{{ID: 7}, {ID: 8}, {ID: 9}},
				{{ID: 10}},
			},
			expectedBlocksHashes: []string{
				"5b7534123197114fa7e7459075f39d89ffab74b5c3f31fad48a025b931ff5a01",
				"7acbd68a6fde6cdee8a8b39c473a1445c6580b3e88cb7566b6397bf59cc55d7c",
				"60d73cf428325259d23acea7fd8f10f56b65a83a1047eae205facafa487b4494",
				"aa54dffd96a821a88b136fd40138a67f8825437bd4ea08ca86247c4eb9fd8ba7",
			},
		},
		{
			name: "2d and 4th batches are invalid",
			batches: [][]Transaction{
				{{ID: 1}, {ID: 2}, {ID: 3}},
				{},
				{{ID: 7}, {ID: 8}, {ID: 9}},
				{},
			},
			expectedErrFromChannel: errNothingToHash,
			expectedBlocksHashes: []string{
				"5b7534123197114fa7e7459075f39d89ffab74b5c3f31fad48a025b931ff5a01",
			},
		},
	}

	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			ctx, cancel := context.WithCancel(context.Background())
			defer cancel()

			batchc := make(chan []Transaction)
			go func() {
				defer close(batchc)
				for _, b := range tt.batches {
					select {
					case <-ctx.Done():
						return
					case batchc <- b:
					}
				}
			}()

			blockc, errc, err := hashTransactions(ctx, batchc)
			require.NoError(t, err)

			hashes := make([]string, 0, len(tt.expectedBlocksHashes))
			for b := range blockc {
				hashes = append(hashes, b.Hash.String())
			}
			assert.Equal(t, tt.expectedBlocksHashes, hashes)

			if tt.expectedErrFromChannel != nil {
				chErr := <-errc
				assert.ErrorIs(t, chErr, tt.expectedErrFromChannel)
			}
			_, ok := <-errc
			assert.False(t, ok)
		})
	}
}

func Test_hashTransactions_NilInput(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	blockc, errc, err := hashTransactions(ctx, nil)
	require.ErrorIs(t, err, errNilChannel)
	assert.Nil(t, blockc)
	assert.Nil(t, errc)
}

func Test_hashTransactions_CancelledContext(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	batchc := make(chan []Transaction, 2)
	batchc <- []Transaction{{ID: 1}, {ID: 2}, {ID: 3}}
	batchc <- []Transaction{{ID: 4}, {ID: 5}, {ID: 6}}

	blockc, errc, err := hashTransactions(ctx, batchc)
	require.NoError(t, err)

	b, ok := <-blockc
	require.True(t, ok)
	assert.Equal(t, "5b7534123197114fa7e7459075f39d89ffab74b5c3f31fad48a025b931ff5a01", b.Hash.String())

	cancel()
	time.Sleep(time.Second)

	_, ok = <-blockc
	assert.False(t, ok)
	_, ok = <-errc
	assert.False(t, ok)
}
