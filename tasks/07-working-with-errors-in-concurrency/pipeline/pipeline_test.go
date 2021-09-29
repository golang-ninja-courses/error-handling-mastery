package pipeline

import (
	"context"
	"errors"
	"testing"

	"go.uber.org/goleak"

	"github.com/stretchr/testify/assert"
)

func TestMain(m *testing.M) {
	goleak.VerifyTestMain(m)
}

func TestSource(t *testing.T) {
	cases := []struct {
		name         string
		input        []Transaction
		output       []Transaction
		errcExpected error
		errExpected  error
	}{
		{
			name:   "ok",
			input:  []Transaction{{1}, {2}, {3}},
			output: []Transaction{{1}, {2}, {3}},
		},
		{
			name:        "empty input",
			input:       []Transaction{},
			errExpected: errEmptyInput,
		},
		{
			name:         "invalid input",
			input:        []Transaction{{1}, {}, {3}},
			output:       []Transaction{{1}},
			errcExpected: errEmptyTx,
		},
	}

	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			out, errc, err := Source(context.Background(), tt.input...)

			if tt.errExpected != nil {
				assert.ErrorIs(t, err, tt.errExpected)
				return
			}

			txs := make([]Transaction, 0, len(tt.input))
			for tx := range out {
				txs = append(txs, tx)
			}
			assert.Equal(t, tt.output, txs)

			assert.ErrorIs(t, tt.errcExpected, <-errc)
		})
	}
}

func TestSource_CancelledContext(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	cancel()

	out, errc, err := Source(ctx, []Transaction{{1}, {2}, {3}}...)

	assert.NoError(t, err)

	var txs []Transaction
	for tx := range out {
		txs = append(txs, tx)
	}

	switch len(txs) {
	case 0:
	case 1:
		assert.Equal(t, int64(1), txs[0].ID)
	default:
		t.Error("out should be empty or contain single value")
	}

	assert.Empty(t, <-errc)
}

func TestAggregate(t *testing.T) {
	batchSize := 2
	in := make(chan Transaction)
	expected := [][]Transaction{{{1}, {2}}, {{3}, {4}}, {{5}}}

	go func() {
		defer close(in)
		for _, tx := range []Transaction{{1}, {2}, {3}, {4}, {5}} {
			in <- tx
		}
	}()

	out, errc, err := Aggregate(context.Background(), batchSize, in)
	assert.NoError(t, err)

	i := 0
	for tx := range out {
		assert.Equal(t, expected[i], tx)
		i++
	}

	assert.NoError(t, <-errc)
}

func TestAggregate_InvalidBatchSize(t *testing.T) {
	for _, batchSize := range []int{0, maxBatchSize + 1} {
		out, errc, err := Aggregate(context.Background(), batchSize, nil)
		assert.Nil(t, out)
		assert.Nil(t, errc)
		assert.ErrorIs(t, err, errInvalidBatchSize)
	}
}

func TestAggregate_CancelledContext(t *testing.T) {
	testCtx, testCtxCancel := context.WithCancel(context.Background())
	defer testCtxCancel()

	ctx, cancel := context.WithCancel(testCtx)
	cancel()

	inFunc := func(len int) <-chan Transaction {
		in := make(chan Transaction)
		go func() {
			defer close(in)
			for i := 1; i <= len; i++ {
				select {
				case in <- Transaction{int64(i)}:
				case <-testCtx.Done():
				}
			}
		}()
		return in
	}

	cases := []struct {
		name      string
		batchSize int
		inLen     int
	}{
		{
			name:      `batch size greater than "in" length`,
			batchSize: 3,
			inLen:     1,
		},
		{
			name:      `batch size equals than "in" length`,
			batchSize: 1,
			inLen:     1,
		},
		{
			name:      `batch size less than "in" length`,
			batchSize: 1,
			inLen:     3,
		},
	}

	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			out, errc, err := Aggregate(ctx, tt.batchSize, inFunc(tt.inLen))
			assert.NoError(t, err)

			var txs [][]Transaction
			for tx := range out {
				txs = append(txs, tx)
			}

			switch len(txs) {
			case 0:
			case 1:
				assert.Equal(t, []Transaction{{1}}, txs[0])
			default:
				t.Error("out should be empty or contain single value")
			}

			assert.Empty(t, <-errc)
		})
	}
}

