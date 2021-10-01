package pipeline

import (
	"bytes"
	"context"
	"io"
	"os"
	"strings"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"go.uber.org/goleak"
)

func TestMain(m *testing.M) {
	goleak.VerifyTestMain(m)
}

func ExamplePipeline() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	err := Pipeline(ctx, 2, os.Stdout,
		Transaction{ID: 100},
		Transaction{ID: 200},
		Transaction{ID: 300})
	if err != nil {
		panic(err)
	}

	// Output:
	// a6c9627e9b7a9724c62a4be4745a3ad01b99e219f17e18f90793983f1e567590
	// 54f46fa1dcd09720d9f4979bd6a4cd1363a9e6118a38580bbbff008e9b59449a
}

func TestPipeline(t *testing.T) {
	cases := []struct {
		name                string
		batchSize           int
		txs                 []Transaction
		expectedErr         error
		expectedHashes      string
		expectedHashesCount int
	}{
		{
			name:        "immediately error from batch",
			batchSize:   0,
			txs:         makeTransactions(3),
			expectedErr: errInvalidBatchSize,
		},
		{
			name:        "error from batch",
			batchSize:   4,
			txs:         []Transaction{{ID: 1}, {ID: 2}, {ID: 3}, {ID: 0}},
			expectedErr: errEmptyTx,
		},
		{
			name:      "positive scenario",
			batchSize: 4,
			txs:       makeTransactions(100),
			expectedHashes: `8b186d4723474e69fd14c28384063e2031d5da66b97844d5973a9e9bf7dcfeeb
07f3af090be94597a0903e91a313b7e461ffbec868eebf9cc88ae75c539ed2eb
42b9b918692a1e74ba4d40386fc76ab6ceb2fa73b53598d29d1089ccdae0c750
c607a9c06c2bfb6e974628c30bd17bed2fe23c613e97c36430cfbf607acb03f7
24e0e160c195b44319e0d84f18124fa7800de7b6681df452d280ae92cc0d6c86
a3f805a47570ea508908e0ab2b1cf8f5218f2ee61a71e16c2bb03553c0cb42ae
a3cce01bacb91aaa72e8682577b12d6e4dab43a6ff8bfe0bcf43f52054ee11c3
874a358653269a55affcca5fa55c75c695de3172c110a4d49bdf036090346674
b418dbba7fe6563ee40ef8c1a834395357a0e741bf3e14bd8e68bb4af6d3a50f
3a8a862e6e349b5e215d7a0a02bacf2398f620eef93c0ec748f14b967b9eabc6
ae02d51274acc771addd0c0db567434a1c9433a996f2536add4ba0dbcda0ee73
97313cf74036b725429ef69ce2bbea204dd4474b12d0cdb681392d4c5f926db3
c470769604483aa52d0a291c3647f3964773411652197be1921d62a4a12dc90a
dc4a415b82eaaa41ad4eb5676ef0416428e7178e6423a753071d43610183ec51
d8bc56364c602e42c0db0de38250d397a0104f15adefd66bdfb3de165126c735
3e17151589260b90fbe0fb652dc1b1e5a21f32d3120f316eb1111e705c7248c6
984d8612105bb1ce8ff8f967d660a49493db8ee62f2e894d7360c993b40d730e
033c75caae1148d459c637e2c98cb8344e783d3ae91e18ba14b3a5bbc7a99341
5b3298ab289ac42bd0c641e83b21f799976d9f1ff42fb259fa9db8e8946add1f
5065bc02e095be9e88166bbe3bc2d4f191d332d2884f823ddbf77da5d65afa25
f73280794c2fd11d11021ccf7b8584a58d137fc086c93c36fb5b9af5d9a3cdf7
8b906c632e5b3c7772aad478cbe51315938f076c73a4891ec5149ba1ba52f416
2ab95e7abd88dcb2662a7206e055a3ee49b9a91ddfd60c2333c1c5735a9f2e4e
cb6095d7b5963e6efcd02d5223a5f0856c3a956a675ef9a62843cb260d3e0a6a
2d476ba57977ecc287d8d56af25360d336ac677f3a388be178f102aa6555dd12`,
		},
		{
			name:                "big data processing",
			batchSize:           2,
			txs:                 makeTransactions(1_000_000),
			expectedHashesCount: 1_000_000 / 2,
		},
	}

	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
			defer cancel()

			buf := bytes.NewBuffer(nil)
			err := Pipeline(ctx, tt.batchSize, buf, tt.txs...)
			require.ErrorIs(t, err, tt.expectedErr)

			if tt.expectedHashes != "" {
				assert.Equal(t, tt.expectedHashes, strings.TrimSpace(buf.String()))
			}

			if tt.expectedHashesCount != 0 {
				received := strings.Split(strings.TrimSpace(buf.String()), "\n")
				assert.Len(t, received, tt.expectedHashesCount)
			}
		})
	}
}

func TestPipeline_Cancelling(t *testing.T) {
	writerLatency := time.Second / 3
	timeout := time.Second

	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	w := slowWriter{
		w:       bytes.NewBuffer(nil),
		latency: writerLatency,
	}
	err := Pipeline(ctx, 4, w, makeTransactions(10_000_000)...)
	require.NoError(t, err)
}

type slowWriter struct {
	w       io.Writer
	latency time.Duration
}

func (w slowWriter) Write(d []byte) (int, error) {
	time.Sleep(w.latency)
	return w.w.Write(d)
}

func makeTransactions(n int) []Transaction {
	txs := make([]Transaction, n)
	for i := 1; i <= n; i++ {
		txs[i-1] = Transaction{ID: int64(i)}
	}
	return txs
}
