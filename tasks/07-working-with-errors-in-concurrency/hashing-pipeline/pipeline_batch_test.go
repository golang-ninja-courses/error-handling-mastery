package pipeline

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func Test_batch(t *testing.T) {
	txs := makeTransactions(10)

	cases := []struct {
		name                   string
		batchSize              int
		txs                    []Transaction
		expectedErr            error
		expectedErrFromChannel error
		expectedBatches        [][]Transaction
	}{
		{
			name:        "too small batch size",
			batchSize:   0,
			txs:         txs,
			expectedErr: errInvalidBatchSize,
		},
		{
			name:        "too big batch size",
			batchSize:   maxBatchSize + 1,
			txs:         txs,
			expectedErr: errInvalidBatchSize,
		},
		{
			name:        "nil input",
			batchSize:   maxBatchSize - 1,
			txs:         nil,
			expectedErr: errEmptyInput,
		},
		{
			name:        "zero length input",
			batchSize:   maxBatchSize - 1,
			txs:         []Transaction{},
			expectedErr: errEmptyInput,
		},
		{
			name:      "batch size 1",
			batchSize: 1,
			txs:       txs,
			expectedBatches: [][]Transaction{
				{{ID: 1}},
				{{ID: 2}},
				{{ID: 3}},
				{{ID: 4}},
				{{ID: 5}},
				{{ID: 6}},
				{{ID: 7}},
				{{ID: 8}},
				{{ID: 9}},
				{{ID: 10}},
			},
		},
		{
			name:      "batch size 2",
			batchSize: 2,
			txs:       txs,
			expectedBatches: [][]Transaction{
				{{ID: 1}, {ID: 2}},
				{{ID: 3}, {ID: 4}},
				{{ID: 5}, {ID: 6}},
				{{ID: 7}, {ID: 8}},
				{{ID: 9}, {ID: 10}},
			},
		},
		{
			name:      "batch size 3",
			batchSize: 3,
			txs:       txs,
			expectedBatches: [][]Transaction{
				{{ID: 1}, {ID: 2}, {ID: 3}},
				{{ID: 4}, {ID: 5}, {ID: 6}},
				{{ID: 7}, {ID: 8}, {ID: 9}},
				{{ID: 10}},
			},
		},
		{
			name:      "batch size 4",
			batchSize: 4,
			txs:       txs,
			expectedBatches: [][]Transaction{
				{{ID: 1}, {ID: 2}, {ID: 3}, {ID: 4}},
				{{ID: 5}, {ID: 6}, {ID: 7}, {ID: 8}},
				{{ID: 9}, {ID: 10}},
			},
		},
		{
			name:      "invalid transaction in second batch",
			batchSize: 3,
			txs: []Transaction{
				{ID: 1},
				{ID: 2},
				{ID: 3},
				{ID: 4},
				{ID: 0},
				{ID: 6},
				{ID: 7},
				{ID: 8},
				{ID: 9},
			},
			expectedBatches: [][]Transaction{
				{{ID: 1}, {ID: 2}, {ID: 3}},
			},
			expectedErrFromChannel: errEmptyTx,
		},
	}

	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			ctx, cancel := context.WithCancel(context.Background())
			defer cancel()

			batchc, errc, err := batch(ctx, tt.batchSize, tt.txs...)
			require.ErrorIs(t, err, tt.expectedErr)

			if tt.expectedErr != nil {
				assert.Nil(t, batchc)
				assert.Nil(t, errc)
				return
			}

			batches := make([][]Transaction, 0, len(tt.expectedBatches))
			for b := range batchc {
				batches = append(batches, b)
			}
			assert.Equal(t, tt.expectedBatches, batches)

			if tt.expectedErrFromChannel != nil {
				chErr := <-errc
				assert.ErrorIs(t, chErr, tt.expectedErrFromChannel)
			}
			_, ok := <-errc
			assert.False(t, ok)
		})
	}
}

func Test_batch_CancelledContext(t *testing.T) {
	const n = 4
	const batchSize = n - 1

	txs := makeTransactions(n)

	t.Run("blocking when sending the first batch", func(t *testing.T) {
		ctx, cancel := context.WithCancel(context.Background())
		defer cancel()

		batchc, errc, err := batch(ctx, batchSize, txs...)
		require.NoError(t, err)

		cancel()
		time.Sleep(time.Second)

		_, ok := <-batchc
		assert.False(t, ok)
		_, ok = <-errc
		assert.False(t, ok)
	})

	t.Run("blocking when sending the last batch", func(t *testing.T) {
		ctx, cancel := context.WithCancel(context.Background())
		defer cancel()

		batchc, errc, err := batch(ctx, batchSize, txs...)
		require.NoError(t, err)

		b, ok := <-batchc
		require.True(t, ok)
		assert.Equal(t, makeTransactions(batchSize), b)

		cancel()
		time.Sleep(time.Second)

		_, ok = <-batchc
		assert.False(t, ok)
		_, ok = <-errc
		assert.False(t, ok)
	})
}