func TestHashTxs(t *testing.T) {
	in := make(chan []Transaction)
	expectedHashes := []string{
		"f981662b1dcd91b2569a56fce8c590b04bc062ee22d459e49bc507638c8099a2",
		"aabd9871539c37bda9f77bf47440df5a57c2a5736a04387d1c3b92dffefa47e4",
		"2183a742b2c40c0f90befe5b460fe0200646897c803b14513d754ab5929b1a2b",
	}

	go func() {
		defer close(in)
		for _, s := range [][]Transaction{{{1}, {2}, {3}}, {{4}, {5}}, {{6}}} {
			in <- s
		}
	}()

	out, errc, err := HashTxs(context.Background(), in)
	assert.NoError(t, err)

	i := 0
	for block := range out {
		assert.Equal(t, expectedHashes[i], block.Hash.String())
		i++
	}

	assert.NoError(t, <-errc)
}

func TestHashTxs_CalculateHashError(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	in := make(chan []Transaction)

	go func() {
		defer close(in)
		for _, s := range [][]Transaction{nil, nil, nil} {
			select {
			case in <- s:
			case <-ctx.Done():
				return
			}
		}
	}()

	_, errc, err := HashTxs(ctx, in)
	assert.NoError(t, err)

	assert.Error(t, <-errc)
}

func TestHashTxs_CancelledContext(t *testing.T) {
	testCtx, testCtxCancel := context.WithCancel(context.Background())
	defer testCtxCancel()

	ctx, cancel := context.WithCancel(testCtx)
	cancel()

	in := make(chan []Transaction)

	go func() {
		defer close(in)
		for _, txs := range [][]Transaction{{{1}}, {{2}}} {
			select {
			case in <- txs:
			case <-testCtx.Done():
			}
		}
	}()

	out, errc, err := HashTxs(ctx, in)

	assert.NoError(t, err)

	var blocks []Block
	for block := range out {
		blocks = append(blocks, block)
	}

	switch len(blocks) {
	case 0:
	case 1:
		assert.Equal(t, "488ad59efeafae6af0be932803b65470b908c33377f4ff99aad1fd8c68c4f463", blocks[0].Hash.String())
	default:
		t.Error("out should be empty or contain single value")
	}

	assert.Empty(t, <-errc)
}

func TestSink(t *testing.T) {
	blocks := []Block{{hash([]byte("1"))}, {hash([]byte("2"))}, {hash([]byte("3"))}}
	in := make(chan Block)
	go func() {
		defer close(in)
		for _, block := range blocks {
			in <- block
		}
	}()

	errc, err := Sink(context.Background(), in)
	assert.NoError(t, err)
	assert.NoError(t, <-errc)
}

func TestMerge(t *testing.T) {
	errs := []error{errors.New("1"), errors.New("2"), errors.New("3")}
	errcs := make([]<-chan error, 0, 3)

	for _, err := range errs {
		errc := make(chan error, 1)
		errc <- err
		close(errc)

		errcs = append(errcs, errc)
	}

	out := Merge(errcs...)

	for err := range out {
		assert.Contains(t, errs, err)
	}
}

func TestPipeline(t *testing.T) {
	cases := []struct {
		name        string
		txs         []Transaction
		expectError bool
	}{
		{
			name: "ok",
			txs:  []Transaction{{1}, {2}, {3}, {4}, {5}},
		},
		{
			name:        "tx with zero ID",
			txs:         []Transaction{{1}, {2}, {0}, {4}, {5}},
			expectError: true,
		},
	}

	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			err := Pipeline(context.Background(), tt.txs...)

			if tt.expectError {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}
